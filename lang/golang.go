package lang

import (
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type goConvert struct {
}

func (g *goConvert) Reference(name string) string {
	//return fmt.Sprintf("*%s", util.CapitalizeFirst(name))
	return name
}

func (g *goConvert) Foundation(name string, format string) string {
	switch name {
	case "integer", "number":
		if format != "" {
			return format
		}
		return "int"
	case "boolean":
		return "bool"
	case "object":
		return "interface{}"
	default:
		return name
	}
}

func (g *goConvert) Array(sub string) string {
	return fmt.Sprintf("[]%s", sub)
}

func (g *goConvert) Generic(parentType string, mode GenericMode, subTypes ...string) string {
	if mode == TypeGenericMode {
		subTypes = lo.Map(subTypes, func(item string, index int) string {
			return fmt.Sprintf("%s any", item)
		})
	}
	return fmt.Sprintf("%s[%s]", parentType, strings.Join(subTypes, ","))
}
