package controller

import (
	"errors"
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

func (n *NoteCtl) CreateNoteCtl(ctx *gin.Context, req types.CreateNoteForm, uc jwt.UserClaims) (result ginx.JsonResult, err error) {
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

func (n *NoteCtl) NoteListByUserPublished(ctx *gin.Context, req types.QueryNoteForm) (result ginx.JsonResult, err error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	userId := ctx.Param("uuid")
	if utils.IsBlank(userId) {
		return ginx.Error(400, "用户ID为空"), errors.New("用户ID为空")
	}
	publish, i, err := n.svc.ListNoteByUserPublish(ctx, userId, req.Size, req.Page)
	return ginx.SuccessPageList(publish, i), nil
}
func (n *NoteCtl) NoteListByUserCollected(ctx *gin.Context, req types.QueryNoteForm) (result ginx.JsonResult, err error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	userId := ctx.Param("uuid")
	if utils.IsBlank(userId) {
		return ginx.Error(400, "用户ID为空"), errors.New("用户ID为空")
	}
	return ginx.SuccessPageList(nil, 0), nil
}
func (n *NoteCtl) NoteListByUserLiked(ctx *gin.Context, req types.QueryNoteForm) (result ginx.JsonResult, err error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	userId := ctx.Param("uuid")
	if utils.IsBlank(userId) {
		return ginx.Error(400, "用户ID为空"), errors.New("用户ID为空")
	}
	return ginx.SuccessPageList(nil, 0), nil
}

// NoteDetail
//
//	@Description: 获取文章详情
//	@receiver n
//	@param ctx
//	@return result
//	@return err
func (n *NoteCtl) NoteDetail(ctx *gin.Context, claim jwt.UserClaims) (result ginx.JsonResult, err error) {
	uuid := ctx.Param("uuid")
	if utils.IsBlank(uuid) {
		return ginx.Error(400, "文章ID为空"), errors.New("文章ID为空")
	}
	note, err := n.svc.GetNoteDetail(ctx, claim.Uid, uuid)
	if err != nil {
		return ginx.Error(500, err.Error()), err
	}
	return ginx.SuccessJson(note), nil
}

func (n *NoteCtl) NaNoteDetail(ctx *gin.Context) (result ginx.JsonResult, err error) {
	uuid := ctx.Param("uuid")
	if utils.IsBlank(uuid) {
		return ginx.Error(400, "文章ID为空"), errors.New("文章ID为空")
	}
	note, err1 := n.svc.GetNoteDetail(ctx, "", uuid)
	if err1 != nil {
		return ginx.Error(500, err.Error()), err1
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
