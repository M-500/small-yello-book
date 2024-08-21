package controller

import (
	"fmt"
	"gin-svc/internal/types"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type NoteCtl struct {
}

func (n *NoteCtl) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/note", ginx.WrapJsonBody[types.CreateNoteForm](n.CreateNoteCtl))
}

func NewNoteCtl() *NoteCtl {
	return &NoteCtl{}
}

func (n *NoteCtl) CreateNoteCtl(ctx *gin.Context, req types.CreateNoteForm) (result ginx.JsonResult, err error) {
	fmt.Println(req)
	return ginx.JsonResult{}, nil
}
