package articles

// Data Access Object

import (
	"fmt"
	"strings"

	dbconn "github.com/bhaskarkc/ffxblue-articles-api/datasource/mysql"
	"github.com/bhaskarkc/ffxblue-articles-api/logger"
	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
	mysql_utils "github.com/bhaskarkc/ffxblue-articles-api/utils/mysql"
)

const (
	insertSQL = `INSERT INTO articles ( title, body, date) VALUES(?, ?, ?);`
	getSQL    = `SELECT articles.*, GROUP_CONCAT(tags.name) as tag FROM articles
					JOIN tag_relation as rel
						ON rel.article_id = articles.id
					JOIN tags
						ON tags.id = rel.tag_id
					WHERE articles.id=?`
)

func (article *Article) Get() *errors.RestErr {
	stmt, err := dbconn.Client.Prepare(getSQL)
	if err != nil {
		logger.Error("error when trying to prepare get article statment", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows := stmt.QueryRow(article.Id)

	var tagStr string
	if err := rows.Scan(
		&article.Id,
		&article.Title,
		&article.Body,
		&article.Date,
		&tagStr,
	); err != nil {
		logger.Error("error when trying to load article result", err)
		return mysql_utils.ParseErr(err)
	}

	article.Tags = strings.Split(tagStr, ",")
	return nil
}

func (article *Article) Save() *errors.RestErr {
	stmt, err := dbconn.Client.Prepare(insertSQL)
	if err != nil {
		fmt.Println(err)
		logger.Error("error when trying to prepare save article statment", err)
		return mysql_utils.ParseErr(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		article.Title,
		article.Body,
		article.Date,
	)

	if err != nil {
		return mysql_utils.ParseErr(err)
	}

	Id, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseErr(err)
	}
	article.Id = Id
	return nil
}
