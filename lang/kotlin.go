package lang

import (
	"fmt"
	"strings"
)

type kotlinConvert struct {
}

func (k *kotlinConvert) Reference(name string) string {
	return name
}

func (k *kotlinConvert) Foundation(name string, format string) string {
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

func (k *kotlinConvert) Map(sub string) string {
	return fmt.Sprintf("Map<String,%s>", sub)
}

func (k *kotlinConvert) Array(sub string) string {
	return fmt.Sprintf("List<%s>", sub)
}

func (k *kotlinConvert) Generic(parentType string, mode GenericMode, subTypes ...string) string {
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
