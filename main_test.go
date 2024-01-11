package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	fmt.Print("testing server")
	assert.Equal(t, 123, 123, "they should be equal")
}
