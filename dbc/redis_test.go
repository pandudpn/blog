package dbc_test

import (
	"testing"
	
	"github.com/pandudpn/blog/dbc"
	"github.com/stretchr/testify/assert"
)

func TestNewConnectionRedis(t *testing.T) {
	t.Run("Test Case #1 Error open connection Redis", func(t *testing.T) {
		var expectedErr error
		
		_, err := dbc.NewConnectionRedis()
		
		assert.Equal(t, expectedErr, err)
	})
}
