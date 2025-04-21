package internal

import (
	"bytes"
	"codegen/tmpl"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"text/template"
)

type writer interface {
	api(output string, name string, engine *template.Template) error
	client(output string, name string, engine *template.Template) error
}

func newWriter(env *Env, refs []*tmpl.Ref, apis []*tmpl.Api) writer {

	//标准输入
	std := &stdWriter{env, refs, apis}

	switch env.Lang {
	case "java":
		return &javaWriter{env, refs, apis}
	case "python":
		return &pythonWriter{std}
	default:
		return std
	}
}

type stdWriter struct {
	env  *Env
	refs []*tmpl.Ref
	apis []*tmpl.Api
}

func (w *stdWriter) api(output string, name string, engine *template.Template) error {
	buf := &bytes.Buffer{}

	//写入 header
	err := engine.ExecuteTemplate(buf, name, w.env)
	if err != nil {
		return err
	}

	// 写入 struct
	for _, ref := range w.refs {

		if ref.Ignore {
			continue
		}

		err := engine.ExecuteTemplate(buf, "struct", ref)
		if err != nil {
			return err
		}
	}

	// 写入 api
	for _, api := range w.apis {
		err := engine.ExecuteTemplate(buf, "api", api)
		if err != nil {
			return err
		}
	}

	//不存在则创建目录
	dir := path.Dir(output)
	_ = os.MkdirAll(dir, os.ModePerm)

	return os.WriteFile(output, buf.Bytes(), os.ModePerm)
}

func (w *stdWriter) client(output string, name string, engine *template.Template) error {
	if output == "" {
		return nil
	}

	if n := engine.Lookup(name); n == nil {
		return nil
	}

	buf := &bytes.Buffer{}
	err := engine.ExecuteTemplate(buf, name, w.env)
	if err != nil {
		return err
	}

	return os.WriteFile(output, buf.Bytes(), os.ModePerm)
}

type javaWriter struct {
	env  *Env
	refs []*tmpl.Ref
	apis []*tmpl.Api
}

func (w *javaWriter) api(output string, name string, engine *template.Template) error {

	//创建目录
	refDir := path.Join(path.Dir(output), "dto")
	_ = os.MkdirAll(refDir, os.ModePerm)

	apiDir := path.Join(path.Dir(output), "api")
	_ = os.MkdirAll(apiDir, os.ModePerm)

	// java 一个文件只能有一个struct 需写入多个文件
	for _, ref := range w.refs {

		if ref.Ignore {
			continue
		}

		fileWriter := &stdWriter{
			env:  w.env,
			refs: []*tmpl.Ref{ref},
			apis: make([]*tmpl.Api, 0),
		}

		structFile := path.Join(output, "dto", fmt.Sprintf("%s.java", ref.Name))
		err := fileWriter.api(structFile, "structHeader", engine)
		if err != nil {
			return err
		}
	}

	for _, api := range w.apis {

		fileWriter := &stdWriter{
			env:  w.env,
			refs: []*tmpl.Ref{},
			apis: []*tmpl.Api{api},
		}

		apiFile := path.Join(output, "api", fmt.Sprintf("%s.java", api.Name))
		err := fileWriter.api(apiFile, "apiHeader", engine)
		if err != nil {
			log.Print("err", err.Error())
			return err
		}
	}

	return nil
}

func (w *javaWriter) client(output string, name string, engine *template.Template) error {

	fileWriter := &stdWriter{
		env:  w.env,
		refs: []*tmpl.Ref{},
		apis: []*tmpl.Api{},
	}

	//接口
	clientFile := path.Join(output, "ApiClient.java")
	_ = fileWriter.client(clientFile, "client", engine)

	return nil
}

type pythonWriter struct {
	*stdWriter
}

func (w *pythonWriter) api(output string, name string, engine *template.Template) error {
	// python 有顺序要求，重新排序
	slices.SortFunc(w.refs, func(a, b *tmpl.Ref) int {
		return a.ReferenceLevel() - b.ReferenceLevel()
	})

	return w.stdWriter.api(output, name, engine)
}

func (w *pythonWriter) client(output string, name string, engine *template.Template) error {
	return w.stdWriter.client(output, name, engine)
}
