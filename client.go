package captain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://cds.captain.ai/"
)

// A Client manages communication with the CAPTAIN API.
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HTTPClient *http.Client

	integrationKey string
	developerKey   string
}

func (c *Client) SetIntegrationKey(key string) {
	c.integrationKey = key
}

func (c *Client) SetDeveloperKey(key string) {
	c.developerKey = key
}

// NewClient returns a new CAPTAIN API client.
func NewClient() *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// Response is a CAPTAIN API response. This wraps the standard http.Response
// returned from CAPTAIN and provides convenient access to things like
// custom headers.
type Response struct {
	*http.Response
	Runtime   time.Duration
	RequestID string
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *Response // HTTP response that caused this error
	Errors   json.RawMessage
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("[%v] %v %v: %d %v %s", r.Response.RequestID,
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Response.Runtime, r.Errors)
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, apiPath string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(apiPath)
	if err != nil {
		return nil, err
	}
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.integrationKey != "" {
		req.Header.Set("X-Integration-Key", c.integrationKey)
	}
	if c.developerKey != "" {
		req.Header.Set("X-Developer-Key", c.developerKey)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	resp, err := c.HTTPClient.Do(req.WithContext(ctx))
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		return err
	}
	defer resp.Body.Close()
	clientResp := newResponse(resp)
	err = CheckResponse(clientResp)
	if err != nil {
		return err
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, clientResp.Body)
		} else {
			decErr := json.NewDecoder(clientResp.Body).Decode(v)
			if decErr != nil {
				return decErr
			}
		}

	}
	return nil
}

func newResponse(r *http.Response) *Response {
	resp := &Response{Response: r}
	resp.RequestID = r.Header.Get("X-Request-Id")
	runtime := r.Header.Get("X-Runtime")
	secs, _ := strconv.ParseFloat(runtime, 64)
	resp.Runtime = time.Duration(secs * float64(time.Second))
	return resp
}

func CheckResponse(r *Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	if r.StatusCode == 422 {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		errorResponse.Errors = json.RawMessage(body)
	}
	return errorResponse
}
