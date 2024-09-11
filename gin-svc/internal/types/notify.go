package types

type NotifyQueryForm struct {
	Page       int    `form:"page" binding:"required"`
	PerPage    int    `form:"per_page" binding:"required"`
	NotifyType string `form:"notify_type" binding:"required"`
}
