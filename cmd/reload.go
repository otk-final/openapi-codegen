package cmd

import (
	"codegen/internal"
	"codegen/lang"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var envFile string
var envLang string
var defaultEnvFileName = "openapi.json"

func init() {

	//reload
	reloadCmd.Flags().StringVarP(&envFile, "file", "f", "", "customize env file")
	reloadCmd.Flags().StringVarP(&envLang, "lang", "l", "", strings.Join(lang.Names(), ","))

}

// reload from env
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload with environment variable configuration file",
	Run: func(cmd *cobra.Command, args []string) {

		//当前执行目录
		if envFile == "" {
			pwd, _ := os.Executable()
			envFile = path.Join(filepath.Dir(pwd), defaultEnvFileName)
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

		var env *internal.Env
		if envLang != "" {
			env, _ = lo.Find(envs, func(item *internal.Env) bool {
				return strings.EqualFold(envLang, item.Lang)
			})
		} else {
			if len(envs) > 0 {
				env = envs[0]
			}
		}

		if env == nil {
			fmt.Println("Not found current env")
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
