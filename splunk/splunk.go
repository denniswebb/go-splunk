package splunk

import (
	"crypto/tls"
	"fmt"
	"net/url"

	"encoding/xml"

	"errors"

	"strings"

	"golang.org/x/tools/blog/atom"
	"gopkg.in/resty.v1"
)

const (
	PathSavedSearchCreate = "saved/searches"
	PathSavedSearch       = "saved/searches/%s"
	PathSavedSearchACL    = "saved/searches/%s/acl"
)

// Client communicates with the Splunk rest endpoint.
type Client struct {
	client *resty.Client
}

// Response is used to store Splunk response Atom feeds
type Response struct {
	Links     map[string]string `json:"links"`
	Origin    string            `json:"origin"`
	Updated   string            `json:"updated"`
	Entry    []Entry       `json:"entry"`
	Messages []string `json:"messages"`
}

// Entry is used to store Splunk response Atom entries
type Entry struct {
	Name    string                 `json:"name"`
	ID      string                 `json:"id"`
	Updated string                 `json:"updated"`
	Links   map[string]string      `json:"links"`
	Author  string                 `json:"author"`
	Content map[string]interface{} `json:"content"`
}


type Entry struct {
	atom.Entry
	Property []Property `xml:"content>dict>key"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
	Raw []byte `xml:",innerxml"`
}

type Dict struct {
	Property []Property `xml:"key"`
}

type List struct {
	Item []string `xml:"item"`
}

type RestError struct {
	Message string `xml:"messages>msg"`
}

// PropertyLookup returns the value of property from an Entry's Property slice
func (entry *Entry) PropertyLookup(name string) *Property {
	return propertyLookup(entry.Property, name)
}

func (dict *Dict) PropertyLookup(name string) *Property {
	return propertyLookup(dict.Property, name)
}

func propertyLookup(p []Property, n string) *Property{
	for _, v := range p {
		if v.Name == n {
			return &v
		}
	}
	return &Property{}
}

func (p *Property) Dict() (d Dict) {
	_ = xml.Unmarshal(p.Raw, &d)
	return
}

func (p *Property) List() (l List) {
	_ = xml.Unmarshal(p.Raw, &l)
	return
}

// PropertyMap returns a map[string][]string from an Entry's Property slice
func (entry *Entry) PropertyMap() (m map[string][]string) {
	return propertyMap(entry.Property)
}

func (dict *Dict) PropertyMap() (m map[string][]string) {
	return propertyMap(dict.Property)
}

func propertyMap(p []Property) (m map[string][]string){
	m = map[string][]string{}
	for _, v := range p {
		m[strings.Replace(v.Name, ".", "_", -1)] = []string{v.Value}
	}
	return
}

// New returns a fully configured client
func New(URL, Username, Password string, InsecureSkipVerify bool) *Client {
	c := &Client{}
	c.client = resty.New().
		SetBasicAuth(Username, Password).
		SetHostURL(fmt.Sprintf("%s/services/", URL)).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetMode("rest").
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: InsecureSkipVerify})

	return c
}

func (c *Client) Get(path string) (f Feed, e error) {
	r, e := c.client.R().
		SetResult(f).
		Get(path)
	if e != nil {
		return
	}

	e = checkStatusCode(r)
	if e != nil {
		return
	}

	e = xml.Unmarshal(r.Body(), &f)
	return
}

func (c *Client) Post(path string, data url.Values) (f Feed, e error) {
	r, e := c.client.R().
		SetMultiValueFormData(data).
		SetResult(f).
		Post(path)
	if e != nil {
		return
	}

	e = checkStatusCode(r)
	if e != nil {
		return
	}

	e = xml.Unmarshal(r.Body(), &f)
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
			re := RestError{}
			xml.Unmarshal(r.Body(), &re)
			eMsg += fmt.Sprintf("\n%s", re.Message)
		}
		e = errors.New(eMsg)
	}
	return
}
