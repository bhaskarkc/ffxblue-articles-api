package services

import (
	"fmt"

	"github.com/bhaskarkc/ffxblue-articles-api/domain/articles"
	"github.com/bhaskarkc/ffxblue-articles-api/domain/tags"
	"github.com/bhaskarkc/ffxblue-articles-api/utils/date"
	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
)

var ArticleService articleServiceInterface = &articleService{}

type articleServiceInterface interface {
	Get(int64) (*articles.Article, *errors.RestErr)
	Create(articles.Article) (*articles.Article, *errors.RestErr)
	RelatedTags([]int64) ([]string, *errors.RestErr)
}

type articleService struct{}

func (s *articleService) Get(Id int64) (*articles.Article, *errors.RestErr) {
	article := &articles.Article{Id: Id}
	article.Date = date.GetNowString()
	if err := article.Get(); err != nil {
		return nil, err
	}
	return article, nil
}

func (s *articleService) Create(article articles.Article) (*articles.Article, *errors.RestErr) {
	if err := article.Validate(); err != nil {
		return nil, err
	}

	article.Date = date.GetNowDBFormat()
	if err := article.Save(); err != nil {
		return nil, err
	}

	var tag = &tags.Tag{}
	var err = &errors.RestErr{}
	for _, tagName := range article.Tags {
		tag.Name = tagName
		if !tag.Exists() {
			tag, err = TagService.Create(*tag)
			fmt.Println(tag)
			if err != nil {
				// TODO: remove tagName from article struct
				continue
			}
		}

		tagRelation := tags.TagRel{
			ArticleId: article.Id,
			TagId:     tag.Id,
			Date:      date.GetNowString(),
		}
		fmt.Println(tagRelation)

		if _, err := TagService.CreateTagRelation(tagRelation); err != nil {
			return nil, err
		}
	}
	return &article, nil
}

func (s *articleService) RelatedTags(Ids []int64) ([]string, *errors.RestErr) {
	return []string{}, nil
}
