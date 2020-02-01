package encryption

import (
	"fmt"
	"testing"

	"github.com/k-washi/bss-utils/uniquegenerator"

	"github.com/k-washi/bss-utils/utils"
	"github.com/tj/assert"
)

func TestGetMD5same(t *testing.T) {
	id := "donaDon^?gaXyz"
	password := "sksS0t3f36"
	testNum := 10000
	base := make([]string, testNum)

	for i := 0; i < testNum; i++ {
		base[i] = GetMD5(id, password)
	}
	uniqueDlid := utils.RemoveUlidDuplicate(base)

	assert.EqualValues(t, len(uniqueDlid), 1)
	assert.Nil(t, Validation(base[0]))
}

func TestGetMD5diff(t *testing.T) {
	id := "donaDon^?gaXyz"

	testNum := 10000
	base := make([]string, testNum)

	for i := 0; i < testNum; i++ {
		base[i] = GetMD5(id, uniquegenerator.Get())
	}
	uniqueDlid := utils.RemoveUlidDuplicate(base)

	assert.EqualValues(t, len(uniqueDlid), testNum)

	fmt.Println("Example:", base[0:5])
}
