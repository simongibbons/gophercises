package cyoa

import (
	"html/template"
	"net/http"
	"net/url"
)

type Handler struct {
	adventure CYOA
	tmpl      *template.Template
}

const (
	arcTemplate = `
<html>
<head>
<title> {{.Title}} </title>
<body>
<h1> {{.Title}} </h1>

{{ range .Story }}
    <p> {{.}} </p>
{{ end }}


{{ if .Options }}
<h3> Where next? </h3>
	{{ range .Options }}
    	<p> <a href="/{{.Arc}}"> {{.Text}} </a> </p>
	{{ end }}
{{ else }}
	<p> <a href="/intro">Start Again</a> </p>
{{ end }}
</body>
</html>
`

	arcTemplateName = "arc"
)

func NewHandler(adventure CYOA) Handler {
	tmpl, err := template.New(arcTemplateName).Parse(arcTemplate)
	if err != nil {
		panic(err)
	}
	return Handler{adventure: adventure, tmpl: tmpl}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := h.adventure.GetArc(getArcName(*r.URL))

	if arc == nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
	h.tmpl.Execute(w, arc)
}

func getArcName(u url.URL) string {
	if u.Path == "/" {
		return "intro"
	}
	return u.Path[1:]
}
