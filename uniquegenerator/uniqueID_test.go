package uniquegenerator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueID(t *testing.T) {
	base := make([]string, 100)

	for i := 0; i < 100; i++ {
		base[i] = GetUlid()
	}
	uniqueDlid := removeUlidDuplicate(base)

	assert.EqualValues(t, len(base), len(uniqueDlid))
	fmt.Println("not unique ulid:", getUlidDuplicateargs(base))
}

func removeUlidDuplicate(args []string) []string {
	results := make([]string, 0, len(args))
	uniqueEncoutered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !uniqueEncoutered[args[i]] {
			uniqueEncoutered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}

func getUlidDuplicateargs(args []string) []string {
	results := make([]string, 0, len(args))
	uniqueEncoutered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !uniqueEncoutered[args[i]] {
			uniqueEncoutered[args[i]] = true
		} else {
			results = append(results, args[i])
		}
	}
	return results
}
