package v2

import (
	"codegen/tmpl"
	"encoding/json"
	"github.com/samber/lo"
	"io"
	"net/http"
)

func LoadParse(addr string) ([]*tmpl.Ref, []*tmpl.Api, error) {

	//fetch
	resp, err := http.Get(addr)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	//read
	body, _ := io.ReadAll(resp.Body)
	var doc = &Doc{}
	err = json.Unmarshal(body, doc)
	if err != nil {
		return nil, nil, err
	}

	//load
	refs := make([]*tmpl.Ref, 0)

	toPropertyType := func(info Property) *tmpl.NamedType {
		return schemaType(info.Schema).Parse()
	}

	//properties
	toProperties := func(info Mode) tmpl.Properties {

		array := make(tmpl.Properties, 0)
		for name, property := range info.Properties {

			array = append(array, &tmpl.Property{
				Name:        name,
				Description: property.Description,
				Type:        toPropertyType(property),
				Format:      property.Format,
				Enums:       []string{},
			})
		}
		return array
	}

	//refs
	schemas := doc.Definitions
	for name, mode := range schemas {

		ref := &tmpl.Ref{
			Name: name,
			Type: &tmpl.NamedType{
				Kind:       tmpl.ReferenceType,
				Expression: name,
			},
			Properties:  toProperties(mode),
			Description: mode.Description,
		}

		refs = append(refs, ref)
	}

	toParameters := func(method string, info Method) tmpl.Parameters {
		parameters := make([]*tmpl.Parameter, 0)
		for _, item := range info.Parameters {

			//过滤body
			if item.In == "body" {
				continue
			}

			tp := schemaType(item.Schema).Parse()

			// 只支持基础数据类型 FoundationType ArrayType
			parameters = append(parameters, &tmpl.Parameter{
				Name:        item.Name,
				Required:    item.Required,
				Type:        tp,
				In:          item.In,
				Format:      item.Schema.Format,
				Description: item.Description,
			})
		}
		return parameters
	}

	toRequest := func(method string, info Method) (nt *tmpl.NamedType) {

		req, ok := lo.Find(info.Parameters, func(item Parameter) bool {
			return item.In == "body"
		})

		if !ok {
			return nil
		}
		//包装类型
		return schemaType(req.Schema).Parse()
	}

	toResponse := func(method string, info Method) *tmpl.NamedType {
		response := info.Responses["200"]
		return schemaType(response.Schema).Parse()
	}

	//paths
	paths := make([]*tmpl.Path, 0)
	for path, item := range doc.Paths {
		for method, fn := range item {

			p := &tmpl.Path{
				Tag:          lo.Ternary(len(fn.Tags) > 0, fn.Tags[0], "UnnamedApi"),
				Name:         fn.OperationId,
				Description:  fn.Description,
				Summary:      fn.Summary,
				Path:         path,
				OriginalPath: path,
				Method:       method,
				Parameters:   toParameters(method, fn),
				Request:      toRequest(method, fn),
				Response:     toResponse(method, fn),
			}
			paths = append(paths, p)
		}
	}

	//分组
	groups := lo.GroupBy(paths, func(item *tmpl.Path) string {
		return item.Tag
	})

	//apis
	apis := make([]*tmpl.Api, 0)
	for _, tag := range doc.Tags {
		apis = append(apis, &tmpl.Api{
			Name:        tag.Name,
			Description: tag.Description,
			Paths:       groups[tag.Name],
		})
	}

	return refs, apis, nil
}
