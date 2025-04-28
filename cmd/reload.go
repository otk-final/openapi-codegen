package cmd

import (
	"codegen/internal"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

var envFile string
var envName string
var defaultEnvFileName = "openapi.json"

func init() {

	//reload
	reloadCmd.Flags().StringVarP(&envFile, "file", "f", "", "customize env file")
	reloadCmd.Flags().StringVarP(&envName, "name", "n", "", "env name")
}

// reload from env
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload with environment variable configuration file",
	Run: func(cmd *cobra.Command, args []string) {

		//当前执行目录
		if envFile == "" {
			pwd, _ := os.Getwd()
			envFile = path.Join(pwd, defaultEnvFileName)
		}

		data, err := os.ReadFile(envFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		envs := make([]*internal.Env, 0)

		//当前配置文件
		err = json.Unmarshal(data, &envs)
		if err != nil {
			fmt.Println(err)
			return
		}

		if envName != "" {
			envs = lo.Filter(envs, func(item *internal.Env, idx int) bool {
				return strings.EqualFold(envName, item.Name)
			})
		}

		if len(envs) == 0 {
			fmt.Println("Not found current env")
			return
		}

		for _, env := range envs {

			fmt.Printf("[%s] Begin Generate\n", env.Name)

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

			fmt.Printf("[%s] Generate SUCCESS\n", env.Name)
		}

		fmt.Println("\nFinished")
	},
}
