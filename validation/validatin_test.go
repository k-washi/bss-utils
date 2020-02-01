package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailValidatin(t *testing.T) {
	testEmail := "abc@gmail.com"
	failEmails := []string{"<abc@gmial.com", "abc@~gmail.com", "abc@gmail.cccccc"}

	assert.Nil(t, Email(testEmail))

	for _, e := range failEmails {
		err := Email(e)
		assert.Error(t, err)
		fmt.Println(err)
	}
}
