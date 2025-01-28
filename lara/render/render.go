package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func (ren *Render) Page(w http.ResponseWriter, r *http.Request,
	view string, variables, data interface{}) error {
	switch strings.ToLower(ren.Renderer) {
	case "go":
		return ren.GoPage(w, r, view, data)
	case "jet":

	}

	return nil
}

func (ren *Render) GoPage(w http.ResponseWriter, r *http.Request,
	view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl",
		ren.RootPath, view))
	if err != nil {
		return err
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, td)
	if err != nil {
		return err
	}

	return nil
}
