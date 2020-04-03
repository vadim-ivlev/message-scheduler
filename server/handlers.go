package server

import (
	"message-scheduler/pkg/prometeo"

	"github.com/gin-gonic/gin"
)

// optionsHandler Для проверки работоспособности.
func optionsHandler(c *gin.Context) {
	c.JSON(200, "")
}

// graphqlHandler Исполняет GraphQL запрос
func graphqlHandler(c *gin.Context) {
	query, variables := GetQueryAndVariables(c)
	result := doGraphQL(query, variables)

	// инкрементируем счетчик ошибок GraphQL
	if len(result.Errors) > 0 {
		prometeo.GraphQLErrorsTotal.Inc()
	}
	c.JSON(200, result)
}
