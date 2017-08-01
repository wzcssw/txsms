package main

import (
	"net/http"
	"txsms/apis"

	"github.com/labstack/echo"
)

/*
	Next:
		1.依赖管理
		2.完善日志
*/

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, txsms!")
	})

	// 使用模板发送
	e.POST("/template", apis.Template)

	// 启动服务
	e.Logger.Fatal(e.Start(":3000"))
}
