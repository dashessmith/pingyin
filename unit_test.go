package pinyin_test

import (
	"testing"

	"github.com/dashessmith/pinyin"
	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	assert.Equal(t, []string{"wo"}, pinyin.HP["我"])
	assert.Equal(t, []string{"ai"}, pinyin.HP["爱"])
	assert.Equal(t, []string{"ni"}, pinyin.HP["你"])
}
func Test_2(t *testing.T) {
	assert.Equal(t, []string{"hang", "xing"}, pinyin.HP["行"])
}
func Test_3(t *testing.T) {
	t.Logf("%v\n", pinyin.PH["ya"])
}
