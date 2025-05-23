package ts

import (
	"codegen/lang"
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

func (t *Convert) Reference(name string) string {
	return name
}

func (t *Convert) Foundation(name string, format string) string {
	switch name {
	case "integer", "int", "long":
		return "number"
	case "object":
		return "any"
	default:
		return name
	}
}

func (t *Convert) Array(sub string) string {
	return fmt.Sprintf("%s[]", sub)
}

func (t *Convert) Map(sub string) string {
	return fmt.Sprintf("Record<string,%s>", sub)
}

func (t *Convert) Generic(parentType string, mode lang.GenericMode, subTypes ...string) string {
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
