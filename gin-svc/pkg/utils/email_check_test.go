package utils

import "testing"

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空字符串",
			args: args{
				email: "",
			},
			want: false,
		},
		{
			name: "非法邮箱",
			args: args{
				email: "123",
			},
			want: false,
		},
		{
			name: "非法邮箱2",
			args: args{
				email: "123@qq.",
			},
			want: false,
		},
		{
			name: "合法邮箱",
			args: args{
				email: "18574945291@163.com",
			},
			want: true,
		},
		{
			name: "合法邮箱2",
			args: args{
				email: "18574945291@qq.com",
			},
			want: true,
		},
		{
			name: "合法邮箱2",
			args: args{
				email: "18574945291@gmail.com",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
