package tmpl

import (
	"codegen/lang"
	"github.com/samber/lo"
)

type Api struct {
	Name        string
	Description string
	Paths       []*Path
}

type Path struct {
	Tag          string
	Name         string
	Description  string
	Summary      string
	OriginalPath string
	Path         string
	Method       string
	Parameters   Parameters
	Queries      Parameters
	Request      *NamedType
	Response     *NamedType
}

type Parameter struct {
	Name        string
	Alias       string
	Required    bool
	In          string
	Type        *NamedType
	Format      string
	Description string
	Default     string
}

type NamedTypeKind uint

// 类型
const (
	ImmutableType NamedTypeKind = 1 << iota

	FoundationType

	MapType

	ArrayType

	ReferenceType

	GenericType
)

type NamedType struct {
	Kind       NamedTypeKind
	Expression string
}

type Parameters = []*Parameter

func (nt *NamedType) GenerateExpression(format string, convert lang.TypeConvert) {
	expression := nt.Expression

	//忽略转换
	if nt.Kind&ImmutableType != 0 {
		return
	}
	nt.Expression = nt.Kind.Parse(expression, format, convert)
}

func (nk NamedTypeKind) Parse(expression string, format string, convert lang.TypeConvert) string {

	if nk&FoundationType != 0 {
		expression = convert.Foundation(expression, format)
	}
	if nk&ReferenceType != 0 {
		expression = convert.Reference(expression)
	}
	if nk&ArrayType != 0 {
		expression = convert.Array(expression)
	}
	if nk&MapType != 0 {
		expression = convert.Map(expression)
	}
	return expression
}

type Ref struct {
	Name        string
	Alias       string
	Type        *NamedType
	Properties  Properties
	Description string
	Summary     string
	Ignore      bool
}

func (r *Ref) ReferenceLevel() int {
	//根据  ReferenceType 排序
	return lo.Reduce(r.Properties, func(agg int, item *Property, index int) int {
		if item.Type.Kind&ReferenceType != 0 {
			return agg + 1
		}
		return agg
	}, 0)

}

type Properties []*Property

func (p Properties) Find(name string) (*Property, bool) {
	return lo.Find(p, func(item *Property) bool {
		return item.Name == name
	})
}

type Generic struct {
	Expression string
	Property   string
}

type Property struct {
	Name        string
	Alias       string
	Description string
	Summary     string
	Type        *NamedType
	Format      string
	Enums       []string
}
