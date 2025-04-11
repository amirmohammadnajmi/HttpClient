package HttpClient

import (
	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	client *resty.Client
}

var client *HttpClient

func init() {
	client = &HttpClient{
		client: resty.New(),
	}
}

func Get(url string, headers map[string]string) (*resty.Response, error) {
	return client.get(url, headers)
}

func Post(url string, headers map[string]string, body interface{}) (*resty.Response, error) {
	return client.post(url, headers, body)
}

func Put(url string, headers map[string]string, body interface{}) (*resty.Response, error) {
	return client.put(url, headers, body)
}

func (hc *HttpClient) get(url string, headers map[string]string) (*resty.Response, error) {
	return hc.client.R().SetHeaders(headers).Get(url)
}

func (hc *HttpClient) post(url string, headers map[string]string, body interface{}) (*resty.Response, error) {
	return hc.client.R().SetHeaders(headers).SetBody(body).Post(url)
}

func (hc *HttpClient) put(url string, headers map[string]string, body interface{}) (*resty.Response, error) {
	return hc.client.R().SetHeaders(headers).SetBody(body).Put(url)
}
