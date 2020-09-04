package app

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"

	_ "blog/configs"
	"blog/global"
	"blog/pkg/util"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		appId     string
		appSecret string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "测试用例",
			args: args{
				appId:     "123",
				appSecret: "456",
			},
			want:    "123",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.appId, tt.args.appSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}

			fmt.Println(got)
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		token string
	}

	claimValue, _ := GenerateToken("202009041031300001", "456")
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{
		{
			name: "测试用例(错误)",
			args: args{
				token: "123",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "测试用例(正确)",
			args: args{
				token: claimValue,
			},
			want: &Claims{
				AppSecret: util.EnCodeMd5("456"),
				AppId:     "202009041031300001",
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(global.JwtSetting.Expire).Unix(),
					Issuer:    global.JwtSetting.Issuer,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJWTSecret(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{
			name: "测试用例正常",
			want: []byte(global.JwtSetting.Secret),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJWTSecret(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJWTSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
