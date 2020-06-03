package emath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivCeil(t *testing.T) {
	assert.True(t, DivCeil(0, 4) == 0)
	assert.True(t, DivCeil(1, 4) == 1)
	assert.True(t, DivCeil(2, 4) == 1)
	assert.True(t, DivCeil(3, 4) == 1)
	assert.True(t, DivCeil(4, 4) == 1)
	assert.True(t, DivCeil(5, 4) == 2)
	assert.True(t, DivCeil(6, 4) == 2)
	assert.True(t, DivCeil(7, 4) == 2)
	assert.True(t, DivCeil(8, 4) == 2)
	assert.True(t, DivCeil(9, 4) == 3)
}

func TestDivRound(t *testing.T) {
	assert.True(t, DivRound(0, 4) == 0)
	assert.True(t, DivRound(1, 4) == 0)
	assert.True(t, DivRound(2, 4) == 1)
	assert.True(t, DivRound(3, 4) == 1)
	assert.True(t, DivRound(4, 4) == 1)
	assert.True(t, DivRound(5, 4) == 1)
	assert.True(t, DivRound(6, 4) == 2)
	assert.True(t, DivRound(7, 4) == 2)
	assert.True(t, DivRound(8, 4) == 2)
	assert.True(t, DivRound(9, 4) == 2)
}
