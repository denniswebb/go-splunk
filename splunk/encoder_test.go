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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	var person struct {
		Name  string `xml:"name"`
		Email string `xml:"email_address"`
	}

	person.Name = "foo"
	person.Email = "foo@bar.com"

	personMap, err := encode(person)

	assert.Nil(t, err, "Expected err to be nil.")
	assert.NotNil(t, personMap, "Expected personMap to not be nil.")
	assert.Equal(t, 2, len(personMap), "Expected 2 but got %d.", len(personMap))
	assert.Equal(t, 0, len(personMap["Name"]), "Expected 0 but got %d.", len(personMap["Name"]))
	assert.Equal(t, 1, len(personMap["name"]), "Expected 1 but got %d.", len(personMap["name"]))
	assert.Equal(t, person.Name, personMap["name"][0], "Expected %s but got %s.", person.Name, personMap["name"][0])
	assert.Equal(t, 0, len(personMap["Email"]), "Expected 0 but got %d.", len(personMap["Email"]))
	assert.Equal(t, 1, len(personMap["email_address"]), "Expected 1 but got %d.", len(personMap["email_address"]))
	assert.Equal(t, person.Email, personMap["email_address"][0], "Expected %s but got %s.", person.Email, personMap["email_address"][0])
}
