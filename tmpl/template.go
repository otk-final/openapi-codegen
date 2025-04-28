package tmpl

import (
	_ "embed"
	"errors"
	"text/template"
)

//go:embed api_ts.tmpl
var ts string

//go:embed api_go.tmpl
var golang string

//go:embed api_kotlin.tmpl
var kotlin string

//go:embed api_swift.tmpl
var swift string

//go:embed api_python.tmpl
var python string

//go:embed api_java.tmpl
var java string

var tmplFunc = template.FuncMap{
	"CapitalizeLetter": CapitalizeFirst,
	"Variable":         Variable,
}

func NewEngine(lang string, style string) (*template.Template, error) {
	tp := template.New(lang)

	tp = tp.Funcs(tmplFunc)

	if style != "" {
		return tp.ParseFiles(style)
	}

	switch lang {
	case "ts", "typescript":
		return tp.Parse(ts)
	case "go", "golang":
		return tp.Parse(golang)
	case "java":
		return tp.Parse(java)
	case "kotlin":
		return tp.Parse(kotlin)
	case "swift":
		return tp.Parse(swift)
	case "python":
		return tp.Parse(python)
	default:
		return nil, errors.New("invalid language")
	}
}
