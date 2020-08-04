package convert

import (
	"testing"
)

func TestStrTo_Int(t *testing.T) {
	tests := []struct {
		name    string
		s       StrTo
		want    int
		wantErr bool
	}{
		{
			name:    "测试01",
			s:       StrTo("012"),
			want:    12,
			wantErr: false,
		},
		{
			name:    "测试abc",
			s:       StrTo("abc"),
			want:    0,
			wantErr: true,
		},
		{
			name:    "测试111",
			s:       StrTo("111"),
			want:    111,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Int() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrTo_MustInt(t *testing.T) {
	tests := []struct {
		name string
		s    StrTo
		want int
	}{
		{
			name: "测试abc",
			s:    StrTo("abc"),
			want: 0,
		},
		{
			name: "测试001",
			s:    StrTo("001"),
			want: 1,
		},
		{
			name: "测试11",
			s:    StrTo("11"),
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MustInt(); got != tt.want {
				t.Errorf("MustInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrTo_MustUInt32(t *testing.T) {
	tests := []struct {
		name string
		s    StrTo
		want uint32
	}{
		{
			name: "测试abc",
			s:    StrTo("abc"),
			want: 0,
		},
		{
			name: "测试001",
			s:    StrTo("001"),
			want: 1,
		},
		{
			name: "测试-1",
			s:    StrTo("-1"),
			want: 4294967295,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MustUInt32(); got != tt.want {
				t.Errorf("MustUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrTo_String(t *testing.T) {
	tests := []struct {
		name string
		s    StrTo
		want string
	}{
		{
			name: "测试abc",
			s:    StrTo("abc"),
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrTo_UInt32(t *testing.T) {
	tests := []struct {
		name    string
		s       StrTo
		want    uint32
		wantErr bool
	}{
		{
			name:    "测试abc",
			s:       StrTo("abc"),
			want:    0,
			wantErr: true,
		},
		{
			name:    "测试123",
			s:       StrTo("123"),
			want:    123,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UInt32()
			if (err != nil) != tt.wantErr {
				t.Errorf("UInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}
