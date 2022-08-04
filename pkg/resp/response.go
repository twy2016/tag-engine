package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code  int         `json:"code,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg,omitempty"`
	Error string      `json:"error,omitempty"`
}

func SendResponseInvalidParameter(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, Response{
		Code:  http.StatusBadRequest,
		Msg:   "参数错误",
		Error: err.Error(),
	})
}

func DataBaseFail(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:  http.StatusInternalServerError,
		Msg:   "数据库操作出错",
		Error: err.Error(),
	})
}

func SendResponse(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: data,
		Msg:  msg,
	})
}
