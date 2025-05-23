package python

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

func (p *Convert) Reference(name string) string {
	return name
}

func (p *Convert) Foundation(name string, format string) string {
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

func (p *Convert) Map(sub string) string {
	return fmt.Sprintf("Dict[str,%s]", sub)
}

func (p *Convert) Array(sub string) string {
	return fmt.Sprintf("List[%s]", sub)
}

func (p *Convert) Generic(parentType string, mode lang.GenericMode, subTypes ...string) string {
	if mode == lang.TypeGenericMode {
		return fmt.Sprintf("%s(Generic[%s])", parentType, strings.Join(subTypes, ", "))
	}
	return fmt.Sprintf("%s[%s]", parentType, strings.Join(subTypes, ", "))
}
