package internal

import (
	"bytes"
	"codegen/tmpl"
	"fmt"
	"os"
	"path/filepath"
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

	file := tmpl.PwdJoinPath(output.File)
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
