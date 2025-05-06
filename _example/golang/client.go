package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ApiClient TODO You need to implement the current interface
type ApiClient[T any] interface {
	Get(name string, params map[string]any) (T, error)

	Post(name string, params any) (T, error)

	Put(name string, params any) (T, error)

	Delete(name string, params any) (T, error)
}

type DefaultClient[T any] struct {
	Host    string
	Headers http.Header
}

func NewClient[T any](host string, headers http.Header) ApiClient[T] {
	return &DefaultClient[T]{host, headers}
}

func (t *DefaultClient[T]) Get(name string, params map[string]any) (T, error) {

	uri, _ := url.Parse(t.Host)
	uri.Path = name

	if params != nil {
		query := url.Values{}
		for k, v := range params {
			query.Set(k, fmt.Sprintf("%v", v))
		}
		uri.RawQuery = query.Encode()
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: t.Headers,
	}

	return t.request(req)
}

func (t *DefaultClient[T]) Post(name string, params any) (T, error) {
	return t.doBody(http.MethodPost, name, params)
}

func (t *DefaultClient[T]) Put(name string, params any) (T, error) {
	return t.doBody(http.MethodPut, name, params)
}

func (t *DefaultClient[T]) Delete(name string, params any) (T, error) {
	return t.doBody(http.MethodDelete, name, params)
}

func (t *DefaultClient[T]) doBody(method string, name string, params any) (T, error) {
	uri, _ := url.Parse(t.Host)
	uri.JoinPath(name)

	req := &http.Request{
		Method: method,
		URL:    uri,
		Header: t.Headers,
	}

	if params != nil {
		body, _ := json.Marshal(params)
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	return t.request(req)
}
func (t *DefaultClient[T]) request(req *http.Request) (T, error) {

	var ret T

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ret, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
