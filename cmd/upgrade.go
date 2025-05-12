package cmd

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade to the latest version or specified version",
	RunE: func(cmd *cobra.Command, args []string) error {
		if repoUrl != "" {
			return downloadInstall(repoUrl)
		}

		//查询 github
		assets, err := downloadCheck(repoVersion)
		if err != nil {
			return err
		}
		return downloadInstall(assets.Url)
	},
}

func init() {
	upgradeCmd.Flags().StringVarP(&repoUrl, "download url", "d", repoUrl, "repo download url")
	upgradeCmd.Flags().StringVarP(&repoVersion, "repo version", "v", repoVersion, "repo version")
}

const repo = "https://api.github.com/repos/otk-final/openapi-codegen/releases"

var repoUrl = ""
var repoVersion = "latest"
var repoName = "openapi"

type Assets struct {
	Name string `json:"name"`
	Url  string `json:"browser_download_url"`
}
type RepoInfo struct {
	Tag         string    `json:"tag"`
	Name        string    `json:"name"`
	CreatedAt   string    `json:"created_at"`
	PublishedAt string    `json:"published_at"`
	Assets      []*Assets `json:"assets"`
}

func downloadCheck(version string) (*Assets, error) {

	//查询
	resp, err := http.Get(repo + "/" + version)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)

	//解码
	var info RepoInfo
	_ = json.Unmarshal(body, &info)

	//获取当前os 安装包
	assets, ok := lo.Find(info.Assets, func(item *Assets) bool {
		return strings.Contains(item.Name, runtime.GOOS)
	})
	if ok {
		return assets, nil
	}
	return nil, errors.New("not found assets")
}

func downloadInstall(url string) error {
	fmt.Printf("start download %s\n", url)

	//下载
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	//临时目录 文件
	tempDir, _ := os.UserCacheDir()
	tempDir = filepath.Join(tempDir, repoName)
	_ = os.MkdirAll(tempDir, os.ModePerm)

	tempFile, err := os.OpenFile(filepath.Join(tempDir, fmt.Sprintf("%d.zip", time.Now().UnixMilli())), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	//写入文件
	bar := progressbar.DefaultBytes(resp.ContentLength, "downloading")
	_, err = io.Copy(io.MultiWriter(tempFile, bar), resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("download completed\n")

	//解压文件
	tempZip, _ := zip.OpenReader(tempFile.Name())
	defer func() {
		_ = tempZip.Close()
		_ = os.Remove(tempFile.Name())
	}()

	//获取指定文件
	newFile, ok := lo.Find(tempZip.File, func(item *zip.File) bool {
		return strings.Contains(item.Name, repoName)
	})
	if !ok {
		return errors.New("not found install file from zip")
	}

	newReader, _ := newFile.Open()
	newBytes, _ := io.ReadAll(newReader)

	//获取真实安装路径
	exePath, _ := os.Executable()
	realPath, _ := filepath.EvalSymlinks(exePath)

	//替换文件 覆盖
	return os.WriteFile(realPath, newBytes, os.ModePerm)
}
