package cmd

import (
	"codegen/internal"
	"codegen/lang"
	"codegen/tmpl"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var initArgs = &internal.Args{
	Name:         "swagger",
	Endpoint:     "http://localhost:8080/v3/api-docs",
	Output:       "src/api.ts",
	ClientOutput: "src/client.ts",
	Lang:         "",
	Style:        "",
	Version:      "",
}

func init() {
	//init
	initCmd.Flags().StringVarP(&initArgs.Version, "version", "v", defaultVersion, "openapi version")
	initCmd.Flags().StringVarP(&initArgs.Endpoint, "endpoint", "e", "", "example：https://{server}:{port}/v3/api-docs")
	initCmd.Flags().StringVarP(&startArgs.Output, "output", "o", "", "api output file")
	initCmd.Flags().StringVarP(&startArgs.ClientOutput, "client_output", "c", "", "client output file")
	initCmd.Flags().StringVarP(&initArgs.Lang, "lang", "l", "", strings.Join(lang.Names(), ","))
	initCmd.Flags().StringVarP(&initArgs.Style, "style", "s", "", "customize template file")

}

// create env file
var initCmd = &cobra.Command{

	Use:   "init",
	Short: "Initialize environment variable configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {

		if initArgs.Name == "" {
			initArgs.Name = initArgs.Lang
		}

		//创建openapi.json
		defaultEnv := &internal.Env{
			Args:   initArgs,
			Ignore: []string{},
			Filter: []string{},
			Alias: internal.Alias{
				Properties: make(map[string]string),
				Modes:      make(map[string]string),
				Types:      make(map[string]string),
				Parameters: make(map[string]string),
			},
			Variables: map[string]string{
				"apiPackage":    "",
				"clientPackage": "",
				"structPackage": "",
			},
			Generics: &internal.Generics{
				Enable: false,
				Unfold: false,
				Expressions: map[string][]string{
					"ApiResult": {"data"},
				},
			},
		}
		//当前执行目录
		envFile := tmpl.PwdJoinPath(defaultEnvFileName)

		envs := make([]*internal.Env, 0)
		envs = append(envs, defaultEnv)

		data, _ := json.MarshalIndent(envs, "", "   ")
		return os.WriteFile(envFile, data, os.ModePerm)
	},
}
