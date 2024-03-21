package blogrenderer

import (
	"fmt"
	"hello_world/blogposts/blogs"
	"io"
)

func Render(w io.Writer, p blogs.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, p.Description)

	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, "Tags: <ul>")

	if err != nil {
		return nil
	}

	for _, tag := range p.Tags {
		_, err := fmt.Fprintf(w, "<li>%s</li>", tag)

		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(w, "</ul>")
	if err != nil {
		return nil
	}

	return err
}
