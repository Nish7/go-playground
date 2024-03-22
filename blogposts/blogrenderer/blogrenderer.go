package blogrenderer

import (
	"hello_world/blogposts/blogs"
	"html/template"
	"io"
)

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(w io.Writer, p blogs.Post) error {
	templ, err := template.New("blogs").Parse(postTemplate)

	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return err
}
