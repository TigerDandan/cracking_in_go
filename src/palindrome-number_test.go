package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	result := isPalindrome(121)
	assert.Equal(t, true, result)

	result = isPalindrome(1221)
	assert.Equal(t, true, result)

	result = isPalindrome(1)
	assert.Equal(t, true, result)

	result = isPalindrome(-1)
	assert.Equal(t, false, result)

	result = isPalindrome(123454321)
	assert.Equal(t, true, result)
}
