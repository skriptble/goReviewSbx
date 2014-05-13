package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, "hello", "hello", "Hello should equal Hello")
	main()
}
