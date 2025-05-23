package swift

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

func (s *Convert) Reference(name string) string {
	return name
}

func (s *Convert) Foundation(name string, format string) string {
	switch name {
	case "integer", "number":
		if format == "int64" {
			return "Int64"
		}
		return "Int"
	case "string":
		return "String"
	case "boolean":
		return "Bool"
	default:
		return name
	}
}

func (s *Convert) Array(sub string) string {
	return fmt.Sprintf("[%s]", sub)
}

func (s *Convert) Map(sub string) string {
	return fmt.Sprintf("[String:%s]", sub)
}

func (s *Convert) Generic(parentType string, mode lang.GenericMode, subTypes ...string) string {
	if mode == lang.TypeGenericMode {
		subTypes = lo.Map(subTypes, func(item string, index int) string {
			return fmt.Sprintf("%s: Codable", item)
		})
	}
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
