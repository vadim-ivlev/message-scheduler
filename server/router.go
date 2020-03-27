package server

import (
	"message-scheduler/pkg/prometeo"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// AddRoutes сопоставляет маршруты функциям контроллера
func AddRoutes(r *gin.Engine) {
	r.Handle("OPTIONS", "/schema", optionsHandler)
	r.Handle("POST", "/schema", grahpqlHandler)
	r.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	// Статика нужна для проверки (на проде не используется)
	r.Static("/public", "./public")
}

// Setup присоединяет функции middleware.
func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default() // output to console
	r.Use(prometeo.CountersMiddleware())
	r.Use(HeadersMiddleware())
	return r
}

// Serve запускает сервер на заданном порту.
func Serve(port string) {
	r := Setup()
	AddRoutes(r)
	_ = r.Run(port)
}
