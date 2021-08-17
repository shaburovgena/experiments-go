package apiserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPIServer_ConfigTest(t *testing.T) {
	conf := newConfig()
	assert.NotNil(t, conf)
	assert.Equal(t, "debug", conf.LogLevel)
}
