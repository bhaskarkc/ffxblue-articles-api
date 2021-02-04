package tags

import (
	"fmt"
	"strconv"
	"strings"

	dbconn "github.com/bhaskarkc/ffxblue-articles-api/datasource/mysql"
	"github.com/bhaskarkc/ffxblue-articles-api/logger"
	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
	mysql_utils "github.com/bhaskarkc/ffxblue-articles-api/utils/mysql"
)

const (
	tagSelect = `SELECT * FROM tags WHERE name=?`

	tagInsertSQL = `INSERT INTO tags ( name ) VALUES(?);`

	tagRelInsertSQL = `
		INSERT INTO tag_relation
			( article_id, tag_id )
			VALUES(?, ?);`

	tagByDateSQL = `
		SELECT t.name as tag, count(r.article_id) as count, GROUP_CONCAT(r.article_id) as articles,
			(
				SELECT GROUP_CONCAT(DISTINCT tg.name)
				FROM tags as tg
				JOIN tag_relation as tr
					ON tr.tag_id = tg.id
				WHERE DATE_FORMAT(tr.date, "%Y%m%d") = ?
					AND tg.name != ?
			) as related_tags
		FROM tags as t
		JOIN tag_relation as r
			ON r.tag_id = t.id
		WHERE t.name = ?
			AND DATE_FORMAT(r.date, "%Y%m%d") = ?
		GROUP BY t.id;
	`
)

func (tag *Tag) Get() *errors.RestErr {
	stmt, err := dbconn.Client.Prepare(tagSelect)
	if err != nil {
		logger.Error("error when trying to prepare save tag statment", err)
		return mysql_utils.ParseErr(err)
	}
	defer stmt.Close()
	if err := stmt.QueryRow(tag.Name).Scan(&tag.Id, &tag.Name); err != nil {
		logger.Error("error: scan result: Tag", err)
		return mysql_utils.ParseErr(err)
	}
	return nil
}

// Save tag row
func (tag *Tag) Save() *errors.RestErr {
	stmt, err := dbconn.Client.Prepare(tagInsertSQL)
	if err != nil {
		logger.Error("error when trying to prepare save tag statment", err)
		return mysql_utils.ParseErr(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(tag.Name)
	if err != nil {
		return mysql_utils.ParseErr(err)
	}

	Id, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseErr(err)
	}
	tag.Id = Id
	return nil
}

// Save tag_relation row
func (tagr *TagRel) Save() *errors.RestErr {
	stmt, err := dbconn.Client.Prepare(tagRelInsertSQL)
	if err != nil {
		logger.Error("error: prepare statement: save tag_relation", err)
		return mysql_utils.ParseErr(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(tagr.ArticleId, tagr.TagId)
	if err != nil {
		return mysql_utils.ParseErr(err)
	}

	Id, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseErr(err)
	}
	tagr.Id = Id
	return nil
}

// List meta information of a specific tag published on a date.
func (tbd *TagByDate) TagByDate() *errors.RestErr {
	stmt, err := dbconn.Client.Prepare(tagByDateSQL)
	if err != nil {
		logger.Error("error: prepare statement: tagByDate", err)
		return mysql_utils.ParseErr(err)
	}
	defer stmt.Close()

	rows := stmt.QueryRow(
		tbd.Date,
		tbd.Tag.Name,
		tbd.Tag.Name,
		tbd.Date,
	)

	var relatedTags string
	var idCSV string
	if err := rows.Scan(
		&tbd.Tag.Name,
		&tbd.Count,
		&idCSV,
		&relatedTags,
	); err != nil {
		logger.Error("error: scan result: tagByDate", err)
		return mysql_utils.ParseErr(err)
	}
	tbd.RelatedTags = strings.Split(relatedTags, ",")

	var articleIds []int64
	for _, v := range strings.Split(idCSV, ",") {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		articleIds = append(articleIds, id)
	}
	tbd.Articles = articleIds
	fmt.Printf("%v\n", tbd)
	return nil
}
