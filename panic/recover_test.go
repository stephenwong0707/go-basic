package panic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	result := parent(panicChild)
	assert.Equal(t, true, result)
	result = parent(child)
	assert.Equal(t, false, result)
}
