package controllers

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func ReturnErrors(c *gin.Context, errs []error) bool {
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, errs)
		return true
	}
	return false
}
