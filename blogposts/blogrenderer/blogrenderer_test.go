package blogrenderer

import (
	"bytes"
	"hello_world/blogposts/blogs"
	"testing"
)

func TestRenderer(t *testing.T) {
	var (
		aPost = blogs.Post{
			Title:       "hello, world",
			Body:        "This is a post",
			Description: "This is a Description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it convertes a single post to html", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello, world</h1><p>This is a Description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`

		if got != want {
			t.Errorf("got %s wanted %s", got, want)
		}
	})
}
