package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var DatabaseURL = "https://rg-push.firebaseio.com"
var CollectionName = "messages"

func main() {
	err := CreateMessage("/topics/rgru", "test golang message", "https://rg.ru", 5, "golang@rg.ru")
	if err != nil {
		log.Fatalln("Error :", err)
	}

	timeit("Push")
}

// functions ------------------------------------
func CreateMessage(to, message, link string, wait int, user string) error {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: DatabaseURL,
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("keys/rg-push-firebase-adminsdk-00.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return err
	}

	client, err := app.Database(ctx)
	if err != nil {
		return err
	}

	_, err = client.NewRef("posts").Push(ctx, map[string]interface{}{
		"to":      to,
		"message": message,
		"link":    link,
		"wait":    wait,
		"user":    user,
	})
	return err
}

// timing -----------------------------------------------------
var t0 = time.Now()
var t1 = time.Now()

func timeit(s string) {
	t1 = time.Now()
	fmt.Println(s, t1.Sub(t0))
	t0 = t1
}

// As an admin, the app has access to read and write all data, regradless of Security Rules
// reading --------------------------------------
// ref := client.NewRef("messages")
// var data map[string]interface{}
// if err := ref.Get(ctx, &data); err != nil {
// 	log.Fatalln("Error reading from database:", err)
// }
// timeit("Get")
// fmt.Println(data)

// writing ---------------------------------------
// Post is a json-serializable type.
