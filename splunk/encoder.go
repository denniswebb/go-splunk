package splunk

import (
	"net/url"

	"github.com/gorilla/schema"
)

var (
	decoder = schema.NewDecoder()
	encoder = schema.NewEncoder()
)

func init() {
	encoder.SetAliasTag("xml")
}

func encode(i interface{}) (r url.Values, e error) {
	r = url.Values{}
	e = encoder.Encode(i, r)
	return
}
