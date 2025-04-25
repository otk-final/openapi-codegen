package cmd

import (
	"codegen/internal"
	"codegen/lang"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var startArgs = &internal.Args{}

func init() {
	//start
	startCmd.Flags().StringVarP(&startArgs.Version, "version", "v", defaultVersion, "openapi version")
	startCmd.Flags().StringVarP(&startArgs.Endpoint, "endpoint", "e", "", "example：https://{server}:{port}/v3/api-docs")
	startCmd.Flags().StringVarP(&startArgs.Output, "output", "o", "", "api output file")
	startCmd.Flags().StringVarP(&startArgs.ClientOutput, "client_output", "c", "", "client output file")
	startCmd.Flags().StringVarP(&startArgs.Lang, "lang", "l", "", strings.Join(lang.Names(), ","))
	startCmd.Flags().StringVarP(&startArgs.Style, "style", "s", "", "customize template file")

}

// run from flag
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Quick start",
	Run: func(cmd *cobra.Command, args []string) {
		//默认配置
		env := &internal.Env{
			Args:   startArgs,
			Ignore: make([]string, 0),
			Filter: make([]string, 0),
			Alias: internal.Alias{
				Properties: make(map[string]string),
				Modes:      make(map[string]string),
				Types:      make(map[string]string),
				Parameters: make(map[string]string),
			},
			Variables: map[string]string{},
			Generics: &internal.Generics{
				Enable:      false,
				Expressions: nil,
			},
		}

		exe, err := internal.New(env, args)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = exe.Run(env.Args)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("SUCCESS")
	},
}
