package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "",
	Short:   "openapi codegen tool",
	Long:    "This is a tool that generates API call code in various programming languages based on the content of an OpenAPI document.",
	Example: "openapi start -e http://localhost:8080/v3/api-doc -l ts -o src/api.ts",
}

var defaultVersion = "v3"

func init() {
	// 根据openapi 生成接口文档
	rootCmd.AddCommand(startCmd, initCmd, reloadCmd)
}

func Run() error {
	return rootCmd.Execute()
}
