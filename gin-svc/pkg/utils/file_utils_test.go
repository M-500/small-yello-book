package utils

import "testing"

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
