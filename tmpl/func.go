package tmpl

func Variable(mapping map[string]string, target string) string {
	return mapping[target]
}
