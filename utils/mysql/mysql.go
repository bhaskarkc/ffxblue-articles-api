package mysql_utils

import (
	"strings"

	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseErr(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error processing request.")
		// return errors.NewInternalServerError(sqlErr.Error())
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicated key")
	}
	// return errors.NewInternalServerError(sqlErr.Error())
	return errors.NewInternalServerError("error processing request.")
}
