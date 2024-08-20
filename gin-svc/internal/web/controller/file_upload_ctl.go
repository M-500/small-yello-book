package controller

import (
	"fmt"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type FileUploadController struct {
}

func NewFileUploadController() BaseController {
	return &FileUploadController{}
}

func (f *FileUploadController) UploadImgCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
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
	return ginx.JsonResult{Data: gin.H{"url": fileURL}}, nil
}

func (f *FileUploadController) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/file/upload", ginx.WrapResponse(f.UploadImgCtl))
}
