package cmd

import (
	"codegen/internal"
	"codegen/tmpl"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var initArgs = &internal.Args{
	Name:     "swagger",
	Endpoint: "http://localhost:8080/v3/api-docs",
	Lang:     "ts",
	Version:  defaultVersion,
}

func init() {
	//init
	initCmd.Flags().StringVarP(&initArgs.Version, "version", "v", initArgs.Version, "openapi version")
	initCmd.Flags().StringVarP(&initArgs.Endpoint, "endpoint", "e", initArgs.Endpoint, "example：https://{server}:{port}/v3/api-docs")
	initCmd.Flags().StringVarP(&initArgs.Lang, "lang", "l", initArgs.Lang, strings.Join(tmpl.LangNames(), ","))
}

// create env file
var initCmd = &cobra.Command{

	Use:   "init",
	Short: "Initialize environment variable configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {

		//创建openapi.json
		defaultEnv := &internal.Env{
			Args:   initArgs,
			Output: tmpl.NewOutputs(initArgs.Lang),
			Ignore: []string{},
			Filter: []string{},
			Alias: internal.Alias{
				Properties: make(map[string]string),
				Modes:      make(map[string]string),
				Types:      make(map[string]string),
				Parameters: make(map[string]string),
			},

			Variables: map[string]string{},
			Generics: &internal.Generics{
				Enable: false,
				Unfold: false,
				Expressions: map[string][]string{
					"ApiResult": {"data"},
				},
			},
			RepeatableOperationId: true,
		}
		//当前执行目录
		envFile := tmpl.PwdJoinPath(defaultEnvFileName)

		envs := make([]*internal.Env, 0)
		envs = append(envs, defaultEnv)

		data, _ := json.MarshalIndent(envs, "", "   ")
		return os.WriteFile(envFile, data, os.ModePerm)
	},
}
