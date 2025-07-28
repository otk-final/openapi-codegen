package v3

import (
	"cmp"
	"codegen/tmpl"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"io"
	"net/http"
	"os"
	"strings"
)

func LoadParse(addr string) ([]*tmpl.Ref, []*tmpl.Api, error) {

	var body []byte
	if strings.HasPrefix(addr, "http") {
		resp, err := http.Get(addr)
		if err != nil {
			return nil, nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, nil, fmt.Errorf("%s %s", http.StatusText(resp.StatusCode), addr)
		}

		defer func() {
			_ = resp.Body.Close()
		}()
		body, _ = io.ReadAll(resp.Body)
	} else {
		//文件读取
		body, _ = os.ReadFile(addr)
	}

	var doc = &Doc{}
	err := json.Unmarshal(body, doc)
	if err != nil {
		return nil, nil, err
	}

	refs := make([]*tmpl.Ref, 0)

	toPropertyType := func(info Property) *tmpl.NamedType {
		return schemaType(info.Schema).Parse()
	}

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

	schemas := doc.Components.Schemas
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

			tp := schemaType(item.Schema).Parse()

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

	toRequest := func(path string, method string, info Method) (nt *tmpl.NamedType) {

		schema := info.RequestBody.Content["application/json"].Schema
		return schemaType(schema).Parse()
	}

	toResponse := func(path string, method string, info Method) *tmpl.NamedType {
		response := info.Responses["200"]

		//未声明响应
		content := response.Content
		if len(content) == 0 {
			fmt.Printf("Waning: undefine response structur [%s %s]\n", method, path)
			return tmpl.VoidNamedType
		}

		return schemaType(content["*/*"].Schema).Parse()
	}

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
				Request:      toRequest(path, method, fn),
				Response:     toResponse(path, method, fn),
			}
			paths = append(paths, p)
		}
	}

	groups := lo.GroupBy(paths, func(item *tmpl.Path) string {
		return item.Tag
	})

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

type schemaType Schema

func (s schemaType) Parse() *tmpl.NamedType {

	var expression string
	var kind tmpl.NamedTypeKind

	if s.Ref != "" {
		kind = tmpl.ReferenceType
		expression = s.Ref
	} else {
		if s.Items != nil {

			expression = cmp.Or(s.Items.Type, s.Items.Ref)
			kind = lo.Ternary(expression == s.Items.Type, tmpl.ArrayType|tmpl.FoundationType, tmpl.ArrayType|tmpl.ReferenceType)
		} else if s.AdditionalProperties != nil {

			expression = cmp.Or(s.AdditionalProperties.Type, s.AdditionalProperties.Ref)
			kind = lo.Ternary(expression == s.AdditionalProperties.Type, tmpl.MapType|tmpl.FoundationType, tmpl.MapType|tmpl.ReferenceType)
		} else {
			expression = cmp.Or(s.Type, s.Ref)
			kind = lo.Ternary(expression == s.Type, tmpl.FoundationType, tmpl.ReferenceType)
		}
	}

	if expression == "" {
		return nil
	}

	return &tmpl.NamedType{Kind: kind, Expression: RefPath(expression).BaseName()}
}
