package response

import (
	"cms/model/common/merr"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间

	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(merr.YES.Code, map[string]interface{}{}, merr.YES.Msg, c)
}

func OkWithMessage(message string, c *gin.Context) {
	if message == "" {
		message = merr.YES.Msg
	}
	Result(merr.YES.Code, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(merr.YES.Code, data, merr.YES.Msg, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	if message == "" {
		message = merr.YES.Msg
	}
	Result(merr.YES.Code, data, message, c)
}

func Fail(c *gin.Context) {
	Result(merr.NO.Code, map[string]interface{}{}, merr.NO.Msg, c)
}

func FailWithMessage(message string, c *gin.Context) {
	if message == "" {
		message = merr.NO.Msg
	}
	Result(merr.NO.Code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	if message == "" {
		message = merr.NO.Msg
	}
	Result(merr.NO.Code, data, message, c)
}

func FailWithCodeDetailed(data interface{}, code int, message string, c *gin.Context) {
	if message == "" {
		message = merr.NO.Msg
	}
	Result(code, data, message, c)
}

func FailWithErr(errMsg merr.Msg, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		errMsg.Code,
		map[string]interface{}{},
		errMsg.Msg,
	})
}
