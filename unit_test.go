package pingyin_test

import (
	"testing"

	"github.com/dashessmith/pingyin"
	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	assert.Equal(t, []string{"wo"}, pingyin.HP["我"])
	assert.Equal(t, []string{"ai"}, pingyin.HP["爱"])
	assert.Equal(t, []string{"ni"}, pingyin.HP["你"])
}
func Test_2(t *testing.T) {
	assert.Equal(t, []string{"hang", "xing"}, pingyin.HP["行"])
}
func Test_3(t *testing.T) {
	t.Logf("%v\n", pingyin.PH["ya"])
}
