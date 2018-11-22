package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"time"
)

// SampleResponse sample
type SampleResponse struct {
	Result baseResponse `json:"result"`
}

type baseResponse struct {
	Status int        `json:"status"`
	Data   Servertime `json:"data"`
}

// Servertime sample
type Servertime struct {
	Timestamp string `json:"timestamp"`
}

// Client sample
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Username   string
	Password   string
}

// NewClient create new Client
func NewClient(urlStr, username, password string) (*Client, error) {
	//if len(username) == 0 {
	//	return nil, errors.New("missing  username")
	//}

	//if len(password) == 0 {
	//	return nil, errors.New("missing  password")
	//}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		//return nil, errors.New("failed to parse url: %s", urlStr)
		return nil, errors.New("failed to parse url")
	}

	c := new(Client)
	c.URL = parsedURL
	c.HTTPClient = &http.Client{
		Timeout: 30 * time.Second,
	}

	//c.Username = username
	//c.Password = password
	return c, nil
}

func (c *Client) SendRequest(ctx context.Context, spath string) (string, error) {

	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return "", err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	// Check status code here…

	bodyString, err := decodeBodyToString(res)
	if err != nil {
		return "", err
	}

	return bodyString, nil
}

func (c *Client) GetServerTime(ctx context.Context) (*SampleResponse, error) {
	spath := "/timestamp.php"

	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check status code here…

	//var servertime interface{}
	var servertime SampleResponse

	if err := decodeBody(res, &servertime); err != nil {
		return nil, err
	}

	return &servertime, nil
}

func (c *Client) GetUser(ctx context.Context, userID, pID string) (string, error) {
	spath := "/oratta_dev.php/user_page/native_login_view"

	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", err
	}

	cookie := &http.Cookie{
		Name:  "PHPSESSID",
		Value: "kpp5cf9j6qe89viiranqs500h1",
	}

	jar.SetCookies(c.URL, []*http.Cookie{cookie})
	c.HTTPClient.Jar = jar

	body := url.Values{}
	body.Add("pid", userID)
	body.Add("platform", pID)
	req, err := c.newRequest(
		ctx,
		"GET",
		spath,
		strings.NewReader(body.Encode()),
	)
	if err != nil {
		return "", err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	// Check status code here…

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
	//req.SetBasicAuth(c.Username, c.Password)

	// set user agent
	//var userAgent = fmt.Sprintf("XXXGoClient/%s (%s)", version, runtime.Version())
	//req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	//b, err := ioutil.ReadAll(resp.Body)
	//if err == nil {
	//	fmt.Println(string(b))
	//}

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
