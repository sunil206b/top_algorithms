package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumBasic(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9

	result := twoSumBasic(input, target)
	expected := []int{0, 1}
	assert.Equal(t, result, expected)
}
