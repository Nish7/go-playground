package blogrenderer

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"hello_world/blogposts/blogs"
	"io"
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

	postRenderer, err := NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it convertes a single post to html", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err = postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}

		posts := []blogs.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err = postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogs.Post{
			Title:       "hello, world",
			Body:        "This is a post",
			Description: "this is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
