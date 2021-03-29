package multiplexer

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRoute_ResolveURL(t *testing.T) {
	type fields struct {
		Pattern string
		Method  string
		Handler http.HandlerFunc
	}
	type args struct {
		urlPath string
	}

	params := []string{"1", "Ehab"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]*string
		wantErr bool
	}{
		{
			"basic",
			fields{
				Pattern: "test",
				Method:  "GET",
				Handler: nil,
			},
			args{urlPath: "test"},
			map[string]*string{},
			false,
		},
		{
			"one_param",
			fields{
				Pattern: "test/{id}",
				Method:  "GET",
				Handler: nil,
			},
			args{urlPath: "test/1"},
			map[string]*string{"id": &params[0]},
			false,
		},
		{
			"two_param",
			fields{
				Pattern: "test/{id}/{name}",
				Method:  "GET",
				Handler: nil,
			},
			args{urlPath: "test/1/Ehab"},
			map[string]*string{"id": &params[0], "name": &params[1]},
			false,
		},
		{
			"not_exists",
			fields{
				Pattern: "test/{id}/{name}",
				Method:  "GET",
				Handler: nil,
			},
			args{urlPath: "test/1/Ehab/test"},
			nil,
			true,
		},
		{
			"not_exists",
			fields{
				Pattern: "test/{id}/{name}",
				Method:  "GET",
				Handler: nil,
			},
			args{urlPath: "ggg/1/Ehab/"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Route{
				Pattern: tt.fields.Pattern,
				Method:  tt.fields.Method,
				Handler: tt.fields.Handler,
			}
			got, err := r.resolveURL(tt.args.urlPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("resolveURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
