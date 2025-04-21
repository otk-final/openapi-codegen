package v3

import "strings"

type Tag struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Path = map[string]Method

type Method struct {
	Tags        []string            `json:"tags,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	OperationId string              `json:"operationId,omitempty"`
	Responses   map[string]Response `json:"responses,omitempty"`
	RequestBody RequestBody         `json:"requestBody,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty"`
}

type Response struct {
	Description string  `json:"description,omitempty"`
	Content     Content `json:"content,omitempty"`
	Ref         string  `json:"$ref,omitempty"`
}

type RequestBody struct {
	Description string  `json:"description,omitempty"`
	Content     Content `json:"content,omitempty"`
	Ref         string  `json:"$ref,omitempty"`
}

type Content = map[string]struct {
	Schema Schema `json:"schema,omitempty"`
}

type Parameter struct {
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Schema      Schema `json:"schema,omitempty"`
	Description string `json:"description,omitempty"`
}

type Schema struct {
	Type   string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Ref    string `json:"$ref,omitempty"`
	Items  struct {
		Type string `json:"type,omitempty"`
		Ref  string `json:"$ref,omitempty"`
	} `json:"items,omitempty"`
	Default interface{} `json:"default,omitempty"`
}

type Mode struct {
	Type          string              `json:"type,omitempty"`
	Properties    map[string]Property `json:"properties,omitempty"`
	Description   string              `json:"description,omitempty"`
	Discriminator struct {
		PropertyName string `json:"propertyName,omitempty"`
	} `json:"discriminator,omitempty"`
}

type Property struct {
	Schema
	Description string `json:"description,omitempty"`
	Enum        []any  `json:"enum,omitempty"`
}

type RefPath string

func (r RefPath) BaseName() string {
	x := r[strings.LastIndex(string(r), "/")+1:]
	return string(x)
}
