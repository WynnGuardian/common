package handlerfunc

import (
	"github.com/gin-gonic/gin"
	"github.com/wynnguardian/common/response"
)

type HandlerFunc func(ctx *gin.Context) response.WGResponse
