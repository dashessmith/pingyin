package pingyin_test

import (
	"testing"

	"github.com/dashessmith/pingyin"
	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	assert.Equal(t, "wo", pingyin.HP["我"])
	assert.Equal(t, "ai", pingyin.HP["爱"])
	assert.Equal(t, "ni", pingyin.HP["你"])
}
