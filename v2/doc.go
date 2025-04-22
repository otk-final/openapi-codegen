package v2

type Doc struct {
	Swagger string `json:"openapi,omitempty"`
	Info    struct {
		Title   string `json:"title,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"info,omitempty"`
	Host        string
	BasePath    string
	Tags        []Tag           `json:"tags,omitempty"`
	Paths       map[string]Path `json:"paths,omitempty"`
	Definitions map[string]Mode `json:"definitions,omitempty"`
}
