package uniquegenerator

import (
	"math/rand"

	"github.com/oklog/ulid/v2"

	"time"
)

var (
	t       = time.Unix(1000000, 0)
	entropy = ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
)

//generateUlid ex. output: 0000XSNJG0MQJHBF4QX1EFD6Y3
func generateUlid() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

//GetUlid ex. output: 0000XSNJG0MQJHBF4QX1EFD6Y3
func GetUlid() string {
	return generateUlid().String()
}
