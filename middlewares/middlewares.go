package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wynnguardian/common/handlerfunc"
	"github.com/wynnguardian/common/response"
)

func CheckOrigin(next handlerfunc.HandlerFunc) handlerfunc.HandlerFunc {
	return func(ctx *gin.Context) response.WGResponse {
		if ctx.GetHeader("Authorization") != os.Getenv("API_TOKEN") {
			return response.ErrUnauthorized
		}
		return next(ctx)
	}
}

func Authorize(next handlerfunc.HandlerFunc) handlerfunc.HandlerFunc {
	return func(ctx *gin.Context) response.WGResponse {
		return next(ctx)
	}
}

func Parse(next handlerfunc.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := next(ctx)
		if len(resp.Body) == 0 {
			resp.Body = "{}"
		}
		ctx.JSON(resp.Status, resp)
	}
}
