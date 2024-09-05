package controller

import (
	"fmt"
	"gin-svc/internal/conf"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type FileController struct {
	cfg conf.ConfigInstance
}

func NewFileController(cfg conf.ConfigInstance) *FileController {
	return &FileController{
		cfg: cfg,
	}
}

func (f *FileController) UploadImgCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ginx.Error(http.StatusBadRequest, "Failed to get file"), err
	}
	// Create static directory if it doesn't exist
	staticDir := "static"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		os.Mkdir(staticDir, os.ModePerm)
	}
	// Save the file to the static directory
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filePath := filepath.Join(staticDir, filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		return ginx.Error(http.StatusInternalServerError, "Failed to save file"), err
	}

	// Return the file access URL
	fileURL := fmt.Sprintf("/%s/%s", staticDir, filename)
	url := f.cfg.Server.FileUploadHost + fileURL
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
func (f *FileController) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/file/upload", ginx.WrapResponse(f.UploadImgCtl))
	group.GET("/file", ginx.WrapResponse(f.ReadFileCtl))
}
