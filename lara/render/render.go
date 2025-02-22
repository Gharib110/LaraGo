package render

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func (ren *Render) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Secure = ren.Secure
	td.ServerName = ren.ServerName
	td.Port = ren.Port

	if ren.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	return td
}

// Page renders a page based on the renderer in Render struct
func (ren *Render) Page(w http.ResponseWriter, r *http.Request,
	view string, variables, data interface{}) error {
	switch strings.ToLower(ren.Renderer) {
	case "go":
		return ren.GoPage(w, r, view, data)
	case "jet":
		return ren.JetPage(w, r, view, variables, data)
	}

	return nil
}

// GoPage renders a go template
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

// JetPage renders a jet template
func (ren *Render) JetPage(w http.ResponseWriter, r *http.Request, templateName string,
	variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	td = ren.defaultData(td, r)

	t, err := ren.JetViews.GetTemplate(templateName + ".jet")
	if err != nil {
		log.Println(err)
		return err
	}

	err = t.Execute(w, vars, td)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
