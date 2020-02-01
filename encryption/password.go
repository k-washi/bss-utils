package encryption

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	md5Length = 32
)

var (
	salt1 string
	salt2 string
)

func init() {
	salt1 = os.Getenv("GO_CRYPTO_SALT1")
	salt2 = os.Getenv("GO_CRYPTO_SALT2")
	if salt1 == "" {
		fmt.Println("salt1 env valiable is not set")
		salt1 = "@#$%"
	}
	if salt2 == "" {
		fmt.Println("salt2 env valiable is not set")
		salt2 = "^&*()"
	}
}

//GetMD5 create md5 encryption
func GetMD5(arg1 string, arg2 string) string {
	h := md5.New()
	io.WriteString(h, arg1)
	pwmd5arg1 := fmt.Sprintf("%x", h.Sum(nil))

	io.WriteString(h, arg2)
	pwmd5arg2 := fmt.Sprintf("%x", h.Sum(nil))

	io.WriteString(h, salt1)
	io.WriteString(h, pwmd5arg1)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5arg2)

	return fmt.Sprintf("%x", h.Sum(nil))
}

//Validation md5 validation
func Validation(input string) error {
	if len(input) != md5Length {
		return errors.New("Md5 validation error: length is diff")
	}

	return nil
}
