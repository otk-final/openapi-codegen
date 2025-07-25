package v2

import (
	"cmp"
	"codegen/tmpl"
	"github.com/samber/lo"
	"strings"
)

type Tag struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Path = map[string]Method

type Method struct {
	Tags        []string            `json:"tags,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Description string              `json:"description,omitempty"`
	OperationId string              `json:"operationId,omitempty"`
	Responses   map[string]Response `json:"responses,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty"`
}

type Response struct {
	Description string  `json:"description,omitempty"`
	Schema      *Schema `json:"schema,omitempty"`
}

type Schema struct {
	Type   string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Ref    string `json:"$ref,omitempty"`
	Items  *struct {
		Type string `json:"type,omitempty"`
		Ref  string `json:"$ref,omitempty"`
	} `json:"items,omitempty"`
	AdditionalProperties *struct {
		Type string `json:"type,omitempty"`
		Ref  string `json:"$ref,omitempty"`
	} `json:"additionalProperties,omitempty"`
	Default interface{} `json:"default,omitempty"`
}

type Parameter struct {
	Schema
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
}

type Mode struct {
	Type        string              `json:"type,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty"`
	Description string              `json:"description,omitempty"`
}

type Property struct {
	Schema
	Description string `json:"description,omitempty"`
	Enum        []any  `json:"enum,omitempty"`
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

	return &tmpl.NamedType{Kind: kind, Expression: ModePath(expression).BaseName()}
}

type ModePath string

func (r ModePath) BaseName() string {
	x := r[strings.LastIndex(string(r), "/")+1:]
	return string(x)
}
