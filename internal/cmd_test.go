package internal

import (
	"regexp"
	"testing"
)

var testArgs = &Args{
	Version:  "v3",
	Endpoint: "http://175.178.57.240/xhm-api/v3/api-docs/api",
	//Endpoint: "http://localhost:8083/v3/api-docs",
	//Endpoint: "http://localhost:8082/v2/api-docs",

	Output:       "/Users/hxy/develops/xhm/XHM-Admin/src/gc/demo.tsx",
	ClientOutput: "/Users/hxy/develops/xhm/XHM-Admin/src/gc/client.tsx",

	//Output:  "/Users/hxy/develops/demo/Codegen/Codegen/Demo.swift",
	//Output:  "/Users/hxy/develops/demo/sts/app/src/main/java/com/otk/sts/internal/Api.kt",

	//Output:       "/Users/hxy/develops/openapi/openapi-codegen/demo/api.go",
	//ClientOutput: "/Users/hxy/develops/openapi/openapi-codegen/demo/client.go",

	//Output:  "/Users/hxy/develops/demo/internal/demo.py",
	//Output:  "/Users/hxy/develops/openapi/openapi-server/v3/src/main/java/com/demo",
	Lang: "ts",
}

func TestCmd(t *testing.T) {

	env := &Env{
		Args: testArgs,
		Ignore: []string{
			"/error",
		},
		Filter: []string{},
		TypeAlias: map[string]string{
			//"JsonNode": "[String:String]",
			//"object":   "[String:String]",
			"JsonNode": "any",
			//"object": "JsonNode",
		},
		PropertyAlias: map[string]string{
			//"type": "type2",
		},
		Generics: &Generics{
			Enable: true,
			Unfold: false,
			Expressions: map[string][]string{
				"ApiResult": {"data"},
				"PageData":  {"entities+"},
			},
		},
		Variables: map[string]string{
			"clientPackage": "demo",
			"structPackage": "com.demo.dto",
			"apiPackage":    "demo",
		},
	}
	exe, err := New(env, []string{})
	if err != nil {
		t.Error(err)
	}
	err = exe.Run(env.Args)

	if err != nil {
		t.Error(err)
	}
}

func TestReg(t *testing.T) {
	text := "ApiResult<csla,abc>"

	re := regexp.MustCompile(`^(\w+)<([\w,\s]+)>$`)
	matches := re.FindStringSubmatch(text)
	t.Log(matches)
	//t.Log(strings.Split(matches[1], ","))
}
