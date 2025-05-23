package lang

import (
	"cmp"
	"fmt"
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
		case "rust":
			return fmt.Sprintf("${%s}", segment)
		case "cs":
			return fmt.Sprintf("{%d}", len(segments)-1)
		default:
			return "%s"
		}
	})

	switch lang {
	case "ts":
		return fmt.Sprintf("`%s`", text)
	case "rust":
		return fmt.Sprintf(`format!("%s",%s)`, text, strings.Join(segments, ","))
	case "cs":
		return fmt.Sprintf(`string.Format("%s",%s)`, text, strings.Join(segments, ","))
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
