package tmpl

import (
	"codegen/tmpl/cpp"
	"codegen/tmpl/cs"
	"codegen/tmpl/dart"
	"codegen/tmpl/golang"
	"codegen/tmpl/java"
	"codegen/tmpl/kotlin"
	"codegen/tmpl/python"
	"codegen/tmpl/rust"
	"codegen/tmpl/swift"
	"codegen/tmpl/ts"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"unicode"
)

var register = map[string]map[string]string{
	"ts":     ts.Templates,
	"go":     golang.Templates,
	"swift":  swift.Templates,
	"python": python.Templates,
	"kotlin": kotlin.Templates,
	"java":   java.Templates,
	"cs":     cpp.Templates,  //TODO
	"cpp":    cs.Templates,   //TODO
	"rust":   rust.Templates, //TODO
	"dart":   dart.Templates, //TODO
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func NewEngine(lang string, name string, style string, variable map[string]string) (*template.Template, error) {
	tp := template.New(lang)

	var tmplFunc = template.FuncMap{
		"Capitalize": Capitalize,
		"Variable": func(key string) string {
			return variable[key]
		},
	}
	tp.Funcs(tmplFunc)

	//自定义模版
	if style != "" {
		return tp.ParseFiles(PwdJoinPath(style))
	}

	mapping, ok := register[lang]
	if !ok {
		return nil, errors.New("invalid language")
	}
	define, ok := mapping[name]
	if !ok {
		return nil, errors.New("invalid language template name")
	}
	return tp.Parse(define)
}

func PwdJoinPath(name string) string {
	if filepath.IsAbs(name) {
		return filepath.Clean(name)
	}
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, name)
}

var langFileSuffix = map[string]string{
	"python": "py",
	"golang": "go",
	"kotlin": "kt",
	"rust":   "rs",
}

func NewOutputs(lang string) map[string]*Output {

	suffix, ok := langFileSuffix[lang]
	if !ok {
		suffix = lang
	}

	return map[string]*Output{
		"api": {
			Header:    []string{},
			File:      fmt.Sprintf("api.%s", suffix),
			Variables: map[string]string{},
		},
		"model": {
			Header:    []string{},
			File:      fmt.Sprintf("model.%s", suffix),
			Variables: map[string]string{},
		},
		"client": {
			Header:    []string{},
			File:      fmt.Sprintf("client.%s", suffix),
			Variables: map[string]string{},
		},
	}
}
