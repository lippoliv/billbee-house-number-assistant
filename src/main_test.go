package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMainText(t *testing.T) {
	// Given

	// When
	text := mainText()

	// Then
	assert.Equal(t, "hello world", text, "The two words should be the same.")
}
