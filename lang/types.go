package lang

import (
	"cmp"
	"fmt"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

type GenericMode int

const (
	// ActualGenericMode 实参类型
	ActualGenericMode GenericMode = iota + 1

	// TypeGenericMode 行参类型
	TypeGenericMode
)

type TypeConvert interface {

	// Reference 引用类型
	Reference(name string) string
	// Foundation 基础类型
	Foundation(name string, format string) string
	// Array 数组
	Array(sub string) string
	// Map 字典
	Map(sub string) string
	// Generic 泛型
	Generic(parentType string, mode GenericMode, subTypes ...string) string
}

var handlers = map[string]TypeConvert{}

func init() {
	handlers["ts"] = &tsConvert{}
	handlers["swift"] = &swiftConvert{}
	handlers["java"] = &javaConvert{}
	handlers["go"] = &goConvert{}
	handlers["kotlin"] = &kotlinConvert{}
	handlers["python"] = &pythonConvert{}
}

func Names() []string {
	return lo.Keys(handlers)
}

func NewConvert(targetLang string, extendTypes map[string]string) TypeConvert {
	return &extendConverts{
		handler: handlers,
		mapping: extendTypes,
		target:  targetLang,
	}
}

type extendConverts struct {
	handler map[string]TypeConvert
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

func (l *extendConverts) Generic(parentType string, mode GenericMode, subTypes ...string) string {
	return l.handler[l.target].Generic(parentType, mode, subTypes...)
}

var reg = regexp.MustCompile(`\{([^/]+?)\}`)

func Format(lang string, path string, alias map[string]string) string {

	var segments []string
	text := reg.ReplaceAllStringFunc(path, func(seg string) string {

		segment := reg.FindStringSubmatch(seg)[1]
		segment = cmp.Or(alias[segment], segment)

		segments = append(segments, segment)

		switch lang {
		case "ts":
			return fmt.Sprintf("${%s}", segment)
		case "go":
			return "%v"
		case "swift":
			return "%@"
		default:
			return "%s"
		}
	})

	switch lang {
	case "ts":
		return fmt.Sprintf("`%s`", text)
	case "go":
		return fmt.Sprintf(`fmt.Sprintf("%s",%s)`, text, strings.Join(segments, ","))
	case "java":
		return fmt.Sprintf(`String.format("%s",%s)`, text, strings.Join(segments, ","))
	case "swift":
		return fmt.Sprintf(`String(format: "%s", %s)`, text, strings.Join(segments, ","))
	case "kotlin", "python":
		return fmt.Sprintf(`"%s".format(%s)`, text, strings.Join(segments, ","))
	}
	return text
}

type TextConvert interface {
	Placeholder(name string) string

	Format(name string, values ...string) string
}
