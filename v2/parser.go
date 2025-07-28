package v2

import (
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

			if item.In == "body" {
				continue
			}

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

		req, ok := lo.Find(info.Parameters, func(item Parameter) bool {
			return item.In == "body"
		})

		if !ok {
			return nil
		}

		return schemaType(req.Schema).Parse()
	}

	toResponse := func(path string, method string, info Method) *tmpl.NamedType {
		response := info.Responses["200"]
		if response.Schema == nil {
			fmt.Printf("Waning: undefine response structur [%s %s]\n", method, path)
			return tmpl.VoidNamedType
		}
		return schemaType(*response.Schema).Parse()
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
