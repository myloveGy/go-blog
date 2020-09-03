package util

import (
	"fmt"
	"testing"
)

func TestEnCodeMd5(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试正确",
			args: args{
				value: "123232323",
			},
			want: "94cf58545fb0aab010e39422dfd27892",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnCodeMd5(tt.args.value); got != tt.want {
				t.Errorf("EnCodeMd5() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}
