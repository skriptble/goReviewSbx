package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, "hello", "hello", "Hello should equal Hello")
	main()
	So(5, ShouldEqual, 5)
}
