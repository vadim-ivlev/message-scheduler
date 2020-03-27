package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {
	t0 := time.Now()
	t1 := time.Now()
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://rg-push.firebaseio.com",
	}
	t1 = time.Now()
	println(t1.Sub(t0))
	t0 = t1

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("keys/rg-push-firebase-adminsdk-00.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	// reading --------------------------------------
	ref := client.NewRef("messages")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	// fmt.Println(data)

	// writing ---------------------------------------
	// Post is a json-serializable type.
	type Post struct {
		Author string `json:"author,omitempty"`
		Title  string `json:"title,omitempty"`
	}

	postsRef := client.NewRef("posts")

	// We can also chain the two calls together
	if _, err := postsRef.Push(ctx, &Post{
		Author: "alanisawesome",
		Title:  "The Turing Machine",
	}); err != nil {
		log.Fatalln("Error pushing child node:", err)
	}
	// END ------------------------------------------
	fmt.Println("hello world")
}
