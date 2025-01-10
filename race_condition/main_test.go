package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Increment(t *testing.T) {
	n := 0
	res := increment(&n, 10)

	assert.Equal(t, 10, res)
}

func Test_IncrementAtomic(t *testing.T) {
	var n32 int32
	res := incrementAtomic(&n32, 10)

	assert.Equal(t, 10, int(res))
}

func Test_IncrementMutex(t *testing.T) {
	n := 0
	res := incrementMutex(&n, 10)

	assert.Equal(t, 10, int(res))
}
