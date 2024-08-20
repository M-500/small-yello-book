package ginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WrapJsonBody[Req any](fn func(*gin.Context, Req) (JsonResult, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, JsonResult{Code: 400, Msg: "bad request", Data: nil})
			return
		}
		res, err := fn(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func WrapQueryBody[Req any](fn func(*gin.Context, Req) (JsonResult, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, JsonResult{Code: 400, Msg: "bad request", Data: nil})
			return
		}
		res, err := fn(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func WrapResponse(fn func(*gin.Context) (JsonResult, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := fn(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func WrapJWT[T any](fn func(ctx *gin.Context, claims T) (JsonResult, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, ok := ctx.Get("UserId")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		uc, ok := val.(T)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		res, err := fn(ctx, uc)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}
