package internal

import (
	"cmp"
	"codegen/lang"
	"codegen/tmpl"
	v2 "codegen/v2"
	v3 "codegen/v3"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"slices"
	"strings"
	"text/template"
)

type Args struct {
	Name         string `json:"name"`
	Endpoint     string `json:"endpoint"`
	Output       string `json:"output"`
	ClientOutput string `json:"client_output"`
	Lang         string `json:"lang"`
	Style        string `json:"style"`
	Version      string `json:"version"`
}

type Env struct {
	*Args
	Ignore                []string          `json:"ignore"`
	Filter                []string          `json:"filter"`
	Alias                 Alias             `json:"alias"`
	Variables             map[string]string `json:"variables"`
	Generics              *Generics         `json:"generics"`
	RepeatableOperationId bool              `json:"repeatable_operation_id"`
}

type Alias struct {
	Properties map[string]string `json:"properties"`
	Modes      map[string]string `json:"modes"`
	Types      map[string]string `json:"types"`
	Parameters map[string]string `json:"parameters"`
}

type Generics struct {
	Enable      bool                `json:"enable"`
	Unfold      bool                `json:"unfold"`
	Expressions map[string][]string `json:"expressions"`
}

type Executor struct {
	env     *Env
	convert lang.TypeConvert
	engine  *template.Template
}

func New(env *Env, args []string) (*Executor, error) {

	tp, err := tmpl.NewEngine(env.Lang, env.Style)
	if err != nil {
		return nil, err
	}

	return &Executor{
		env:    env,
		engine: tp,
	}, nil
}

