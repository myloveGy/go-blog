package upload

import (
	"mime/multipart"
	"os"
	"strings"
	"testing"

	_ "blog/configs"
)

func TestGetFileExt(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试正确",
			args: args{name: "my.txt"},
			want: ".txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileExt(tt.args.name); got != tt.want {
				t.Errorf("GetFileExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试正确",
			args: args{name: "my.txt"},
			want: "6864f389d9876436bc8778ff071d1b6c.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileName(tt.args.name); got != tt.want {
				t.Errorf("GetFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSavePath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "测试",
			want: "storage/uploads",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSavePath(); got != tt.want {
				t.Errorf("GetSavePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckSavePath(t *testing.T) {
	type args struct {
		dst string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试文件存在",
			args: args{"/go-project/blog"},
			want: false,
		},
		{
			name: "测试文件不存在",
			args: args{"/go-project/storage/upload"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSavePath(tt.args.dst); got != tt.want {
				t.Errorf("CheckSavePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckContainExt(t *testing.T) {
	type args struct {
		t    FileType
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试文件后缀存在",
			args: args{name: "my.jpeg", t: TypeImage},
			want: true,
		},
		{
			name: "测试文件后缀不存在",
			args: args{name: "my.txt", t: TypeImage},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckContainExt(tt.args.t, tt.args.name); got != tt.want {
				t.Errorf("CheckContainExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPermission(t *testing.T) {
	type args struct {
		dst string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试权限",
			args: args{
				dst: "/go-project/blog",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermission(tt.args.dst); got != tt.want {
				t.Errorf("CheckPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSavePath(t *testing.T) {
	type args struct {
		dst  string
		perm os.FileMode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试",
			args: args{
				dst:  "/go-project/blog/storage/upload",
				perm: os.ModePerm,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateSavePath(tt.args.dst, tt.args.perm); (err != nil) != tt.wantErr {
				t.Errorf("CreateSavePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	os.Remove("/go-project/blog/storage/upload")
}

func TestCheckMaxSize(t *testing.T) {
	type args struct {
		t FileType
		f multipart.File
	}

	f, _ := os.Open("/go-project/blog/storage/logs/app.log")
	defer f.Close()

	f1, _ := os.Open("/go-project/blog/testdata/test.jpg")
	defer f1.Close()

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试正常",
			args: args{
				t: TypeImage,
				f: f,
			},
			want: false,
		},
		{
			name: "测试文件超大",
			args: args{
				t: TypeImage,
				f: f1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckMaxSize(tt.args.t, tt.args.f); got != tt.want {
				t.Errorf("CheckMaxSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	filebContents = "Another test file."
)

const messageWithFileWithoutName = `
--MyBoundary
Content-Disposition: form-data; name="file"; filename="file1.txt"
Content-Type: text/plain

` + filebContents + `
--MyBoundary--
`

func TestSaveFile(t *testing.T) {
	type args struct {
		file *multipart.FileHeader
		dst  string
	}

	r := multipart.NewReader(strings.NewReader(messageWithFileWithoutName), "MyBoundary")
	form1, err := r.ReadForm(25)
	if err != nil {
		t.Fatal("ReadForm:", err)
	}
	defer form1.RemoveAll()

	// 删除测试文件
	_ = os.Remove("/go-project/blog/testdata/copy.txt")

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试正常",
			args: args{
				file: form1.File["file"][0],
				dst:  "/go-project/blog/testdata/copy.txt",
			},
			wantErr: false,
		},
		{
			name: "测试 multipart.FileHeader 文件错误",
			args: args{
				file: &multipart.FileHeader{Filename: "test.jpeg"},
				dst:  "/go-project/blog/testdata/copy.txt",
			},
			wantErr: true,
		},
		{
			name: "测试写入文件存在",
			args: args{
				file: form1.File["file"][0],
				dst:  "/go-project/blog/testdata1/test12.jpg",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveFile(tt.args.file, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("SaveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// 删除测试文件
	_ = os.Remove("/go-project/blog/testdata/copy.txt")
}
