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

//go:embed api_dart.tmpl
var dart string

//go:embed api_java.tmpl
var java string

//go:embed api_cpp.tmpl
var cpp string

//go:embed api_rust.tmpl
var rust string

//go:embed api_oc.tmpl
var oc string

var tmplFunc = template.FuncMap{
	"CapitalizeLetter": CapitalizeFirst,
	"Variable":         Variable,
}

// NewEngine 模版
func NewEngine(lang string, style string) (*template.Template, error) {
	tp := template.New(lang)

	//扩展方法
	tp = tp.Funcs(tmplFunc)

	//自定义模版
	if style != "" {
		return tp.ParseFiles(style)
	}

	switch lang {
	case "ts", "js":
		return tp.Parse(ts)
	case "go", "golang":
		return tp.Parse(golang)
	case "c++", "cpp":
		return tp.Parse(cpp)
	case "java":
		return tp.Parse(java)
	case "kotlin":
		return tp.Parse(kotlin)
	case "swift":
		return tp.Parse(swift)
	case "python":
		return tp.Parse(python)
	case "dart":
		return tp.Parse(dart)
	case "oc":
		return tp.Parse(oc)
	case "rust":
		return tp.Parse(rust)
	default:
		return nil, errors.New("invalid language")
	}
}
