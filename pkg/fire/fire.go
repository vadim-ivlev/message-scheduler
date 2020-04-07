/*
Основная функциональность пакета.

Определяет функции для создания, мотификации, и удаления записей о новых сообщениях.

Функции образующие интерфейс GraphQL внесены в отдельный файл fire_graphql.go
*/

package fire

import (
	"context"
	"io/ioutil"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"
)

type appParams struct {
	DatabaseURL     string `yaml:"database_url"`
	CredentialsFile string `yaml:"credentials_file"`
	CollectionName  string `yaml:"collection_name"`
}

// Params параметры приложения
var Params appParams

// ReadConfig читает YAML
func ReadConfig(fileName string, env string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	envParams := make(map[string]appParams)
	err = yaml.Unmarshal(yamlFile, &envParams)
	if err != nil {
		panic(err)
	}
	Params = envParams[env]
	return
}

// createMessage Создает сообщение в коллекции сообщений Firebase.
// Имя коллекция задается параметрами определенными в configs/firebase.yaml.
// Входные параметры:
// to - Кому посылается сообщение. Токен пользователя или идентификатор топика в виде: /topics/topic_name
// message - Текст сообщения
// link - Ссылка в сообщении
// icon - Иконка сообщения
// wait - Время ожидания отсылки сообщения в минутах
// userEmail - Email пользователя который создал сообщение
func createMessage(to, message, link, icon string, wait int, userEmail string) error {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: Params.DatabaseURL,
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(Params.CredentialsFile)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return err
	}

	client, err := app.Database(ctx)
	if err != nil {
		return err
	}

	_, err = client.NewRef(Params.CollectionName).Push(ctx, map[string]interface{}{
		"to":      to,
		"message": message,
		"link":    link,
		"icon":    icon,
		"wait":    wait,
		"user":    userEmail,
	})
	return err
}

// updateMessage модифицирует сообщение в коллекции сообщений Firebase.
// Имя коллекция задается параметрами определенными в configs/firebase.yaml.
// Входные параметры:
// messageID - Идентификатор сообщения
// to - Кому посылается сообщение. Токен пользователя или идентификатор топика в виде: /topics/topic_name
// message - Текст сообщения
// link - Ссылка в сообщении
// icon - Иконка сообщения
// wait - Время ожидания отсылки сообщения в минутах
// userEmail - Email пользователя который создал сообщение
func updateMessage(messageID, to, message, link, icon string, wait int, userEmail string) error {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: Params.DatabaseURL,
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(Params.CredentialsFile)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return err
	}

	client, err := app.Database(ctx)
	if err != nil {
		return err
	}

	err = client.NewRef(Params.CollectionName).Child(messageID).Update(ctx, map[string]interface{}{
		"to":      to,
		"message": message,
		"link":    link,
		"icon":    icon,
		"wait":    wait,
		"user":    userEmail,
	})
	return err
}

// deleteMessage Удаляет сообщение из коллекции сообщений Firebase.
// Имя коллекция задается параметрами определенными в configs/firebase.yaml.
// параметры:
// messageID - Идентификатор сообщения
func deleteMessage(messageID string) error {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: Params.DatabaseURL,
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(Params.CredentialsFile)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return err
	}

	client, err := app.Database(ctx)
	if err != nil {
		return err
	}

	return client.NewRef(Params.CollectionName).Child(messageID).Delete(ctx)
}
