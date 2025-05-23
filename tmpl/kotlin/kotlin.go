package kotlin

import (
	"codegen/tmpl"
	_ "embed"
	"fmt"
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

func (k *Convert) Reference(name string) string {
	return name
}

func (k *Convert) Foundation(name string, format string) string {
	switch name {
	case "integer", "number":
		if format == "int64" {
			return "Long"
		}
		return "Int"
	case "string":
		return "String"
	case "boolean":
		return "Boolean"
	default:
		return name
	}
}

func (k *Convert) Map(sub string) string {
	return fmt.Sprintf("Map<String,%s>", sub)
}

func (k *Convert) Array(sub string) string {
	return fmt.Sprintf("List<%s>", sub)
}

func (k *Convert) Generic(parentType string, mode tmpl.GenericMode, subTypes ...string) string {
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
