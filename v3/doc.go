package v3

type Doc struct {
	Openapi string `json:"openapi,omitempty"`
	Info    struct {
		Title   string `json:"title,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"info,omitempty"`
	Servers []struct {
		Url         string `json:"url,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"servers,omitempty"`
	Tags       []Tag           `json:"tags,omitempty"`
	Paths      map[string]Path `json:"paths,omitempty"`
	Components struct {
		Schemas map[string]Mode `json:"schemas,omitempty"`
	} `json:"components,omitempty"`
}
