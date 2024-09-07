package controller

import (
	"errors"
	"fmt"
	"gin-svc/internal/conf"
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/ginx"
	"gin-svc/pkg/utils"
	"github.com/gin-gonic/gin"
)

type NoteCtl struct {
	svc service.NoteService
	cfg *conf.ConfigInstance
}

func NewNoteCtl(svc service.NoteService, cfg *conf.ConfigInstance) *NoteCtl {
	return &NoteCtl{
		svc: svc,
		cfg: cfg,
	}
}

func (n *NoteCtl) RegisterRoute(group *gin.RouterGroup) {
	group.POST("/notes", ginx.WrapJsonBodyAndClaims[types.CreateNoteForm, jwt.UserClaims](n.CreateNoteCtl))
	group.PUT("/notes/:uuid/pass", ginx.WrapResponse(n.PassNoteCtl))
	group.GET("/notes", ginx.WrapQueryBody[types.QueryNoteForm](n.NoteListCtl))
	group.GET("/notes/detail/:uuid", ginx.WrapResponse(n.NoteDetail))                        // 获取文章详情信息
	group.GET("/feed/notes", ginx.WrapQueryBody[types.FeedNoteQueryForm](n.FeedNoteListCtl)) // 获取推荐文章列表  后续要改成feed流模式
}

func (n *NoteCtl) CreateNoteCtl(ctx *gin.Context, req types.CreateNoteForm, uc jwt.UserClaims) (result ginx.JsonResult, err error) {
	fmt.Println(uc)
	err = n.svc.CreateNote(ctx, req.ToNoteDomain(n.cfg.Server), uc.Uid)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.Success(), nil
}

func (n *NoteCtl) PassNoteCtl(ctx *gin.Context) (result ginx.JsonResult, err error) {
	uuid := ctx.Param("uuid")
	if utils.IsBlank(uuid) {
		return ginx.Error(400, "文章ID为空"), errors.New("文章ID为空")
	}
	err = n.svc.PassNote(ctx, uuid)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.Success(), nil
}

func (n *NoteCtl) NoteListCtl(ctx *gin.Context, req types.QueryNoteForm) (result ginx.JsonResult, err error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	notes, total, err := n.svc.ListNote(ctx, req.State, req.Page, req.Size)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.SuccessPageList(notes, total), nil
}

// NoteDetail
//
//	@Description: 获取文章详情
//	@receiver n
//	@param ctx
//	@return result
//	@return err
func (n *NoteCtl) NoteDetail(ctx *gin.Context) (result ginx.JsonResult, err error) {
	uuid := ctx.Param("uuid")
	if utils.IsBlank(uuid) {
		return ginx.Error(400, "文章ID为空"), errors.New("文章ID为空")
	}
	note, err := n.svc.GetNoteDetail(ctx, uuid)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.SuccessJson(note), nil
}

func (n *NoteCtl) FeedNoteListCtl(ctx *gin.Context, req types.FeedNoteQueryForm) (result ginx.JsonResult, err error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	notes, err := n.svc.FeedListNote(ctx, req.TagId, req.Page, req.Size)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.SuccessPageList(notes, 0), nil
}
