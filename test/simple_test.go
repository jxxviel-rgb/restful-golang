package test

import (
	"testing"

	"github.com/jxxviel-rgb/restful-golang/simple"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, error := simple.InitializedService(false)
	assert.Nil(t, error)
	assert.NotNil(t, simpleService)
}
func TestSimpleServiceFailed(t *testing.T) {
	simpleService, error := simple.InitializedService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, error)
}
