package scago

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	assert := assert.New(t)
	got, err := NewCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Fatalf("NewCategory returned error: %s", err)
	}
	assert.Equal(got.identifier, "K")
	assert.Equal(got.pattern, "(a|b|c)")
}
