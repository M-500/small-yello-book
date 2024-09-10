package utils

import (
	"crypto/md5"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func FileMd5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to compute md5: %w", err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func GetFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func GetImageSize(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()
	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}
	// 获取图片的宽度和高度
	bounds := img.Bounds()
	width := bounds.Dx()  // 图片宽度
	height := bounds.Dy() // 图片高度

	return width, height, nil
}
