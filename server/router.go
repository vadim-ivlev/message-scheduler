package server

import (
	"message-scheduler/pkg/prometeo"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Setup присоединяет функции middleware.
func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default() // output to console
	r.Use(prometeo.CountersMiddleware())
	r.Use(HeadersMiddleware())
	return r
}

// AddRoutes сопоставляет маршруты функциям контроллера
func AddRoutes(r *gin.Engine) {
	// для префлайт запросов
	r.Handle("OPTIONS", "/schema", optionsHandler)
	// для GraphQL запросов
	r.Handle("POST", "/schema", graphqlHandler)
	// для метрик Прометея
	r.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	// для тестового приложения
	r.Static("/public", "./public")
}

// Serve запускает сервер на заданном порту.
func Serve(port string) {
	r := Setup()
	AddRoutes(r)
	_ = r.Run(port)
}
