package cmd

import (
	"codegen/internal"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "openapi",
	Short: "A Gen Code tools ",
	RunE: func(cmd *cobra.Command, args []string) error {
		runner, err := internal.New(cmdEnv, cmdArgs, args)
		if err != nil {
			return err
		}
		return runner.Run(cmdArgs, args)
	},
}

var cmdArgs = &internal.Args{}
var cmdEnv *internal.Env

func init() {

	// 根据openapi 生成接口文档
	rootCmd.Flags().StringVarP(&cmdArgs.Version, "version", "v", "v3", "openapi 版本")
	rootCmd.Flags().StringVarP(&cmdArgs.Endpoint, "endpoint", "e", "", "接口地址 https://xx")
	rootCmd.Flags().StringVarP(&cmdArgs.Output, "output", "o", "./internal", "输出目录")
	rootCmd.Flags().StringVarP(&cmdArgs.Lang, "lang", "l", "ts", "代码语言（如 go, java, ts）")
	rootCmd.Flags().StringVarP(&cmdArgs.Style, "style", "s", "", "自定义模版")

}

func Run() error {
	//当前执行目录
	pwd, _ := os.Executable()
	filepath.Dir(pwd)

	return rootCmd.Execute()
}
