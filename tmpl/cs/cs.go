package cs

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

func (j *Convert) Reference(name string) string {
	return name
}

func (j *Convert) Foundation(name string, format string) string {
	switch name {
	case "integer", "number":
		if format == "int64" {
			return "long"
		}
		return "int"
	case "string":
		return "string"
	case "boolean":
		return "bool"
	default:
		return name
	}
}

func (j *Convert) Map(sub string) string {
	return fmt.Sprintf("Dictionary<string,%s>", sub)
}

func (j *Convert) Array(sub string) string {
	return fmt.Sprintf("List<%s>", sub)
}

func (j *Convert) Generic(parentType string, mode lang.GenericMode, subTypes ...string) string {
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
