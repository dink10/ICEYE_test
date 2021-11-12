package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/motemen/go-loghttp"
	"github.com/sirupsen/logrus"
)

// Client is used to implement http client
type Client struct {
	*http.Client
}

// NewClient creates http Client
func NewClient(cfg *Config) *Client {
	tr := loghttp.Transport{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: cfg.MaxConn,
			TLSHandshakeTimeout: time.Duration(cfg.HandshakeTimeout),
		},
		LogRequest: func(req *http.Request) {
			var bodyBytes []byte
			if cfg.LogRequests {
				if req.Body != nil {
					var err error
					bodyBytes, err = ioutil.ReadAll(req.Body)
					if err != nil {
						logrus.Error(err)
					}
				}
				req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}

			if cfg.LogRequests {
				logrus.Debugf("API request [%p] %s %s %s", req, req.Method, req.URL, string(bodyBytes))
			}
		},
		LogResponse: func(resp *http.Response) {
			var bodyBytes []byte
			if cfg.LogRequestBody {
				if resp.Body != nil {
					var err error
					bodyBytes, err = ioutil.ReadAll(resp.Body)
					if err != nil {
						logrus.Error(err)
					}
				}
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}
			if cfg.LogRequests {
				logrus.Debugf("API response [%p] %d %s %s", resp.Request, resp.StatusCode, resp.Request.URL, string(bodyBytes))
			}
		},
	}

	return &Client{&http.Client{Transport: &tr, Timeout: time.Duration(cfg.ClientTimeOut) * time.Second}}
}

type Options struct {
	Headers map[string]string
	Params  map[string]string
	Body    []byte
}

// DoRequest - do request
// nolint: gocyclo
func (c Client) DoRequest(url, method string, opts Options) ([]byte, error) {

	if url == "" {
		return nil, fmt.Errorf("%s", "requst URL is empty")
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(opts.Body))
	if err != nil {
		return nil, err
	}
	if req.URL == nil {
		return nil, fmt.Errorf("empty request URL object")
	}
	q := req.URL.Query()
	for k, v := range opts.Params {
		q.Add(k, v)
	}
	if q != nil {
		req.URL.RawQuery = q.Encode()
	}

	for k, v := range opts.Headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("response body is empty")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		errB := resp.Body.Close()
		if errB != nil {
			logrus.Error(errB)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d; url %s; response error body: %s", resp.StatusCode, url, body)
	}

	return body, nil
}
