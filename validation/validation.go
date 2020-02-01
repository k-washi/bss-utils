package validation

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	emailPattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

var (
	emailRegexp = regexp.MustCompile(emailPattern)
)

//Email validate email
func Email(e string) error {
	if !emailRegexp.MatchString(e) {
		return errors.New(fmt.Sprintln("Email address is invalid: ", emailPattern))
	}

	return nil
}
