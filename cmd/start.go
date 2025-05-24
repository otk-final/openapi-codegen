package cmd

import (
	"codegen/internal"
	"codegen/tmpl"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {

	//init
	startCmd.Flags().StringVarP(&initArgs.Version, "version", "v", initArgs.Version, "openapi version")
	startCmd.Flags().StringVarP(&initArgs.Endpoint, "endpoint", "e", initArgs.Endpoint, "example：https://{server}:{port}/v3/api-docs")
	startCmd.Flags().StringVarP(&initArgs.Lang, "lang", "l", initArgs.Lang, strings.Join(tmpl.LangNames(), ","))
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Quick start",
	Run: func(cmd *cobra.Command, args []string) {

		env := &internal.Env{
			Args:   initArgs,
			Ignore: make([]string, 0),
			Filter: make([]string, 0),
			Output: tmpl.NewOutputs(initArgs.Lang),
			Alias: internal.Alias{
				Properties: make(map[string]string),
				Models:     make(map[string]string),
				Types:      make(map[string]string),
				Parameters: make(map[string]string),
			},
			Variables: map[string]string{},
			Generics: &internal.Generics{
				Enable:      false,
				Unfold:      false,
				Expressions: map[string][]string{},
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
