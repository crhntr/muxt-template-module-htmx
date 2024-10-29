package hypertext

import (
	"context"
	"embed"
	"html/template"
)

//go:embed templates/*.gohtml
var templateFiles embed.FS

var templates = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

type Server struct{}

type IndexData struct {
	Name string
}

func (srv *Server) Index(_ context.Context) IndexData {
	return IndexData{
		Name: "friend",
	}
}
