package splunk

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func TestPropertyLookup(t *testing.T) {

}

func TestPropertyMap(t *testing.T) {

}

func TestNew(t *testing.T) {

}

func TestGet(t *testing.T) {

}

func TestPost(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestCheckStatusCode(t *testing.T) {
	codes := map[int]int{
		1: http.StatusCreated,
		2: http.StatusOK,
		3: http.StatusInternalServerError,
		4: http.StatusNotFound,
		5: http.StatusTemporaryRedirect,
	}

	reqCount := 0

	testHandler := func(w http.ResponseWriter, req *http.Request) {
		reqCount += 1
		w.WriteHeader(codes[reqCount])
	}

	testServer := httptest.NewServer(http.HandlerFunc(testHandler))
	defer testServer.Close()

	for reqCount < 5 {
		resp, err := resty.R().Get(testServer.URL)

		if err != nil {
			log.Panic(err)
		}

		if reqCount <= 2 {
			assert.Nil(t, checkStatusCode(resp))
		} else {
			assert.NotNil(t, checkStatusCode(resp))
		}
	}
}
