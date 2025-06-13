package internal

import (
	"bytes"
	"codegen/tmpl"
	"fmt"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type writer interface {
	write(output *tmpl.Output, apis []*tmpl.Api, models []*tmpl.Ref, engine *template.Template) error
}

func newWriter(name string, env *Env) writer {

	//标准输入
	std := &stdWriter{name, env}

	//java 需要输出多个文件
	if env.Lang == "java" {
		return &javaWriter{name, env, std}
	}
	return std
}

type stdWriter struct {
	name string
	env  *Env
}

func (w *stdWriter) write(output *tmpl.Output, apis []*tmpl.Api, models []*tmpl.Ref, engine *template.Template) error {
	buf := &bytes.Buffer{}

	models = lo.Filter(models, func(item *tmpl.Ref, index int) bool {
		return !item.Ignore
	})

	data := struct {
		Output *tmpl.Output
		Env    *Env
		Apis   []*tmpl.Api
		Models []*tmpl.Ref
	}{
		Output: output,
		Apis:   apis,
		Models: models,
		Env:    w.env,
	}

	//写入 header
	err := engine.Execute(buf, data)
	if err != nil {
		return err
	}

	//格式化路径参数
	textFile := output.File
	if len(apis) == 1 {
		textFile = strings.ReplaceAll(textFile, "{api}", apis[0].Name)
	}
	if len(models) == 1 {
		textFile = strings.ReplaceAll(textFile, "{model}", models[0].Name)
	}

	file := tmpl.PwdJoinPath(textFile)
	dir := filepath.Dir(file)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Output %s \n", file)

	return os.WriteFile(file, buf.Bytes(), os.ModePerm)
}

type javaWriter struct {
	name string
	env  *Env
	std  writer
}

// java 一个文件只能有一个model 或者 api 需写入多个文件
func (w *javaWriter) write(output *tmpl.Output, apis []*tmpl.Api, models []*tmpl.Ref, engine *template.Template) error {
	if w.name == "api" {
		for _, api := range apis {
			err := w.std.write(output, []*tmpl.Api{api}, models, engine)
			if err != nil {
				return err
			}
		}
		return nil
	} else if w.name == "model" {
		for _, model := range models {
			err := w.std.write(output, apis, []*tmpl.Ref{model}, engine)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return w.std.write(output, apis, models, engine)
}
