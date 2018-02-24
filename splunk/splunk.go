package splunk

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/url"

	"gopkg.in/resty.v1"
)

const (
	PathSavedSearchCreate = "/services/saved/searches"
	PathSavedSearch       = "/services/saved/searches/%s"
)

// Client communicates with the Splunk rest endpoint.
type Client struct {
	client *resty.Client
}

// Feed is used to store Splunk response Atom feeds
type Feed struct {
	Links    map[string]string `schema:"-" json:"links"`
	Origin   string            `schema:"-" json:"origin"`
	Updated  string            `schema:"-" json:"updated"`
	Entry    []Entry           `schema:"-" json:"entry"`
	Messages []Message         `schema:"-" json:"messages"`
}

// Entry is used to store Splunk response Atom entries
type Entry struct {
	Name    string                 `schema:"-" json:"name"`
	ID      string                 `schema:"-" json:"id"`
	Updated string                 `schema:"-" json:"updated"`
	Links   map[string]string      `schema:"-" json:"links"`
	Author  string                 `schema:"-" json:"author"`
	Content map[string]interface{} `schema:"-" json:"content"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// New returns a fully configured client
func New(URL, Username, Password string, InsecureSkipVerify bool) *Client {
	c := &Client{}
	c.client = resty.New().
		SetBasicAuth(Username, Password).
		SetHostURL(fmt.Sprintf("%s", URL)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetQueryParam("output_mode", "json").
		SetMode("rest").
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: InsecureSkipVerify})

	return c
}

func (c *Client) Get(path string) (b []byte, e error) {
	r, e := c.client.R().
		Get(path)
	if e != nil {
		return
	}

	e = checkStatusCode(r)
	if e != nil {
		return
	}

	b = r.Body()
	return
}

func (c *Client) Post(path string, data url.Values) (b []byte, e error) {
	r, e := c.client.R().
		SetMultiValueFormData(data).
		Post(path)
	if e != nil {
		return
	}

	e = checkStatusCode(r)
	if e != nil {
		return
	}

	b = r.Body()
	return
}

func (c *Client) Delete(path string) (e error) {
	r, e := c.client.R().Delete(path)
	if e != nil {
		return
	}

	e = checkStatusCode(r)
	return
}

func checkStatusCode(r *resty.Response) (e error) {
	s := r.StatusCode()
	if !(s >= 200 && s <= 299) {
		eMsg := fmt.Sprintf("Unexpected response from Splunk: %d", s)
		if s >= 400 && s <= 499 {
			eMsg += fmt.Sprintf("\n%s", string(r.Body()))
		}
		e = errors.New(eMsg)
	}
	return
}
