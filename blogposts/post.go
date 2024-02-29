package blogposts

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

func readBody(scanner *bufio.Scanner) (body string) {
	scanner.Scan() // skip the ---

	buf := bytes.Buffer{}

	for scanner.Scan() { // Scan() return a boolean, until it done
		fmt.Fprintf(&buf, scanner.Text()+"\n")
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
