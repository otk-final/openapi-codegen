package lang

import (
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type swiftConvert struct {
}

func (s *swiftConvert) Reference(name string) string {
	return name
}

func (s *swiftConvert) Foundation(name string, format string) string {
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

func (s *swiftConvert) Array(sub string) string {
	return fmt.Sprintf("[%s]", sub)
}

func (s *swiftConvert) Map(sub string) string {
	return fmt.Sprintf("[String:%s]", sub)
}

func (s *swiftConvert) Generic(parentType string, mode GenericMode, subTypes ...string) string {
	if mode == TypeGenericMode {
		subTypes = lo.Map(subTypes, func(item string, index int) string {
			return fmt.Sprintf("%s: Codable", item)
		})
	}
	return fmt.Sprintf("%s<%s>", parentType, strings.Join(subTypes, ","))
}
