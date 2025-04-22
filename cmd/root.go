package cmd

import (
	"codegen/internal"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "A Gen Code tools",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// run from flag
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A Gen Code tools",
	RunE: func(cmd *cobra.Command, args []string) error {
		//默认配置
		env := &internal.Env{
			Args:          startArgs,
			Ignore:        make([]string, 0),
			Filter:        make([]string, 0),
			TypeAlias:     make(map[string]string),
			PropertyAlias: make(map[string]string),
			Variables:     map[string]string{},
			Generics: &internal.Generics{
				Enable:      false,
				Expressions: nil,
			},
		}

		exe, err := internal.New(env, args)
		if err != nil {
			return err
		}
		return exe.Run(env.Args)
	},
}

// create env file
var initCmd = &cobra.Command{

	Use:   "init",
	Short: "init env file",
	RunE: func(cmd *cobra.Command, args []string) error {

		//创建openapi.json
		defaultEnv := &internal.Env{
			Args:          initArgs,
			Ignore:        []string{},
			Filter:        []string{},
			TypeAlias:     map[string]string{},
			PropertyAlias: map[string]string{},
			Variables: map[string]string{
				"apiPackage":    "",
				"clientPackage": "",
				"structPackage": "",
			},
			Generics: &internal.Generics{
				Enable:      false,
				Expressions: map[string][]string{},
			},
		}
		//当前执行目录
		pwd, _ := os.Executable()
		envFile := path.Join(filepath.Dir(pwd), "openapi.json")

		data, _ := json.MarshalIndent(&defaultEnv, "", "   ")
		return os.WriteFile(envFile, data, os.ModePerm)
	},
}

// reload from env
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reload openapi ",
	RunE: func(cmd *cobra.Command, args []string) error {

		//当前执行目录
		if envFile == "" {
			pwd, _ := os.Executable()
			envFile = path.Join(filepath.Dir(pwd), "openapi.json")
		}

		data, err := os.ReadFile(envFile)
		if err != nil {
			return err
		}

		//当前配置文件
		env := &internal.Env{}
		err = json.Unmarshal(data, env)
		if err != nil {
			return err
		}

		exe, err := internal.New(env, args)
		if err != nil {
			return err
		}

		return exe.Run(env.Args)
	},
}

var startArgs = &internal.Args{}
var initArgs = &internal.Args{}
var envFile string

func init() {

	// 根据openapi 生成接口文档

	//start
	startCmd.Flags().StringVarP(&startArgs.Version, "version", "v", "v3", "openapi version")
	startCmd.Flags().StringVarP(&startArgs.Endpoint, "endpoint", "e", "", "server api https://")
	startCmd.Flags().StringVarP(&startArgs.Output, "output", "o", "", "output directory")
	startCmd.Flags().StringVarP(&startArgs.Lang, "lang", "l", "", "target code language go, java, ts")
	startCmd.Flags().StringVarP(&startArgs.Style, "style", "s", "", "custom template file")

	//init
	initCmd.Flags().StringVarP(&initArgs.Version, "version", "v", "v3", "openapi version")
	initCmd.Flags().StringVarP(&initArgs.Endpoint, "endpoint", "e", "", "server api https://")
	initCmd.Flags().StringVarP(&initArgs.Output, "output", "o", "", "output directory")
	initCmd.Flags().StringVarP(&initArgs.Lang, "lang", "l", "", "target code language go, java, ts")
	initCmd.Flags().StringVarP(&initArgs.Style, "style", "s", "", "custom template file")

	//reload
	reloadCmd.Flags().StringVarP(&envFile, "env", "e", "", "custom env file")

	rootCmd.AddCommand(startCmd, initCmd, reloadCmd)
}

func Run() error {
	return rootCmd.Execute()
}
