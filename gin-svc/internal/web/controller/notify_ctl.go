package controller

import (
	"context"
	"gin-svc/internal/types"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/ginx"
)

func NewNotifyCtl() *NotifyCtl {
	return &NotifyCtl{}
}

type NotifyCtl struct {
}

func (n *NotifyCtl) NotifyListCtl(ctx *context.Context, query types.NotifyQueryForm, uc jwt.UserClaims) (result ginx.JsonResult, err error) {
	return ginx.Success(), nil
}
