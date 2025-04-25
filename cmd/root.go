package cmd

import (
	"codegen/internal"
	"codegen/lang"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:     "",
	Short:   "openapi codegen tool",
	Long:    "This is a tool that generates API call code in various programming languages based on the content of an OpenAPI document.",
	Example: "openapi start -e http://localhost:8080/v3/api-doc -l ts -o src/api.ts",
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

// create env file
var initCmd = &cobra.Command{

	Use:   "init",
	Short: "Initialize environment variable configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {

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
	Short: "Reload with environment variable configuration file",
	Run: func(cmd *cobra.Command, args []string) {

		//当前执行目录
		if envFile == "" {
			pwd, _ := os.Executable()
			envFile = path.Join(filepath.Dir(pwd), "openapi.json")
		}

		data, err := os.ReadFile(envFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		//当前配置文件
		env := &internal.Env{}
		err = json.Unmarshal(data, env)
		if err != nil {
			fmt.Println(err)
			return
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

var startArgs = &internal.Args{}
var initArgs = &internal.Args{}
var envFile string

var defaultVersion = "v3"

func init() {

	// 根据openapi 生成接口文档

	//start
	startCmd.Flags().StringVarP(&startArgs.Version, "version", "v", defaultVersion, "openapi version")
	startCmd.Flags().StringVarP(&startArgs.Endpoint, "endpoint", "e", "", "example：https://{server}:{port}/v3/api-docs")
	startCmd.Flags().StringVarP(&startArgs.Output, "output", "o", "", "api output file")
	startCmd.Flags().StringVarP(&startArgs.ClientOutput, "client_output", "c", "", "client output file")
	startCmd.Flags().StringVarP(&startArgs.Lang, "lang", "l", "", strings.Join(lang.Names(), ","))
	startCmd.Flags().StringVarP(&startArgs.Style, "style", "s", "", "customize template file")

	//init
	initCmd.Flags().StringVarP(&initArgs.Version, "version", "v", defaultVersion, "openapi version")
	initCmd.Flags().StringVarP(&initArgs.Endpoint, "endpoint", "e", "", "example：https://{server}:{port}/v3/api-docs")
	initCmd.Flags().StringVarP(&startArgs.Output, "output", "o", "", "api output file")
	initCmd.Flags().StringVarP(&startArgs.ClientOutput, "client_output", "c", "", "client output file")
	initCmd.Flags().StringVarP(&initArgs.Lang, "lang", "l", "", strings.Join(lang.Names(), ","))
	initCmd.Flags().StringVarP(&initArgs.Style, "style", "s", "", "customize template file")

	//reload
	reloadCmd.Flags().StringVarP(&envFile, "env", "e", "", "customize env file")

	rootCmd.AddCommand(startCmd, initCmd, reloadCmd)
}

func Run() error {
	return rootCmd.Execute()
}
