package dart

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

func (j *Convert) Reference(name string) string {
	return name
}

func (j *Convert) Foundation(name string, format string) string {
	switch name {
	case "integer", "number":
		if format == "int64" {
			return "Long"
		}
		return "Integer"
	case "string":
		return "String"
	case "boolean":
		return "Boolean"
	case "object":
		return "JsonNode"
	default:
		return name
	}
}

func (j *Convert) Map(sub string) string {
	return fmt.Sprintf("Map<String,%s>", sub)
}

func (j *Convert) Array(sub string) string {
	return fmt.Sprintf("List<%s>", sub)
}

func (j *Convert) Generic(parentType string, mode tmpl.GenericMode, subTypes ...string) string {
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
