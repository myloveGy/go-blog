package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"

	_ "blog/configs"
)

func TestBindAndValid(t *testing.T) {
	type args struct {
		c *gin.Context
		v interface{}
	}

	c1, _ := gin.CreateTestContext(httptest.NewRecorder())
	c1.Request, _ = http.NewRequest("GET", "http://example.com/?page=9", nil)

	c2, _ := gin.CreateTestContext(httptest.NewRecorder())
	c2.Request, _ = http.NewRequest("POST", "http://example.com", bytes.NewBufferString(`{"username":"123456"}`))

	c3, _ := gin.CreateTestContext(httptest.NewRecorder())
	c3.Request, _ = http.NewRequest("POST", "http://example.com", bytes.NewBufferString(`{"username":"123456"}`))
	c3.Request.Header.Add("Content-Type", "application/json; charset=utf-8")

	type typeA struct {
		Username string `json:"username" binding:"required"`
	}

	tests := []struct {
		name  string
		args  args
		want  bool
		want1 ValidErrors
	}{
		{
			name: "测试用例一",
			args: args{
				c: c1,
				v: nil,
			},
			want:  false,
			want1: nil,
		},
		{
			name: "测试用例二(验证不通过)",
			args: args{
				c: c2,
				v: &typeA{},
			},
			want: true,
			want1: ValidErrors{&ValidError{
				Key:     "typeA.Username",
				Message: "Key: 'typeA.Username' Error:Field validation for 'Username' failed on the 'required' tag",
			}},
		},
		{
			name: "测试用例三(验证通过)",
			args: args{
				c: c3,
				v: &typeA{},
			},
			want:  false,
			want1: nil,
		},
		{
			name: "测试用例四(验证错误)",
			args: args{
				c: c3,
				v: nil,
			},
			want:  true,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BindAndValid(tt.args.c, tt.args.v)
			if got != tt.want {
				t.Errorf("BindAndValid() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BindAndValid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestValidError_Error(t *testing.T) {
	type fields struct {
		Key     string
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "测试用例一",
			fields: fields{Key: "user_id", Message: "不能为空"},
			want:   "不能为空",
		},
		{
			name:   "测试用例二",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ValidError{
				Key:     tt.fields.Key,
				Message: tt.fields.Message,
			}
			if got := v.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidErrors_Error(t *testing.T) {
	tests := []struct {
		name string
		v    ValidErrors
		want string
	}{
		{
			name: "测试用例一",
			v:    ValidErrors{},
			want: "",
		},
		{
			name: "测试用例二",
			v: ValidErrors{
				&ValidError{Key: "user_id", Message: "不能为空"},
				&ValidError{Key: "username", Message: "不能少于两个字符"},
			},
			want: "不能为空,不能少于两个字符",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidErrors_Errors(t *testing.T) {
	tests := []struct {
		name string
		v    ValidErrors
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}
