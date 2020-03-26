package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://rg-push.firebaseio.com",
	}
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
	ref := client.NewRef("messages")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)

	fmt.Println("hello world")

	// Post is a json-serializable type.
	type Post struct {
		Author string `json:"author,omitempty"`
		Title  string `json:"title,omitempty"`
	}

	postsRef := ref.Child("posts")

	// newPostRef, err := postsRef.Push(ctx, nil)
	// if err != nil {
	// 	log.Fatalln("Error pushing child node:", err)
	// }

	// if err := newPostRef.Set(ctx, &Post{
	// 	Author: "gracehop",
	// 	Title:  "Announcing COBOL, a New Programming Language",
	// }); err != nil {
	// 	log.Fatalln("Error setting value:", err)
	// }

	// We can also chain the two calls together
	if _, err := postsRef.Push(ctx, &Post{
		Author: "alanisawesome",
		Title:  "The Turing Machine",
	}); err != nil {
		log.Fatalln("Error pushing child node:", err)
	}

}
