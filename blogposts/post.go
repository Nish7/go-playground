package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFiIe io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFiIe)

	readline := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	titleLine := readline()[7:]
	descriptionLine := readline()[13:]

	return Post{Title: titleLine, Description: descriptionLine}, nil
}
