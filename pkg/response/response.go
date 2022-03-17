// Package response is gin response
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is a response struct
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	// Success is a success code
	Success = 0
	// Error is a error code
	Error = 7
)

// Result 返回结果封装
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// OK 返回成功（不带数据）
func OK(msg string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, c)
}

// Fail 返回失败（不带数据）
func Fail(msg string, c *gin.Context) {
	Result(Error, map[string]interface{}{}, msg, c)
}

// OKWithData 返回成功（带数据）
func OKWithData(data interface{}, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

// FailWithData 返回失败（带数据）
func FailWithData(data interface{}, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
