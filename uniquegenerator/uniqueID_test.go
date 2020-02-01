package uniquegenerator

import (
	"fmt"
	"testing"

	"github.com/k-washi/bss-utils/utils"

	"github.com/stretchr/testify/assert"
)

func TestUniqueID(t *testing.T) {
	base := make([]string, 100)

	for i := 0; i < 100; i++ {
		base[i] = Get()
	}
	uniqueDlid := utils.RemoveUlidDuplicate(base)

	assert.EqualValues(t, len(base), len(uniqueDlid))
	fmt.Println("not unique ulid:", utils.GetUlidDuplicateargs(base))
	assert.Nil(t, Validation(base[0]))
}
