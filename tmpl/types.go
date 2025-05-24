package tmpl

import (
	"codegen/lang"
	"codegen/tmpl/cs"
	"codegen/tmpl/golang"
	"codegen/tmpl/java"
	"codegen/tmpl/kotlin"
	"codegen/tmpl/python"
	"codegen/tmpl/rust"
	"codegen/tmpl/swift"
	"codegen/tmpl/ts"
	"fmt"
	"github.com/samber/lo"
)

var handlers = map[string]lang.TypeConvert{}

func init() {
	handlers["ts"] = &ts.Convert{}
	handlers["swift"] = &swift.Convert{}
	handlers["java"] = &java.Convert{}
	handlers["go"] = &golang.Convert{}
	handlers["kotlin"] = &kotlin.Convert{}
	handlers["python"] = &python.Convert{}
	handlers["cs"] = &cs.Convert{}
	handlers["rust"] = &rust.Convert{}
}

func LangNames() []string {
	return lo.Keys(handlers)
}

func NewLangConvert(targetLang string, extendTypes map[string]string) lang.TypeConvert {
	return &extendConverts{
		handler: handlers,
		mapping: extendTypes,
		target:  targetLang,
	}
}

type extendConverts struct {
	handler map[string]lang.TypeConvert
	mapping map[string]string
	target  string
}

func (l *extendConverts) Reference(name string) string {
	v, ok := l.mapping[name]
	if ok {
		return v
	}
	return l.handler[l.target].Reference(name)
}

func (l *extendConverts) Foundation(name string, format string) string {

	v, ok := l.mapping[name]
	if ok {
		return v
	}

	//组合类型 type + format
	combination := fmt.Sprintf("%s+%s", name, format)
	v, ok = l.mapping[combination]
	if ok {
		return v
	}

	return l.handler[l.target].Foundation(name, format)
}

func (l *extendConverts) Array(sub string) string {
	return l.handler[l.target].Array(sub)
}

func (l *extendConverts) Map(sub string) string {
	return l.handler[l.target].Map(sub)
}

func (l *extendConverts) Generic(parentType string, mode lang.GenericMode, subTypes ...string) string {
	return l.handler[l.target].Generic(parentType, mode, subTypes...)
}
