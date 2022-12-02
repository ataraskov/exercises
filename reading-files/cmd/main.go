package main

import (
	"blogposts"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	for i, p := range posts {
		log.Printf("[%d] %s\n", i, p)
	}
}
