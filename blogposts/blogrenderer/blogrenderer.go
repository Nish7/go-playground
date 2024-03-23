package blogrenderer

import (
	"embed"
	"hello_world/blogposts/blogs"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (pr *PostRenderer) Render(w io.Writer, p blogs.Post) error { // why pointer to a postRenderer, not a instance? --> why pointer reciever was used because of mutability
	return pr.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (pr *PostRenderer) RenderIndex(w io.Writer, p []blogs.Post) error {
	return pr.templ.ExecuteTemplate(w, "index.gohtml", p)
}
