package service

import (
	"jtyl_feishu_mt/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func HandleError(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, model.Response{
		Code:    statusCode,
		Message: message,
		Error:   err.Error(),
	})
}
