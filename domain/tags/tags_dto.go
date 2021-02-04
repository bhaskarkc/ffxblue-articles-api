package tags

import (
	"time"

	errors "github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
)

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"tag"`
}

type TagRel struct {
	Id        int64
	ArticleId int64
	TagId     int64
	Date      string
}

type TagByDate struct {
	Tag
	Count       int64    `json:"count"`
	Articles    []int64  `json:"articles"`
	RelatedTags []string `json:"related_tags"`
	Date        string   `json:"-"`
}

func (tag *Tag) Validate() *errors.RestErr {
	if tag.Name == "" {
		return errors.NewBadRequestError("tag name can not be empty")
	}
	return nil
}

func (tag *Tag) Exists() bool {
	if err := tag.Get(); err != nil {
		return false
	}
	return true
}

func (tagRelation *TagRel) Validate() *errors.RestErr {
	if tagRelation.ArticleId < 1 {
		return errors.NewBadRequestError("ArticleId for tag relation can not be empty")
	}

	if tagRelation.TagId < 1 {
		return errors.NewBadRequestError("TagId for tag relation can not empty")
	}
	return nil
}

func (tagByDate *TagByDate) Validate() *errors.RestErr {
	// https://golang.org/pkg/time/#Parse
	_, err := time.Parse("20060102", tagByDate.Date)
	if err != nil {
		return errors.NewBadRequestError("Date in URI param is invalid")
	}
	return nil
}
