package main

import (
	"hello_world/blogposts/blogs"
	"log"
	"os"
)

func main() {
	posts, err := blogs.NewPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
}
