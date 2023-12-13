package contentful

import (
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func Test_contentfulAPI_urlFactory(t *testing.T) {
	type args struct {
		entryID string
	}
	tests := []struct {
		name    string
		api     contentfulAPI
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Successful case",
			api:  contentfulAPI{},
			args: args{
				entryID: "testtest",
			},
			want:    "https://cdn.contentful.com/spaces/xxxxxxxxxxxx/entries/testtest?access_token=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := contentfulAPI{}
			got, err := a.urlFactory(tt.args.entryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("contentfulAPI.urlFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("contentfulAPI.urlFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contentfulAPI_Get(t *testing.T) {
	type args struct {
		entryID string
	}
	tests := []struct {
		name       string
		api        contentfulAPI
		args       args
		statusCode int
		wantErr    bool
	}{
		{
			name: "Successful case",
			api:  contentfulAPI{},
			args: args{
				entryID: os.Getenv("CONTENTFUL_ENTRY_ID"),
			},
			statusCode: 200,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := contentfulAPI{}
			got, err := a.Get(tt.args.entryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("contentfulAPI.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.StatusCode, tt.statusCode) {
				t.Errorf("contentfulAPI.Get() = %v, want %v", got, tt.statusCode)
			}
		})
	}
}

func Test_contentfulAPI_JSONParse(t *testing.T) {
	api := NewContentfulAPIFactory()
	entryID := os.Getenv("CONTENTFUL_ENTRY_ID")
	res, err := api.Get(entryID)
	if err != nil || res.StatusCode != 200 {
		log.Fatalln(err)
		return
	}
	defer res.Body.Close()

	type args struct {
		res       *http.Response
		jqPattern []string
	}
	tests := []struct {
		name    string
		api     contentfulAPI
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Successful case",
			api:  contentfulAPI{},
			args: args{
				res: res,
				jqPattern: []string{
					".sys.id",
					".fields.name",
					".sys.createdAt",
				},
			},
			want: []string{
				entryID,
				"蜂蜜豆乳クランベリー",
				"2023-07-06T08:32:22.090Z",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := contentfulAPI{}
			got, err := a.JSONParse(tt.args.res, tt.args.jqPattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("contentfulAPI.JSONParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("contentfulAPI.JSONParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
