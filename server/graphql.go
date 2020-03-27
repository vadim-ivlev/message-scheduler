package server

import (
	"message-scheduler/pkg/fire"
	"message-scheduler/pkg/prometeo"

	"github.com/graphql-go/graphql"
)

// FUNCTIONS *******************************************************

// // jsonStringToMap преобразует строку JSON в map[string]interface{}
// func jsonStringToMap(s string) map[string]interface{} {
// 	m := make(map[string]interface{})
// 	_ = json.Unmarshal([]byte(s), &m)
// 	return m
// }

// // getParamsFromBody извлекает параметры запроса из тела запроса
// func getParamsFromBody(c *gin.Context) (map[string]interface{}, error) {
// 	r := c.Request
// 	mb := make(map[string]interface{})
// 	if r.ContentLength > 0 {
// 		errBodyDecode := json.NewDecoder(r.Body).Decode(&mb)
// 		return mb, errBodyDecode
// 	}
// 	return mb, errors.New("No body")
// }

// G R A P H Q L ********************************************************************************

var queryObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ping": &graphql.Field{
			Type:        graphql.String,
			Description: "quick test",
			Args:        graphql.FieldConfigArgument{},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return "pong", nil
			},
		},
	},
})

var mutationObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create_message": fire.CreateMessage(),
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryObject,
	Mutation: mutationObject,
})

// GraphQL исполняет GraphQL запрос
func GraphQL(payload *Payload) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  payload.Query,
		VariableValues: payload.Variables,
	})
	if len(result.Errors) > 0 {
		// инкрементируем счетчик ошибок GraphQL
		prometeo.GraphQLErrorsTotal.Inc()
	}
	return result
}
