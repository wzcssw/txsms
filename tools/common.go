package tools

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/shesuyo/yunpian"
)

// YunpianKey YunpianKey
var YunpianKey = "b92ee0b369f4749409bdb76bd2d13776"

// APIInstance APIInstance
var APIInstance *yunpian.YunpianAPI

// FloatToString  to convert a float number to a string
func FloatToString(inputnum float64) string {
	return strconv.FormatFloat(inputnum, 'f', -1, 64)
}

// ReturnJSONResult 通用返回结果
func ReturnJSONResult(c echo.Context, success bool, msg string, e error) error {
	msgHead := ""
	if success {
		msgHead = "发送成功"
	} else {
		if e == nil {
			msgHead = "发送失败:" + msg
		} else {
			msgHead = "发送失败:" + e.Error()
		}
	}
	return c.JSON(http.StatusOK, &map[string]interface{}{"success": success, "msg": msgHead})
}
