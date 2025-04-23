package lang

import (
	"fmt"
	"strings"
)

type pythonConvert struct {
}

func (p *pythonConvert) Reference(name string) string {
	return name
}

func (p *pythonConvert) Foundation(name string, format string) string {
	switch name {
	case "integer", "number":
		return "int"
	case "boolean":
		return "bool"
	case "object":
		return "Any"
	case "string":
		return "str"
	default:
		return name
	}
}

func (p *pythonConvert) Map(sub string) string {
	return fmt.Sprintf("dict[str:%s]", sub)
}

func (p *pythonConvert) Array(sub string) string {
	return fmt.Sprintf("[%s]", sub)
}

func (p *pythonConvert) Generic(parentType string, mode GenericMode, subTypes ...string) string {
	if mode == TypeGenericMode {
		return fmt.Sprintf("%s(Generic[%s])", parentType, strings.Join(subTypes, ", "))
	}
	return fmt.Sprintf("%s[%s]", parentType, strings.Join(subTypes, ", "))
}
