package handlers

import (
	"LaraGo/data"
	"LaraGo/lara"
	"github.com/CloudyKit/jet/v6"
	"net/http"
)

type Handlers struct {
	App    *lara.Lara
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}

}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	testData := "SessionTestData"

	h.App.Session.Put(r.Context(), "test", testData)

	myValue := h.App.Session.GetString(r.Context(), "test")

	vars := make(jet.VarMap)
	vars.Set("test", myValue)

	err := h.App.Render.JetPage(w, r, "session", nil, vars)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering Jet Page: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error rendering Jet Page"))
		return
	}
}