func (e *Executor) Run(cmd *Args) error {

	LANG := cmd.Lang
	typeConvert := lang.NewConvert(LANG, e.env.Alias.Types)
	e.convert = typeConvert

	var outRefs []*tmpl.Ref
	var outApis []*tmpl.Api
	var err error

	if e.env.Version == "v2" {
		outRefs, outApis, err = v2.LoadParse(e.env.Endpoint)
	} else if e.env.Version == "v3" {
		outRefs, outApis, err = v3.LoadParse(e.env.Endpoint)
	} else {
		return errors.New("invalid version")
	}

	if err != nil {
		return err
	}

	// 过滤 path
	filter := e.env.Filter
	ignore := e.env.Ignore
	lo.ForEach(outApis, func(item *tmpl.Api, index int) {
		item.Paths = lo.Filter(item.Paths, func(p *tmpl.Path, index int) bool {
			path := p.Path
			_, ignoreMatch := lo.Find(ignore, func(p string) bool {
				return strings.HasPrefix(path, p)
			})
			if ignoreMatch {
				return false
			}

			if len(filter) > 0 {
				_, filterMatch := lo.Find(filter, func(p string) bool {
					return strings.HasPrefix(path, p)
				})
				return filterMatch
			}
			return true
		})
	})
	outApis = lo.Filter(outApis, func(item *tmpl.Api, index int) bool {
		return len(item.Paths) > 0
	})

	slices.SortFunc(outRefs, func(a, b *tmpl.Ref) int {
		return cmp.Compare(a.Name, b.Name)
	})

	slices.SortFunc(outApis, func(a, b *tmpl.Api) int {
		return cmp.Compare(a.Name, b.Name)
	})

	//是否需要生成泛型
	if e.env.Generics.Enable {

		expressions := e.env.Generics.Expressions

		generics := resolvingGenerics(e.convert, expressions, outRefs)

		mgr := &nestingGenericManage{
			generics: generics,
			refs:     outRefs,
		}

		newRefs := make([]*tmpl.Ref, 0)
		newRefs = append(newRefs, generics...)

		for _, ref := range outRefs {

			refName := ref.Name
			//是否为泛型，前缀是否一致
			matched, ok := lo.Find(generics, func(item *tmpl.Ref) bool {
				return strings.HasPrefix(refName, item.Name)
			})

			if ok {
				//递归识别泛型
				val := &nestingGeneric{
					Kind:       tmpl.GenericType,
					Expression: matched.Name,
					Subs:       make([]*nestingGeneric, 0),
				}
				mgr.recursion(val, matched, ref)

				//是否展开泛型，只获取子类型，方便业务直接使用
				if e.env.Generics.Unfold {
					val.unfold()
				}

				//设置泛型>不可变类型
				ref.Type.Kind = tmpl.GenericType | tmpl.ImmutableType
				ref.Type.Expression = val.discriminator(typeConvert)
			}

			newRefs = append(newRefs, ref)
		}
		outRefs = newRefs
	}

	//根据类型 生成属性表达式
	for _, ref := range outRefs {

		//排序
		properties := ref.Properties
		slices.SortFunc(properties, func(a, b *tmpl.Property) int {
			return cmp.Compare(a.Name, b.Name)
		})

		//属性转换
		for _, property := range properties {

			if property.Type != nil {
				property.Type.GenerateExpression(property.Format, typeConvert)
			}

			property.Alias = cmp.Or(e.env.Alias.Properties[property.Name], property.Name)
		}

		//是否确定导出
		kind := ref.Type.Kind
		if kind&tmpl.ImmutableType != 0 && kind&tmpl.GenericType != 0 {
			ref.Ignore = true
		}
		if len(ref.Properties) == 0 {
			ref.Ignore = true
		}

		//别名
		ref.Alias = cmp.Or(e.env.Alias.Modes[ref.Name], ref.Name)
	}

	//查询
	findType := func(currenType *tmpl.NamedType) *tmpl.NamedType {

		currenName := currenType.Expression
		match, ok := lo.Find(outRefs, func(item *tmpl.Ref) bool {
			return item.Name == currenName
		})

		if ok && match.Type.Kind&tmpl.GenericType != 0 {
			return match.Type
		}

		//默认类型
		currenType.GenerateExpression("", e.convert)
		return currenType
	}

	//匹配路径参数
	joinPath := func(path *tmpl.Path) string {
		parameters := path.Parameters

		//存在路径参数
		parameters = lo.Filter(parameters, func(item *tmpl.Parameter, idx int) bool {
			return item.In == "path"
		})

		if len(parameters) > 0 {
			return lang.Format(LANG, path.OriginalPath, e.env.Alias.Parameters)
		}
		return fmt.Sprintf(`"%s"`, path.Path)
	}

	//根据类型 生成表达式
	for _, api := range outApis {

		paths := api.Paths
		slices.SortFunc(paths, func(a, b *tmpl.Path) int {
			return cmp.Compare(a.Name, b.Name)
		})

		for _, path := range paths {

			//参数转换
			for _, parameter := range path.Parameters {
				parameter.Type.GenerateExpression(parameter.Format, typeConvert)
				//别名
				parameter.Alias = cmp.Or(e.env.Alias.Parameters[parameter.Name], parameter.Name)
			}

			//路径
			path.Path = joinPath(path)

			//全量参数
			slices.SortFunc(path.Parameters, func(a, b *tmpl.Parameter) int {
				return cmp.Compare(a.Name, b.Name)
			})

			if path.Request != nil {
				path.Parameters = append(path.Parameters, &tmpl.Parameter{
					Name:  "body",
					Alias: "body",
					In:    "body",
					Type:  findType(path.Request),
				})
			}

			//query参数
			queries := lo.Filter(path.Parameters, func(item *tmpl.Parameter, index int) bool {
				return item.In == "query"
			})
			slices.SortFunc(queries, func(a, b *tmpl.Parameter) int {
				return cmp.Compare(a.Name, b.Name)
			})
			path.Queries = queries

			if path.Response != nil {
				path.Response = findType(path.Response)
			}

			//方法名 不同Tag下 方法重名下容易生成 add_1,update_39
			path.Name = lo.TernaryF(e.env.RepeatableOperationId, func() string {
				return strings.FieldsFunc(path.Name, func(r rune) bool {
					return r == '_' || r == '-'
				})[0]
			}, func() string {
				return path.Name
			})

		}

		api.Paths = paths
	}

	//写入接口文件
	w := newWriter(e.env, outRefs, outApis)

	//写入接口
	err = w.api(e.env.Output, "header", e.engine)
	if err != nil {
		return err
	}

	//写入接口适配
	err = w.client(e.env.ClientOutput, "client", e.engine)
	return err
}
