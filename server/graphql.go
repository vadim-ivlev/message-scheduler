package server

import (
	"message-scheduler/pkg/fire"

	"github.com/graphql-go/graphql"
)

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
		"delete_message": fire.DeleteMessage(),
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryObject,
	Mutation: mutationObject,
})

// GraphQL исполняет GraphQL запрос
func doGraphQL(query string, variables map[string]interface{}) *graphql.Result {
	return graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: variables,
	})
}
