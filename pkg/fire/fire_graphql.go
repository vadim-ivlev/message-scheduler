/*
GraphQL интерфейс к функциям пакета.

Этот файл содержит функции файла fire.go упакованые согласно требованиям GraphQL.
*/

package fire

import (
	"github.com/graphql-go/graphql"
)

// CreateMessage добавляет сообщение в таблицу сообщений Firebase
func CreateMessage() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "добавить новое сообщение",
		Args: graphql.FieldConfigArgument{
			"to": &graphql.ArgumentConfig{
				Type:         graphql.NewNonNull(graphql.String),
				Description:  "Токен пользователя или имя топика в виде '/topics/topicName'",
				DefaultValue: "/topics/rgru",
			},
			"message": &graphql.ArgumentConfig{
				Type:         graphql.String,
				Description:  "Текст сообщения",
				DefaultValue: "Тестовое сообщение",
			},
			"link": &graphql.ArgumentConfig{
				Type:         graphql.String,
				Description:  "Ссылка",
				DefaultValue: "https://rg.ru",
			},
			"icon": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "URL Иконки сообщения",
				// DefaultValue: "https://rg.ru/favicon.ico",
				DefaultValue: "",
			},
			"wait": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				Description:  "Задержка сообщения в минутах",
				DefaultValue: 5,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			to := p.Args["to"].(string)
			message := p.Args["message"].(string)
			link := p.Args["link"].(string)
			wait := p.Args["wait"].(int)
			icon := p.Args["icon"].(string)
			user := "golang@rg.ru"
			err := createMessage(to, message, link, icon, wait, user)
			if err != nil {
				return "error", err
			}
			return "ok", nil
		},
	}
}

// UpdateMessage изменяет сообщение в таблице сообщений Firebase
func UpdateMessage() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "добавить новое сообщение",
		Args: graphql.FieldConfigArgument{
			"message_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Идентификатор сообщения",
			},
			"to": &graphql.ArgumentConfig{
				Type:         graphql.NewNonNull(graphql.String),
				Description:  "Токен пользователя или имя топика в виде '/topics/topicName'",
				DefaultValue: "/topics/rgru",
			},
			"message": &graphql.ArgumentConfig{
				Type:         graphql.String,
				Description:  "Текст сообщения",
				DefaultValue: "Тестовое сообщение",
			},
			"link": &graphql.ArgumentConfig{
				Type:         graphql.String,
				Description:  "Ссылка",
				DefaultValue: "https://rg.ru",
			},
			"icon": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "URL Иконки сообщения",
				// DefaultValue: "https://rg.ru/favicon.ico",
				DefaultValue: "",
			},
			"wait": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				Description:  "Задержка сообщения в минутах",
				DefaultValue: 5,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			messageID := p.Args["message_id"].(string)
			to := p.Args["to"].(string)
			message := p.Args["message"].(string)
			link := p.Args["link"].(string)
			icon := p.Args["icon"].(string)
			wait := p.Args["wait"].(int)
			user := "golang@rg.ru"
			err := updateMessage(messageID, to, message, link, icon, wait, user)
			if err != nil {
				return "error", err
			}
			return "ok", nil
		},
	}
}

// DeleteMessage удаляет сообщение по его идентификатору
func DeleteMessage() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Удалить сообщение",
		Args: graphql.FieldConfigArgument{
			"message_id": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Идентификатор сообщения",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			messageID := p.Args["message_id"].(string)
			err := deleteMessage(messageID)
			if err != nil {
				return "error", err
			}
			return "ok", nil
		},
	}
}
