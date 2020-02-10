package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client http client
type Client struct {
	URL       string
	Method    string
	ReqHeader map[string]string
	ReqBody   []byte
	ResHeader map[string]string
	ResBody   []byte
}

// SetUrl set http url
func (c *Client) SetURL(url string) *Client {
	c.URL = url
	return c
}

// SetMethod set http method
func (c *Client) SetMethod(method string) *Client {
	c.Method = method
	return c
}

// SetReqHeader set http request header
func (c *Client) SetReqHeader(header map[string]string) *Client {
	if header == nil {
		return c
	}
	head := make(map[string]string)
	for k, v := range header {
		head[k] = v
	}
	c.ReqHeader = head
	return c
}

// SetReqBody set http request body
func (c *Client) SetReqBody(body []byte) *Client {
	if len(body) == 0 {
		return c
	}
	c.ReqBody = body
	return c
}

// Do run http request
func (c *Client) Do() error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	var (
		request *http.Request
		err     error
	)
	if c.Method == "" {
		c.Method = "GET"
	}
	switch c.Method {
	case "GET", "DELETE":
		request, err = http.NewRequest(c.Method, c.URL, nil)
	case "POST", "PUT":
		request, err = http.NewRequest(c.Method, c.URL, bytes.NewBuffer(c.ResBody))
	default:
		return errors.New(fmt.Sprintf("Unknown request method: %s", c.Method))
	}
	if err != nil {
		return err
	}
	// request header
	if c.ReqHeader != nil && len(c.ReqHeader) > 0 {
		for k, v := range c.ReqHeader {
			request.Header.Add(k, v)
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// response header
	c.ResHeader = make(map[string]string)
	for k, v := range response.Header {
		tmp := ""
		for _, val := range v {
			tmp = fmt.Sprintf("%s%s", tmp, val)
		}
		c.ResHeader[k] = tmp
	}
	// head["StatusCode"] = fmt.Sprintf("%d",response.StatusCode)
	// response body
	buffer := [128]byte{}
	result := bytes.NewBuffer(nil)
	for {
		n, err := response.Body.Read(buffer[0:])
		if err != nil {
			if err == io.EOF {
				result.Write(buffer[0:n])
				break
			}
			return err
		}
		result.Write(buffer[0:n])
	}
	c.ResBody = result.Bytes()
	return nil
}
