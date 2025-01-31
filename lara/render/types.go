package render

import "github.com/CloudyKit/jet/v6"

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	GenericData     map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}
