package articles

import (
	"reflect"
	"testing"

	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
)

func TestArticle_Validate(t *testing.T) {
	type fields struct {
		Id    int64
		Title string
		Date  string
		Body  string
		Tags  []string
	}
	tests := []struct {
		name   string
		fields fields
		want   *errors.RestErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			article := &Article{
				Id:    tt.fields.Id,
				Title: tt.fields.Title,
				Date:  tt.fields.Date,
				Body:  tt.fields.Body,
				Tags:  tt.fields.Tags,
			}
			if got := article.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Article.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
