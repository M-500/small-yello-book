package controller

import (
	"fmt"
	"gin-svc/internal/conf"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type FileController struct {
	cfg     *conf.ConfigInstance
	mClient *minio.Client
}

func NewFileController(cfg *conf.ConfigInstance, client *minio.Client) *FileController {
	return &FileController{
		cfg:     cfg,
		mClient: client,
	}
}

func (f *FileController) UploadImgCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		return ginx.Error(http.StatusBadRequest, "Failed to get file"), err
	}
	// Create static directory if it doesn't exist
	bucketName := f.cfg.Minio.BucketName
	objectName := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)
	exists, err := f.mClient.BucketExists(ctx, bucketName)
	if err != nil {
		return ginx.Error(http.StatusInternalServerError, "Failed to check bucket"), err
	}
	if !exists {
		if err := f.mClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
			return ginx.Error(http.StatusInternalServerError, "Failed to create bucket"), err
		}
	}
	// 将文件上传到 MinIO
	_, err = f.mClient.PutObject(ctx, bucketName, objectName, file, header.Size, minio.PutObjectOptions{
		ContentType: header.Header.Get("Content-Type"),
	})
	// 上传成功，返回文件访问 URL
	url := fmt.Sprintf("http://%s/%s/%s", f.cfg.Minio.Endpoint, bucketName, objectName)
	return ginx.SuccessJson(url), nil
}

func (f *FileController) ReadFileCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
	filePath := ctx.Query("path")
	if filePath == "" {
		return ginx.Error(http.StatusBadRequest, "File path is required"), fmt.Errorf("file path is required")
	}

	// Ensure the file path is within the static directory
	staticDir := "static"
	fullPath := filepath.Join(staticDir, filepath.Clean(filePath))
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return ginx.Error(http.StatusNotFound, "File not found"), err
	}

	// Read the file content
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return ginx.Error(http.StatusInternalServerError, "Failed to read file"), err
	}

	return ginx.JsonResult{Data: gin.H{"content": string(content)}}, nil
}
