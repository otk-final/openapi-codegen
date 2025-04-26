package internal

import (
	"testing"
)

var testArgs = &Args{
	Version: "v3",
	//Endpoint: "http://175.178.57.240/xhm-api/v3/api-docs/api",
	Endpoint: "http://localhost:8083/v3/api-docs",
	//Endpoint: "http://localhost:8082/v2/api-docs",

	//Output:       "/Users/hxy/develops/openapi/client-ts/src/demo.ts",
	//ClientOutput: "/Users/hxy/develops/openapi/client-ts/src/client.ts",

	Output:       "/Users/hxy/develops/openapi/client-swift-x/client-api/Api.swift",
	ClientOutput: "/Users/hxy/develops/openapi/client-swift-x/client-api/Client.swift",

	//Output:       "/Users/hxy/develops/openapi/client-kotlin/app/src/main/java/com/demo/Api.kt",
	//ClientOutput: "/Users/hxy/develops/openapi/client-kotlin/app/src/main/java/com/demo/Client.kt",

	//Output:       "/Users/hxy/develops/openapi/openapi-codegen/demo/api.go",
	//ClientOutput: "/Users/hxy/develops/openapi/openapi-codegen/demo/client.go",

	//Output:  "/Users/hxy/develops/demo/internal/demo.py",
	//Output:  "/Users/hxy/develops/openapi/openapi-server/v3/src/main/java/com/demo",
	Lang: "swift",
}

func TestCmd(t *testing.T) {

	env := &Env{
		Args: testArgs,
		Ignore: []string{
			"/error",
			"/pt",
		},
		Filter: []string{
			//"/pb/product",
		},
		Alias: Alias{
			Properties: map[string]string{},
			Modes:      map[string]string{},
			Types: map[string]string{
				"JsonNode": "[String:String]",
			},
			Parameters: map[string]string{
				//"type": "type2",
			},
		},
		Generics: &Generics{
			Enable: true,
			Unfold: false,
			Expressions: map[string][]string{
				"ApiResult": {"data"},
				"PageData":  {"entities+"},
				"MapResult": {"data~"},
			},
		},
		Variables: map[string]string{
			"clientPackage": "com.demo",
			"structPackage": "com.demo",
			"apiPackage":    "com.demo",
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
