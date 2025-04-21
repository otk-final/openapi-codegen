package lang

import (
	"fmt"
	"strings"
)

type javaConvert struct {
}

func (j *javaConvert) Reference(name string) string {
	return name
}

func (j *javaConvert) Foundation(name string, format string) string {
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

func (j *javaConvert) Array(sub string) string {
	return fmt.Sprintf("List<%s>", sub)
}

func (j *javaConvert) Generic(parentType string, mode GenericMode, subTypes ...string) string {
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
