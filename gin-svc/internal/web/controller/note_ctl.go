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
	group.POST("/notes", ginx.WrapJsonBody[types.CreateNoteForm](n.CreateNoteCtl))
	group.GET("/notes", ginx.WrapQueryBody[types.QueryNoteForm](n.NoteListCtl))
}

func (n *NoteCtl) CreateNoteCtl(ctx *gin.Context, req types.CreateNoteForm) (result ginx.JsonResult, err error) {
	err = n.svc.CreateNote(ctx, req.ToNoteDomain())
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.Success(), nil
}

func (n *NoteCtl) NoteListCtl(ctx *gin.Context, req types.QueryNoteForm) (result ginx.JsonResult, err error) {
	notes, total, err := n.svc.ListNote(ctx, req.State, req.Page, req.Size)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.SuccessPageList(notes, total), nil
}
