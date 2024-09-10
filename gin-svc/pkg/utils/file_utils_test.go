package utils

import (
	"fmt"
	"testing"
)

func TestFileMd5(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileMd5(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileMd5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileMd5() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetImageSize(t *testing.T) {
	got, got1, err := GetImageSize("/Users/wulinlin/workspace/github.com/wll/small-yello-book/gin-svc/static/1725640013_\U0001FAF5港星最帅之最！不接受反驳_4_杨平民的生活_来自小红书网页版.jpg")
	fmt.Println("丢雷", got1, got, err)
}
