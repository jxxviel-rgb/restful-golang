package test

import (
	"testing"

	"github.com/jxxviel-rgb/restful-golang/simple"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Fikri")
	assert.NotNil(t, connection)
	cleanup()
}
