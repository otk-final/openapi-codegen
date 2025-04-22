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
	"regexp"
	"slices"
	"strings"
	"text/template"
)

type Args struct {
	Endpoint     string `json:"endpoint"`
	Output       string `json:"output"`
	ClientOutput string `json:"client_output"`
	Lang         string `json:"lang"`
	Style        string `json:"style"`
	Version      string `json:"version"`
}

type Env struct {
	*Args
	Ignore        []string          `json:"ignore"`
	Filter        []string          `json:"filter"`
	TypeAlias     map[string]string `json:"typeAlias"`
	PropertyAlias map[string]string `json:"propertyAlias"`
	Variables     map[string]string `json:"variables"`
	Generics      *Generics         `json:"generics"`
}

type Generics struct {
	Enable      bool                `json:"enable"`
	Expressions map[string][]string `json:"expressions"`
}

type Executor struct {
	env     *Env
	convert lang.TypeConvert
	tp      *template.Template
}

func New(env *Env, args []string) (*Executor, error) {

	//初始化模版
	tp, err := tmpl.NewEngine(env.Lang)
	if err != nil {
		return nil, err
	}

	return &Executor{
		env: env,
		tp:  tp,
	}, nil
}

func (e *Executor) Run(cmd *Args) error {

	//target language
	LANG := cmd.Lang
	typeConvert := lang.NewConvert(LANG, e.env.TypeAlias)
	e.convert = typeConvert

	var outRefs []*tmpl.Ref
	var outApis []*tmpl.Api
	var err error

	//加载数据
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
	// path 过滤
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

	//排序
	slices.SortFunc(outRefs, func(a, b *tmpl.Ref) int {
		return cmp.Compare(a.Name, b.Name)
	})

	slices.SortFunc(outApis, func(a, b *tmpl.Api) int {
		return cmp.Compare(a.Name, b.Name)
	})

	//是否需要生成泛型
	if e.env.Generics.Enable {

		// 匹配通用泛型
		expressions := e.env.Generics.Expressions

		//查找泛型
		generics := resolvingGenerics(e.convert, expressions, outRefs)

		mgr := &nestingGenericManage{
			generics: generics,
			refs:     outRefs,
		}

		newRefs := make([]*tmpl.Ref, 0)
		newRefs = append(newRefs, generics...)

		//修改 ref 所有匹配类型 为 泛型
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

				//设置泛型>不可变类型
				ref.Type.Kind = tmpl.GenericType | tmpl.ImmutableType
				ref.Type.Expression = val.discriminator(typeConvert)
			}

			//TODO 属性中是否存在泛型
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

			property.Type.GenerateExpression(property.Format, typeConvert)

			property.Alias = cmp.Or(e.env.PropertyAlias[property.Name], property.Name)
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
		ref.Alias = cmp.Or(e.env.PropertyAlias[ref.Name], ref.Name)
	}

	//查询
	findType := func(currenType *tmpl.NamedType) *tmpl.NamedType {
		currenName := currenType.Expression
		match, ok := lo.Find(outRefs, func(item *tmpl.Ref) bool {
			return item.Name == currenName
		})
		if ok {
			return match.Type
		}

		//默认类型
		currenType.GenerateExpression("", e.convert)
		return currenType
	}

	//匹配路径参数

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
				parameter.Alias = cmp.Or(e.env.PropertyAlias[parameter.Name], parameter.Name)
			}

			//排序
			slices.SortFunc(path.Parameters, func(a, b *tmpl.Parameter) int {
				return cmp.Compare(a.Name, b.Name)
			})

			//路径
			path.Path = joinPathVariable(LANG, path)

			//重新匹配
			if path.Request != nil {
				path.Request = findType(path.Request)
			}

			if path.Response != nil {
				path.Response = findType(path.Response)
			}
		}

		api.Paths = paths
	}

	//写入接口文件
	w := newWriter(e.env, outRefs, outApis)

	//写入接口
	err = w.api(e.env.Output, "header", e.tp)

	//写入接口适配
	err = w.client(e.env.ClientOutput, "client", e.tp)

	return err
}

// 匹配路径参数
func joinPathVariable(language string, path *tmpl.Path) string {

	parameters := path.Parameters

	//存在路径参数
	paths := lo.Filter(parameters, func(item *tmpl.Parameter, idx int) bool {
		return item.In == "path"
	})

	rawPath := path.Path
	if len(paths) > 0 {

		reg := regexp.MustCompile(`\{([^/]+?)\}`)
		rawPath = reg.ReplaceAllStringFunc(rawPath, func(seg string) string {
			key := reg.FindStringSubmatch(seg)[1]
			switch language {
			case "swift":
				return fmt.Sprintf(`\(%s)`, key)
			case "go", "golang":
				return `" + fmt.Sprintf("%v",` + key + `) +"`
			case "java":
				return fmt.Sprintf(`" + %s +"`, key)
			default:
				return fmt.Sprintf(`" + params.%s +"`, key)
			}
		})
	}

	rawPath = `"` + rawPath + `"`
	//处理结尾 +""
	if strings.HasSuffix(rawPath, `+""`) {
		return rawPath[:strings.LastIndex(rawPath, `+""`)]
	}
	return rawPath
}
