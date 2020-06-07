package pill

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPill(t *testing.T) {
	pk1 := Aspirin
	assert.Equal(t, "Aspirin", pk1.String())
}
