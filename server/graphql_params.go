package server

import (
	"errors"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// getParamsFromBody извлекает параметры запроса из тела запроса
func getParamsFromBody(c *gin.Context) (map[string]interface{}, error) {
	r := c.Request
	params := make(map[string]interface{})
	if r.ContentLength > 0 {
		err := json.NewDecoder(r.Body).Decode(&params)
		return params, err
	}
	return params, errors.New("No body")
}

// GetQueryAndVariables извлекает "query", "variables"
// Обрабатывает данные как из полей форм так и из тела запроса.
// Данные полученные из тела запроса перекрывают данные POST, которые перекрывают данные GET.
func GetQueryAndVariables(c *gin.Context) (query string, variables map[string]interface{}) {

	// Берем query и variables из Form Data
	query = c.PostForm("query")
	variables = make(map[string]interface{})
	_ = json.Unmarshal([]byte(c.PostForm("variables")), &variables)

	// если есть тело запроса то берем query и variables из Request Payload (для Altair)
	params, err := getParamsFromBody(c)
	if err == nil {
		query, _ = params["query"].(string)
		variables, _ = params["variables"].(map[string]interface{})
	}

	return
}
