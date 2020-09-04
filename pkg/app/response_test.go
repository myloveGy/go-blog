package app

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"

	error2 "blog/pkg/error"

	_ "blog/configs"
)

func TestNewResponse(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			name: "测试用例一(正常)",
			args: args{
				ctx: &gin.Context{},
			},
			want: &Response{
				Ctx: &gin.Context{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponse(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_ToErrorResponse(t *testing.T) {

	type args struct {
		resRecorder *httptest.ResponseRecorder
		err         *error2.Error
	}

	tests := []struct {
		name string
		args args
		want string
		code int
	}{
		{
			name: "测试用例一",
			args: args{
				resRecorder: httptest.NewRecorder(),
				err:         error2.InvalidParams,
			},
			want: "{\"code\":40001,\"msg\":\"参数错误\"}",
			code: http.StatusBadRequest,
		},
		{
			name: "测试用例二",
			args: args{
				resRecorder: httptest.NewRecorder(),
				err:         error2.NotFound.WithDetails("页面不存在"),
			},
			want: "{\"code\":40004,\"details\":[\"页面不存在\"],\"msg\":\"Not Found\"}",
			code: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c1, _ := gin.CreateTestContext(tt.args.resRecorder)
			r := &Response{
				Ctx: c1,
			}

			r.ToErrorResponse(tt.args.err)

			got := tt.args.resRecorder.Body.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("r.ToErrorResponse() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.args.resRecorder.Code, tt.code) {
				t.Errorf("code error r.ToResponse() = %v, want %v", tt.args.resRecorder.Code, tt.code)
			}
		})
	}
}

func TestResponse_ToResponse(t *testing.T) {
	type args struct {
		resRecorder *httptest.ResponseRecorder
		data        interface{}
	}
	tests := []struct {
		name string
		args args
		want string
		code int
	}{
		{
			name: "测试用例一",
			args: args{
				resRecorder: httptest.NewRecorder(),
				data:        1234,
			},
			want: "1234",
			code: http.StatusOK,
		},
		{
			name: "测试用例二",
			args: args{
				resRecorder: httptest.NewRecorder(),
				data:        nil,
			},
			want: "{}",
			code: http.StatusOK,
		},
		{
			name: "测试用例三",
			args: args{
				resRecorder: httptest.NewRecorder(),
				data: map[string]interface{}{
					"username": "jinxing.liu",
					"age":      23,
				},
			},
			want: "{\"age\":23,\"username\":\"jinxing.liu\"}",
			code: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c1, _ := gin.CreateTestContext(tt.args.resRecorder)
			r := &Response{
				Ctx: c1,
			}

			r.ToResponse(tt.args.data)

			got := tt.args.resRecorder.Body.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("r.ToResponse() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.args.resRecorder.Code, tt.code) {
				t.Errorf("r.ToResponse() = %v, want %v", tt.args.resRecorder.Code, tt.code)
			}
		})
	}
}

func TestResponse_ToResponseList(t *testing.T) {
	type args struct {
		resRecorder *httptest.ResponseRecorder
		list        interface{}
		total       int
		url         string
	}
	tests := []struct {
		name string
		args args
		code int
		want string
	}{
		{
			name: "测试用例一",
			args: args{
				resRecorder: httptest.NewRecorder(),
				list:        []int{1, 2, 3, 4, 5},
				total:       100,
				url:         "http://example.com/?page=2",
			},
			code: http.StatusOK,
			want: "{\"list\":[1,2,3,4,5],\"pager\":{\"page\":2,\"page_size\":10,\"total\":100}}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c1, _ := gin.CreateTestContext(tt.args.resRecorder)
			c1.Request, _ = http.NewRequest("GET", tt.args.url, nil)
			r := &Response{Ctx: c1}
			r.ToResponseList(tt.args.list, tt.args.total)

			got := tt.args.resRecorder.Body.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("r.ToResponse() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.args.resRecorder.Code, tt.code) {
				t.Errorf("code error: r.ToResponse() = %v, want %v", tt.args.resRecorder.Code, tt.code)
			}
		})
	}
}
