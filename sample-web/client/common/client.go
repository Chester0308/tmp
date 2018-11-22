package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"time"
)

// Client sample
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Username   string
	Password   string
	Timeout    int
}

// ClientOption functional option
type ClientOption func(*Client) error

// BasicAuth functional option
func BasicAuth(name, password string) ClientOption {
	return func(c *Client) error {
		c.Username = name
		c.Password = password
		return nil
	}
}

// Timeout functional option
func Timeout(t time.Duration) ClientOption {
	return func(c *Client) error {
		c.HTTPClient.Timeout = t
		return nil
	}
}

// NewClient create new Client
func NewClient(urlStr string, options ...ClientOption) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.New("failed to parse url")
	}

	// initialize default value
	c := &Client{
		URL: parsedURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	// functional option
	for _, option := range options {
		option(c)
	}

	fmt.Printf("\n\nBasic Auth\n  name = %s \n  password = %s\n\n", c.Username, c.Password)
	return c, nil
}

// GetFullName get request sample
func (c *Client) GetFullName(ctx context.Context, fname, lname string) (string, error) {
	spath := "/fullname"
	method := "GET"

	req, err := c.newRequest(ctx, method, spath, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("firstname", fname)
	q.Add("lastname", lname)
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	// Check status code here

	bodyString, err := decodeBodyToString(res)
	if err != nil {
		return "", err
	}

	return bodyString, nil
}

// PostMessage post request sample
func (c *Client) PostMessage(ctx context.Context, name, message string) (string, error) {
	spath := "/message"
	method := "POST"
	body := url.Values{}
	body.Add("name", name)
	body.Add("message", message)

	req, err := c.newRequest(ctx, method, spath, strings.NewReader(body.Encode()))
	if err != nil {
		return "", err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	// Check status code here

	bodyString, err := decodeBodyToString(res)
	if err != nil {
		return "", err
	}

	return bodyString, nil
}

// SendRequest send request
func (c *Client) SendRequest(ctx context.Context, method, spath string) (string, error) {
	req, err := c.newRequest(ctx, method, spath, nil)
	if err != nil {
		return "", err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	// Check status code here

	// Cookie ???
	//jar, err := cookiejar.New(nil)
	//if err != nil {
	//	return "", err
	//}

	//cookie := &http.Cookie{
	//	Name:  "hoge",
	//	Value: "huga",
	//}

	//jar.SetCookies(c.URL, []*http.Cookie{cookie})
	//c.HTTPClient.Jar = jar

	// Cookie ???
	cookies := res.Cookies()
	for _, v := range cookies {
		fmt.Printf("%v\n", v)
	}

	bodyString, err := decodeBodyToString(res)
	if err != nil {
		return "", err
	}

	return bodyString, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// set basic auth
	req.SetBasicAuth(c.Username, c.Password)

	// set user agent
	//var userAgent = fmt.Sprintf("XXXGoClient/%s (%s)", version, runtime.Version())
	//req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func decodeBodyToString(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(b))
		return "", err
	}

	return string(b), nil
}

func assert(data interface{}) {
	switch data.(type) {
	case string:
		fmt.Print(data.(string))
	case float64:
		fmt.Print(data.(float64))
	case bool:
		fmt.Print(data.(bool))
	case nil:
		fmt.Print("null")
	case []interface{}:
		fmt.Print("[")
		for _, v := range data.([]interface{}) {
			assert(v)
			fmt.Print(" ")
		}
		fmt.Print("]")
	case map[string]interface{}:
		fmt.Print("{")
		for k, v := range data.(map[string]interface{}) {
			fmt.Print(k + ":")
			assert(v)
			fmt.Print(" ")
		}
		fmt.Print("}")
	default:
	}
}
