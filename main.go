package main

import (
	"net/http"
	"txsms/apis"
	"txsms/tools"

	"github.com/labstack/echo"
	"github.com/shesuyo/yunpian"
)

/*
	Next:
		1.依赖管理
		2.完善日志
*/

func main() {
	e := echo.New()
	tools.APIInstance = yunpian.NewYunpianAPI(tools.YunpianKey)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, txsms!")
	})

	// 使用模板发送
	e.POST("/template", apis.Template)

	// 启动服务
	e.Logger.Fatal(e.Start(":3000"))
}
