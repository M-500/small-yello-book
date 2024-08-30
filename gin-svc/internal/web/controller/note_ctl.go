package controller

import (
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type NoteCtl struct {
	svc service.NoteService
}

func NewNoteCtl(svc service.NoteService) *NoteCtl {
	return &NoteCtl{svc: svc}
}

func (n *NoteCtl) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/note", ginx.WrapJsonBody[types.CreateNoteForm](n.CreateNoteCtl))
}

func (n *NoteCtl) CreateNoteCtl(ctx *gin.Context, req types.CreateNoteForm) (result ginx.JsonResult, err error) {
	err = n.svc.CreateNote(ctx, req.ToNoteDomain())
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.Success(), nil
}
