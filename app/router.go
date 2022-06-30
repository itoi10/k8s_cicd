package main

import (
	"echo_sample/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	/* ミドルウェア */
	// アクセスログのようなリクエスト単位のログを出力する
	e.Use(middleware.Logger())
	// panicを起こしてもサーバを落とさずエラーレスポンスを返せるようにする
	e.Use(middleware.Recover())

	/* 静的ファイルを提供 */
	// GET /
	e.File("/", "public/index.html")

	/* 認証エンドポイント */
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	/* APIエンドポイント */
	api := e.Group("/api")
	// apiはJWT認証が必要
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.POST("/todos", handler.AddTodo)
	api.GET("/todos", handler.GetTodos)
	api.PUT("/todos/:id/completed", handler.UpdateTodo)
	api.DELETE("/todos/:id", handler.DeleteTodo)

	return e
}
