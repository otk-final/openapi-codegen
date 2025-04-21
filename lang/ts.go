package lang

import (
	"fmt"
	"strings"
)

type tsConvert struct {
}

func (t *tsConvert) Reference(name string) string {
	return name
}

func (t *tsConvert) Foundation(name string, format string) string {
	switch name {
	case "integer", "int", "long":
		return "number"
	case "object":
		return "any"
	default:
		return name
	}
}

func (t *tsConvert) Array(sub string) string {
	return fmt.Sprintf("%s[]", sub)
}

func (t *tsConvert) Generic(parentType string, mode GenericMode, subTypes ...string) string {
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
