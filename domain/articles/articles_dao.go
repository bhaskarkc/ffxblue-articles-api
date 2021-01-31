package articles

// Data Access Object

import (
	dbconn "github.com/bhaskarkc/ffxblue-article-api/datasource/mysql"
	"github.com/bhaskarkc/ffxblue-article-api/logger"
	"github.com/bhaskarkc/ffxblue-article-api/utils/errors"
	mysql_utils "github.com/bhaskarkc/ffxblue-article-api/utils/mysql"
)

const (
	insertSQL = `INSERT INTO articles ( title, body, date) VALUES(?, ?, ?);`
	getSQL    = `SELECT * from articles WHERE id=?`
)

func (article *Article) Get() *errors.RestErr {

	stmt, err := dbconn.Client.Prepare(getSQL)
	if err != nil {
		logger.Error("error when trying to prepare get article statment", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(article.Id)

	if err := result.Scan(
		&article.Id,
		&article.Title,
		&article.Body,
		&article.Date,
	); err != nil {
		logger.Error("error when trying to get user data", err)
		return mysql_utils.ParseErr(err)
	}
	return nil
}

func (article *Article) Save() *errors.RestErr {

	stmt, err := dbconn.Client.Prepare(insertSQL)
	if err != nil {
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
