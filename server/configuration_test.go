package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	t.Run("get configuration", func(t *testing.T) {
		plugin := &Plugin{}
		assert.NotNil(t, plugin.getConfiguration())
	})
}
