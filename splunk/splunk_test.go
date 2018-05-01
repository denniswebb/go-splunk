/* Upside Travel, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

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
