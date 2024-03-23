package blogs

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine("Title: "),
		Description: readMetaLine("Description: "),
		Tags:        strings.Split(readMetaLine("Tags: "), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	buf := bytes.Buffer{}

	scanner.Scan()

	for scanner.Scan() {
		fmt.Fprintf(&buf, scanner.Text()+"\n")
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
