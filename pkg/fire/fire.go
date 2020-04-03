package fire

import (
	"context"
	"io/ioutil"

	firebase "firebase.google.com/go"
	"github.com/graphql-go/graphql"
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
			user := "golang@rg.ru"
			err := createMessage(to, message, link, wait, user)
			if err != nil {
				return "error", err
			}
			return "ok", nil
		},
	}
}

func createMessage(to, message, link string, wait int, user string) error {
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
		"wait":    wait,
		"user":    user,
	})
	return err
}
