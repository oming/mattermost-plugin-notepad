package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	t.Run("get configuration", func(t *testing.T) {
		plugin := &Plugin{}
		fmt.Printf("Bookmark Content = %#v", plugin.getConfiguration().BookmarkContent)
		assert.NotNil(t, plugin.getConfiguration())
	})
}
