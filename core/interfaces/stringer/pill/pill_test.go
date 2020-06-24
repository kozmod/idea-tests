package pill

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPill(t *testing.T) {
	pk1 := Aspirin
	assert.Equal(t, "Aspirin", pk1.String())
}
