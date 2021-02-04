package services

import (
	"reflect"
	"testing"

	"github.com/bhaskarkc/ffxblue-articles-api/domain/tags"
	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
)

func Test_tagService_Create(t *testing.T) {
	type args struct {
		tag tags.Tag
	}
	tests := []struct {
		name  string
		t     *tagService
		args  args
		want  *tags.Tag
		want1 *errors.RestErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &tagService{}
			got, got1 := t.Create(tt.args.tag)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagService.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("tagService.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_tagService_CreateTagRelation(t *testing.T) {
	type args struct {
		tagRelation tags.TagRel
	}
	tests := []struct {
		name  string
		t     *tagService
		args  args
		want  *tags.TagRel
		want1 *errors.RestErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &tagService{}
			got, got1 := t.CreateTagRelation(tt.args.tagRelation)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagService.CreateTagRelation() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("tagService.CreateTagRelation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_tagService_GetTagByDate(t *testing.T) {
	type args struct {
		tagName    string
		dateString string
	}
	tests := []struct {
		name  string
		t     *tagService
		args  args
		want  *tags.TagByDate
		want1 *errors.RestErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &tagService{}
			got, got1 := t.GetTagByDate(tt.args.tagName, tt.args.dateString)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagService.GetTagByDate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("tagService.GetTagByDate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
