package services

import (
	"fmt"

	"github.com/bhaskarkc/ffxblue-articles-api/domain/tags"
	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
)

var TagService tagServiceInterface = &tagService{}

type tagServiceInterface interface {
	Create(tags.Tag) (*tags.Tag, *errors.RestErr)
	CreateTagRelation(tags.TagRel) (*tags.TagRel, *errors.RestErr)
	GetTagByDate(string, string) (*tags.TagByDate, *errors.RestErr)
}

type tagService struct{}

func (t *tagService) Create(tag tags.Tag) (*tags.Tag, *errors.RestErr) {
	if err := tag.Validate(); err != nil {
		return nil, err
	}

	if err := tag.Save(); err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t *tagService) CreateTagRelation(tagRelation tags.TagRel) (*tags.TagRel, *errors.RestErr) {
	if err := tagRelation.Validate(); err != nil {
		return nil, err
	}
	if err := tagRelation.Save(); err != nil {
		return nil, err
	}
	return &tagRelation, nil
}

func (t *tagService) GetTagByDate(tagName string, dateString string) (*tags.TagByDate, *errors.RestErr) {
	tagByDate := &tags.TagByDate{
		Tag:  tags.Tag{Name: tagName},
		Date: dateString,
	}

	if !tagByDate.Tag.Exists() {
		return nil, errors.NewBadRequestError(fmt.Sprintf("Tag '%s' does not exists", tagName))
	}

	if err := tagByDate.Validate(); err != nil {
		return nil, err
	}

	if err := tagByDate.TagByDate(); err != nil {
		return nil, err
	}
	return tagByDate, nil
}
