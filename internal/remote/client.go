// Copyright 2019 yhyzgn enp-simple
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-05-14 15:28
// version: 1.0.0
// desc   : 

package remote

import (
	"context"
	"encoding/json"
	"fmt"
	"gox-simple/internal/built"
	"gox-simple/internal/res"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ContentTypeJSON = "application/json;charset=UTF-8"
	ContentTypeForm = "application/x-www-form-urlencoded"
)

type Method string

type Client struct {
	ctx     context.Context
	method  Method
	baseURL string
	path    string
	header  http.Header
	query   url.Values
	form    url.Values
	data    map[string]interface{}
}

func NewClient(method Method) *Client {
	ctx, _ := context.WithTimeout(context.TODO(), 6*time.Second)
	return &Client{
		ctx:    ctx,
		method: method,
		header: http.Header{},
		query:  url.Values{},
		form:   url.Values{},
		data:   make(map[string]interface{}),
	}
}

func GET() *Client {
	return NewClient("GET")
}

func POST() *Client {
	return NewClient("POST")
}

func PUT() *Client {
	return NewClient("PUT")
}

func DELETE() *Client {
	return NewClient("DELETE")
}

func (c *Client) BaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

func (c *Client) Path(path string) *Client {
	c.path = path
	return c
}

func (c *Client) Header(key, value string) *Client {
	c.header.Set(key, value)
	return c
}

func (c *Client) ContentType(contentType string) *Client {
	return c.Header("Content-Type", contentType)
}

func (c *Client) Cookie(name string, value interface{}) *Client {
	return c.Header("Cookie", fmt.Sprintf("%v=%v", name, value))
}

func (c *Client) Query(name string, value interface{}) *Client {
	c.query.Set(name, fmt.Sprintf("%v", value))
	return c
}

func (c *Client) Form(name string, value interface{}) *Client {
	c.form.Set(name, fmt.Sprintf("%v", value))
	return c
}

func (c *Client) Data(name string, value interface{}) *Client {
	c.data[name] = value
	return c
}

func (c *Client) Response(bean interface{}) (result *res.Response, err error) {
	return c.Execute(newJsonResolver(bean))
}

func (c *Client) Execute(resolver ResultResolver) (result *res.Response, err error) {
	contentType := c.header.Get("Content-Type")
	if contentType == "" {
		contentType = ContentTypeForm
		c.ContentType(contentType)
	}

	c.Header("User-Agent", built.FullNameWithBuildDate)

	uri, err := c.uri()
	if err != nil {
		return
	}

	var req *http.Request
	if c.method == "POST" || c.method == "PUT" || c.method == "DELETE" {
		var reader *strings.Reader
		if contentType == ContentTypeForm {
			reader = strings.NewReader(c.form.Encode())
		} else if contentType == ContentTypeJSON {
			bs, e := json.Marshal(c.data)
			if e != nil {
				err = e
				return
			}
			reader = strings.NewReader(string(bs))
		} else {
			// 默认
			reader = strings.NewReader(c.form.Encode())
		}
		temp, e := http.NewRequestWithContext(c.ctx, string(c.method), uri.String(), reader)
		if e != nil {
			err = e
			return
		}
		req = temp
	} else if c.method == "GET" {
		uri.RawQuery = c.query.Encode()
		temp, e := http.NewRequestWithContext(c.ctx, string(c.method), uri.String(), nil)
		if e != nil {
			err = e
			return
		}
		req = temp
	} else {
		err = fmt.Errorf("unknown http method [%v]", c.method)
		return
	}

	if req == nil {
		err = fmt.Errorf("req is nil pointer")
		return
	}

	req.Header = c.header

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return
	}

	if r.StatusCode == http.StatusOK {
		bs, e := ioutil.ReadAll(r.Body)
		if e != nil {
			err = e
			return
		}
		if resolver != nil {
			return resolver.Resolve(bs)
		}
	} else {
		err = fmt.Errorf("http response error with status code [%d]", r.StatusCode)
	}

	return
}

func (c *Client) uri() (*url.URL, error) {
	return url.Parse(c.baseURL + c.path)
}
