package server

import (
	"github.com/gin-gonic/gin"
)

// HeadersMiddleware добавляет HTTP заголовки к ответу сервера
func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		origin := c.GetHeader("Origin")
		// fmt.Println("Origin:", origin, "c.Request.Host:", c.Request.Host)
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		// Если конечное приложение  не установило Access-Control-Allow-Credentials добавляем его
		if c.GetHeader("Access-Control-Allow-Credentials") == "" {
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		// c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Max-Age", "600")
		c.Next()
	}
}
