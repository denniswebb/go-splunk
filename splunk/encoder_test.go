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
