// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httpclient

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var client *HttpClient
var once sync.Once

type HttpClient struct {
	client      *http.Client
	transport   *http.Transport
	method      string
	timeout     time.Duration
	url         string
	body        io.Reader // Use POST/PUT/DELETE
	queryParams Params
	debug       bool
	mux         sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc
}

func NewHttpClientOr() *HttpClient {
	once.Do(func() {
		client = &HttpClient{
			client: &http.Client{},
		}
	})
	return client
}

func (cli *HttpClient) init() *HttpClient {
	cli.method = ""
	cli.url = ""
	cli.body = nil
	cli.queryParams = nil
	cli.debug = false
	cli.transport = nil
	return cli
}

func (cli *HttpClient) newRequest(method string, url string) *HttpClient {
	cli.mux.Lock()
	defer cli.mux.Unlock()

	// Reset parameters.
	cli.init()

	// New parameters.
	cli.method = method
	cli.url = url
	return cli
}

func (cli *HttpClient) Body(body io.Reader) *HttpClient {
	cli.mux.Lock()
	defer cli.mux.Unlock()

	cli.body = body
	return cli
}

func (cli *HttpClient) BodyWithJSON(s string) *HttpClient {
	cli.mux.Lock()
	defer cli.mux.Unlock()

	cli.body = strings.NewReader(s)
	return cli
}

func (cli *HttpClient) QueryParams(params Params) *HttpClient {
	cli.mux.Lock()
	defer cli.mux.Unlock()

	cli.queryParams = params
	return cli
}

func (cli *HttpClient) Timeout(wait time.Duration) *HttpClient {
	cli.mux.Lock()
	defer cli.mux.Unlock()

	cli.timeout = wait
	return cli
}

func (cli *HttpClient) InsecureSkipVerify(skip bool) *HttpClient {
	cli.mux.Lock()
	defer cli.mux.Unlock()

	if skip {
		if cli.transport == nil {
			cli.transport = &http.Transport{}
		}

		if cli.transport.TLSClientConfig == nil {
			cli.transport.TLSClientConfig = &tls.Config{}
		}
		cli.transport.TLSClientConfig.InsecureSkipVerify = true
	}
	return cli
}

func (cli *HttpClient) Debug(debug bool) *HttpClient {
	cli.mux.Lock()
	defer cli.mux.Unlock()

	cli.debug = debug
	return cli
}

func (cli *HttpClient) Do() (*http.Response, error) {
	cli.mux.Lock()
	defer cli.mux.Unlock()

	// Use transport.
	cli.useTransport()

	// New request.
	req, err := http.NewRequest(cli.method, cli.url, cli.body)
	if err != nil {
		return nil, err
	}

	// Use query parameters,
	// if request Method is GET and call QueryParams method,
	// then use
	cli.useQueryParams(req)

	// Show debug information.
	cli.showDebug(req)

	// Use timeout.
	req = cli.useTimeout(req)
	if cli.cancel != nil {
		defer cli.cancel()
	}

	// Execute http request.
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cli *HttpClient) useQueryParams(req *http.Request) {
	if cli.queryParams != nil {
		query := req.URL.Query()

		for _, param := range cli.queryParams {
			query.Add(param.Key, param.Value)
		}

		req.URL.RawQuery = query.Encode()
	}
}

func (cli *HttpClient) useTransport() {
	if cli.transport != nil {
		cli.client.Transport = cli.transport
	}
}

func (cli *HttpClient) showDebug(req *http.Request) {
	if cli.debug {
		log.Printf("fill url: %s\n", req.URL.String())
	}
}

func (cli *HttpClient) useTimeout(req *http.Request) *http.Request {
	if cli.timeout != 0 {
		cli.ctx, cli.cancel = context.WithTimeout(context.Background(), cli.timeout)
		return req.WithContext(cli.ctx)
	}
	return req
}

func Get(url string) *HttpClient {
	return NewHttpClientOr().newRequest(http.MethodGet, url)
}

func Post(url string) *HttpClient {
	return NewHttpClientOr().newRequest(http.MethodPost, url)
}
