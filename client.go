package trakt

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/jingweno/go-sawyer"
	"github.com/jingweno/go-sawyer/hypermedia"
)

func NewClient(apiKey string, authMethod AuthMethod) *Client {
	return NewClientWith(TraktAPIURL, UserAgent, apiKey, authMethod, nil)
}

func NewClientWith(baseURL string, UserAgent string, apiKey string, authMethod AuthMethod, httpClient *http.Client) *Client {
	client, _ := sawyer.NewFromString(baseURL, httpClient)
	return &Client{Client: client, UserAgent: UserAgent, ApiKey: apiKey, AuthMethod: authMethod}
}

type Client struct {
	*sawyer.Client

	UserAgent  string
	AuthMethod AuthMethod
	ApiKey     string
	rootRels   hypermedia.Relations
}

func (c *Client) NewRequest(urlStr string) (req *Request, err error) {
	// TODO Move this to a more sane place
	if strings.Contains(urlStr, "?") == true {
		urlStr += "&extended=full,images"
	} else {
		urlStr += "?extended=full,images"
	}
	req, err = newRequest(c, urlStr)
	if err != nil {
		return
	}

	c.applyRequestHeaders(req)

	return
}

// a GET request with specific media type set
func (c *Client) getBody(url *url.URL, mediaType string) (patch io.ReadCloser, result *Result) {
	result = sendRequest(c, url, func(req *Request) (*Response, error) {
		req.Header.Set("Accept", mediaType)
		return req.Get(nil)
	})

	if result.Response != nil {
		patch = result.Response.Body
	}

	return
}

func (c *Client) get(url *url.URL, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Get(output)
	})
}

func (c *Client) post(url *url.URL, input interface{}, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Post(input, output)
	})
}

func (c *Client) put(url *url.URL, input interface{}, output interface{}) *Result {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Put(input, output)
	})
}

func (c *Client) delete(url *url.URL, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Delete(output)
	})
}

func (c *Client) patch(url *url.URL, input interface{}, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Patch(input, output)
	})
}

func (c *Client) upload(uploadUrl *url.URL, asset io.ReadCloser, contentType string, contentLength int64) (result *Result) {
	req, err := c.NewRequest(uploadUrl.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	req.Header.Set("Content-Type", contentType)
	req.ContentLength = contentLength

	req.Body = asset
	sawyerResp := req.Request.Post()

	resp, err := NewResponse(sawyerResp)
	return newResult(resp, err)
}

func (c *Client) applyRequestHeaders(req *Request) {
	req.Header.Set("Accept", DefaultMediaType)
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("trakt-api-version", TraktAPIVersion)
	req.Header.Set("trakt-api-key", c.ApiKey)

	if tokenAuth, ok := c.AuthMethod.(TokenAuth); ok {
		req.Header.Set("Authorization", tokenAuth.String())
	}

	// Go doesn't apply `Host` on the header, instead it consults `Request.Host`
	// Populate `Host` if it exists in `Client.Header`
	// See Bug https://code.google.com/p/go/issues/detail?id=7682
	host := c.Header.Get("Host")
	if host != "" {
		req.Request.Host = host
	}

	return
}

func sendRequest(c *Client, url *url.URL, fn func(r *Request) (*Response, error)) (result *Result) {
	req, err := c.NewRequest(url.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	resp, err := fn(req)
	result = newResult(resp, err)

	return
}
