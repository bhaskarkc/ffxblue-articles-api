package dbconn

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bhaskarkc/ffxblue-articles-api/utils"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&timeout=5s&tls=false&autocommit=true",
		utils.GetEnv("DB_USER", "root"),
		utils.GetEnv("DB_PASS", "password"),
		utils.JoinString(
			utils.GetEnv("DB_HOST", "localhost"), ":",
			utils.GetEnv("DB_PORT", "3306")),
		utils.GetEnv("DB_NAME", "ffxblue-articles"),
	)

	var err error

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	// We are reusing the Client therefore,
	// closing is not necessary. :)
	// defer Client.Close()

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("DB connection successful.")
}
