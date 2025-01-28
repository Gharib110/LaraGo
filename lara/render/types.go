package render

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
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
