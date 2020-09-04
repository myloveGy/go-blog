package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	_ "blog/configs"
	"blog/global"
)

func TestGetPage(t *testing.T) {
	type args struct {
		c *gin.Context
	}

	c1, _ := gin.CreateTestContext(httptest.NewRecorder())
	c1.Request, _ = http.NewRequest("GET", "http://example.com/?page=9", nil)

	c2, _ := gin.CreateTestContext(httptest.NewRecorder())
	c2.Request, _ = http.NewRequest("GET", "http://example.com/?page=1000", nil)

	c3, _ := gin.CreateTestContext(httptest.NewRecorder())
	c3.Request, _ = http.NewRequest("GET", "http://example.com/?page=-10", nil)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试用例一",
			args: args{
				c: c1,
			},
			want: 9,
		},
		{
			name: "测试用例二",
			args: args{
				c: c2,
			},
			want: 1000,
		},
		{
			name: "测试用例三(不正常)",
			args: args{
				c: c3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPage(tt.args.c); got != tt.want {
				t.Errorf("GetPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPageOffset(t *testing.T) {
	type args struct {
		page int
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试用例一(正常)",
			args: args{
				page: 1,
				size: 10,
			},
			want: 0,
		},
		{
			name: "测试用例二(不正常)",
			args: args{
				page: -1,
				size: 10,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPageOffset(tt.args.page, tt.args.size); got != tt.want {
				t.Errorf("GetPageOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPageSize(t *testing.T) {
	type args struct {
		c *gin.Context
	}

	c1, _ := gin.CreateTestContext(httptest.NewRecorder())
	c1.Request, _ = http.NewRequest("GET", "http://example.com/?page_size=0", nil)

	c2, _ := gin.CreateTestContext(httptest.NewRecorder())
	c2.Request, _ = http.NewRequest("GET", "http://example.com/?page_size=1000", nil)

	c3, _ := gin.CreateTestContext(httptest.NewRecorder())
	c3.Request, _ = http.NewRequest("GET", "http://example.com/?page_size=20", nil)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试用例一(小于等于0)",
			args: args{
				c: c1,
			},
			want: global.AppSetting.DefaultPageSize,
		},
		{
			name: "测试用例二(大于最大值100)",
			args: args{
				c: c2,
			},
			want: global.AppSetting.MaxPageSize,
		},
		{
			name: "测试用例三(正常值20)",
			args: args{
				c: c3,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPageSize(tt.args.c); got != tt.want {
				t.Errorf("GetPageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
