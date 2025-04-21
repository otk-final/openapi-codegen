package internal

import (
	"cmp"
	"codegen/lang"
	"codegen/tmpl"
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
)

type nestingGeneric struct {
	Kind       tmpl.NamedTypeKind
	Expression string
	Format     string
	Subs       []*nestingGeneric
}

func (n *nestingGeneric) discriminator(convert lang.TypeConvert) string {
	lines := make([]string, 0)
	for _, sub := range n.Subs {
		var expression = sub.Expression
		if sub.Kind&tmpl.GenericType != 0 {
			expression = sub.discriminator(convert)
		} else {
			expression = sub.Kind.Parse(expression, sub.Format, convert)
		}
		lines = append(lines, expression)
	}
	return convert.Generic(n.Expression, lang.ActualGenericMode, lines...)
}

type nestingGenericManage struct {
	generics []*tmpl.Ref
	refs     []*tmpl.Ref
}

// 递归
func (n *nestingGenericManage) recursion(current *nestingGeneric, generic, example *tmpl.Ref) {

	//同序 替换占位符
	genericProperties := generic.Properties
	slices.SortFunc(genericProperties, func(t1, t2 *tmpl.Property) int {
		return cmp.Compare(t1.Name, t2.Name)
	})
	exampleProperties := example.Properties
	slices.SortFunc(exampleProperties, func(t1, t2 *tmpl.Property) int {
		return cmp.Compare(t1.Name, t2.Name)
	})

	for idx, property := range genericProperties {

		//过滤非泛型属性
		if property.Type.Kind&tmpl.GenericType == 0 {
			continue
		}

		//同序 获取属性
		exampleProperty := exampleProperties[idx]
		nextKind := exampleProperty.Type.Kind
		nextFormat := exampleProperty.Format
		nextExpression := exampleProperty.Type.Expression

		//判断 属性类型是否为嵌套泛型
		nextGeneric, genericOk := lo.Find(n.generics, func(item *tmpl.Ref) bool {
			return strings.HasPrefix(nextExpression, item.Name)
		})

		//参考
		nextExample, exampleOk := lo.Find(n.refs, func(item *tmpl.Ref) bool {
			return item.Name == nextExpression
		})

		next := &nestingGeneric{
			Subs: make([]*nestingGeneric, 0),
		}

		if genericOk && exampleOk {
			//泛型
			next.Expression = nextGeneric.Name
			next.Kind = property.Type.Kind
			next.Format = property.Format

			n.recursion(next, nextGeneric, nextExample)
		} else {
			//基础属性类型
			next.Kind = nextKind

			//防止数组类型冲突，排除
			if property.Type.Kind&tmpl.ArrayType != 0 {
				next.Kind = nextKind ^ tmpl.ArrayType
			}
			next.Format = nextFormat
			next.Expression = nextExpression
		}
		current.Subs = append(current.Subs, next)
	}

}

var placeholders = []string{"T", "A", "B", "C", "D", "E", "F", "G", "H"}

// 匹配泛型
func resolvingGenerics(convert lang.TypeConvert, expressions map[string][]string, refs []*tmpl.Ref) []*tmpl.Ref {

	var defines = make([]*tmpl.Ref, 0)
	for wrapper, discriminators := range expressions {

		//查询已声明数据
		match, ok := lo.Find(refs, func(item *tmpl.Ref) bool {
			return strings.HasPrefix(item.Name, wrapper)
		})
		if !ok {
			continue
		}

		properties := make(tmpl.Properties, 0)
		idx := 0
		//替换泛型属性的 类型
		for _, property := range match.Properties {

			var expression string
			var kind tmpl.NamedTypeKind
			// 泛型  使用自定义的属性
			if lo.Contains(discriminators, property.Name) {
				//对象泛型
				expression = placeholders[idx]
				kind = tmpl.ImmutableType | tmpl.GenericType | property.Type.Kind //属性设置不可变
			} else if lo.Contains(discriminators, fmt.Sprintf("%s+", property.Name)) { //集合泛型
				//集合泛型
				expression = convert.Array(placeholders[idx])
				kind = tmpl.ImmutableType | tmpl.ArrayType | tmpl.GenericType | property.Type.Kind //属性设置不可变
			} else {
				//其他类型
				properties = append(properties, property)
				continue
			}

			//生成表达式
			property = &tmpl.Property{
				Name:        property.Name,
				Description: property.Description,
				Type: &tmpl.NamedType{
					Kind:       kind,
					Expression: expression,
				},
			}
			properties = append(properties, property)

			idx++
		}

		placeholders := placeholders[:idx]
		tp := &tmpl.NamedType{
			Kind:       tmpl.ImmutableType, //类型设置不可变
			Expression: convert.Generic(wrapper, lang.TypeGenericMode, placeholders...),
		}

		ref := &tmpl.Ref{
			Name:        wrapper,
			Type:        tp,
			Properties:  properties,
			Description: "",
		}
		defines = append(defines, ref)
	}

	return defines
}
