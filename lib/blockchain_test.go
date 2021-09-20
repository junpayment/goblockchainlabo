package lib

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		b := NewBlock("test", FirstHash)
		assert.NotNil(t, b)
		log.Println(b)
	})
}
