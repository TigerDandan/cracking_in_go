package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
		result := twoSum([]int{2, 7, 3, 4}, 9)
		assert.Equal(t, []int{0, 1}, result)
}
