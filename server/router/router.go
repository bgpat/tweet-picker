package router

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func Initialize(r *gin.Engine) {
	InitializeAPI(r.Group("/api"))
}
