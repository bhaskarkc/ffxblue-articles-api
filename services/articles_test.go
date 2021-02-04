package services

import (
	"reflect"
	"testing"

	"github.com/bhaskarkc/ffxblue-articles-api/domain/articles"
	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
)

func Test_articleService_Get(t *testing.T) {
	type args struct {
		Id int64
	}
	tests := []struct {
		name  string
		s     *articleService
		args  args
		want  *articles.Article
		want1 *errors.RestErr
	}{
		{
			"test",
			&articleService{},
			args{1},
			&articles.Article{},
			&errors.RestErr{
				Message: "database error",
				Status:  500,
				Error:   "internal_server_error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &articleService{}
			got, got1 := s.Get(tt.args.Id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("articleService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_articleService_Create(t *testing.T) {
	type args struct {
		article articles.Article
	}
	tests := []struct {
		name  string
		s     *articleService
		args  args
		want  *articles.Article
		want1 *errors.RestErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &articleService{}
			got, got1 := s.Create(tt.args.article)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("articleService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_articleService_RelatedTags(t *testing.T) {
	type args struct {
		Ids []int64
	}
	tests := []struct {
		name  string
		s     *articleService
		args  args
		want  []string
		want1 *errors.RestErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &articleService{}
			got, got1 := s.RelatedTags(tt.args.Ids)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleService.RelatedTags() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("articleService.RelatedTags() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
