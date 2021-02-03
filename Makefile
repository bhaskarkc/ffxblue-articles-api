VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
PACKAGES := $(shell go list ./... | grep -v /vendor/)

## Load .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DSN="mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"

# https://github.com/golang-migrate/migrate#cli-usage
MIGRATE := docker run -t -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.14.1 -path=/migrations/ -database ${DSN}

PID_FILE := './.pid'
FSWATCH_FILE := './fswatch.cfg'

.PHONY: default
default: help

.PHONY: help
help: ## Prints this help screen.
	@printf "================================================\t\t\n"
	@printf "\t FFXBlue Articles API \t\t\n"
	@printf "================================================\t\t\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed -e "s/Makefile://" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: run
run: ## run the API server
	go run main.go

.PHONY: run-restart
run-restart: ## restart the API server
	@pkill -P `cat $(PID_FILE)` || true
	@printf '%*s\n' "80" '' | tr ' ' -
	@echo "Source file changed. Restarting server..."
	@go run main.go & echo $$! > $(PID_FILE)
	@printf '%*s\n' "80" '' | tr ' ' -

run-live: ## run the API server with live reload support (requires fswatch)
	@go run main.go & echo $$! > $(PID_FILE)
	@fswatch -x -o --event Created --event Updated --event Renamed -r controllers datasource domain logger server services utils | xargs -n1 -I {} make run-restart

.PHONY: build
build:  ## build the API server binary
	CGO_ENABLED=0 go build -a -o api-server .

.PHONY: up
up: ## Spins up docker container
	docker-compose up --build

.PHONY: down
down: ## Tear down the docker container
	docker-compose down

recreate: ## Force recreate and start the docker container
	docker-compose up --build --force-recreate

dbshow: ## Show tables
	docker exec -it $(appContainer) bash -c 'mysqlshow $(containerDbCred) $(filter-out $@,$(MAKECMDGOALS))'

dbschema: ## Dump mysql db
	docker exec -it $(appContainer) bash -c 'mysqldump ${containerDbCred}  ${DB_DATABASE}'

.PHONY: build-docker
build-docker: ## build the API server as a docker image
	docker build -f Dockerfile -t api-server .

.PHONY: clean
clean: ## remove temporary files
	rm -rf api-server coverage.out coverage-all.out

.PHONY: db-start
db-start: ## start the database server
	@mkdir -p testdata/mysql
	docker run --rm --name mysqlContainer -v $(shell pwd)/testdata:/testdata \
		-v $(shell pwd)/testdata/mysql:/var/lib/mysql/data \
		-e MYSQL_PASSWORD=mysql -e MYSQL_DB=ffxblue-articles -d -p 3306:3306 mysql

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop mysqlContainer

.PHONY: testdata
testdata: ## populate the database with test data
	make migrate-down
	@echo "Populating test data..."
	@docker exec -it mysqlContainer mysql "$(APP_DSN)" -f /testdata/testdata.sql

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
		$(MIGRATE) create -ext sql -dir /migrations/ $${name// /_}

# https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md#forcing-your-database-version
# https://github.com/dylanaraps/pure-bash-bible/blob/master/README.md#replacement
.PHONY: migrate-fix
migrate-fix: ## Fix dirty db
	@read -p "Enter the rollback version: " version; \
	$(MIGRATE) force $${version%%[^0-9]*} # remove non numeric chars using parameter expansion
