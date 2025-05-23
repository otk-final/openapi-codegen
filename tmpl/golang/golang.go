package golang

import (
	"codegen/lang"
	_ "embed"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

//go:embed api.tmpl
var templateApi string

//go:embed client.tmpl
var templateClient string

//go:embed model.tmpl
var templateModel string

var Templates = map[string]string{
	"api":    templateApi,
	"model":  templateModel,
	"client": templateClient,
}

type Convert struct {
}

func (g *Convert) Reference(name string) string {
	return fmt.Sprintf("*%s", name)
}

func (g *Convert) Foundation(name string, format string) string {
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

func (g *Convert) Map(sub string) string {
	return fmt.Sprintf("map[string]%s", sub)
}

func (g *Convert) Array(sub string) string {
	return fmt.Sprintf("[]%s", sub)
}

func (g *Convert) Generic(parentType string, mode lang.GenericMode, subTypes ...string) string {
	if mode == lang.TypeGenericMode {
		subTypes = lo.Map(subTypes, func(item string, index int) string {
			return fmt.Sprintf("%s any", item)
		})
	}
	return fmt.Sprintf("%s[%s]", parentType, strings.Join(subTypes, ","))
}
