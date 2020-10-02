package OrderBookSnapStore

import (
	"context"
	"io"
	"net/http"
	"time"
	"github.com/bitly/go-simplejson"
)

func MakeHeaders(ctx context.Context) map[string]string {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Cache-Control"] = "no-cache"
	return header
}


func makeReqHeaders(req *http.Request, headers map[string]string) *http.Request {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req
}

func getClient(
	timeout time.Duration) (client *http.Client) {

	tr := &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		DisableCompression:  true,
	}

	client = &http.Client{Transport: tr, Timeout: timeout}
	return
}

func Request(
	method string,
	url string,
	headers map[string]string,
	timeout time.Duration,
	body io.Reader) (http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return http.Response{}, err
	}

	req = makeReqHeaders(req, headers)

	client := getClient(timeout)
	resp, err := client.Do(req)

	if err != nil {
		return http.Response{}, err
	}

	if resp != nil {
		return *resp, nil
	}

	return http.Response{}, nil
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}
