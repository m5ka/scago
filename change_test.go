package scago

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseChange(t *testing.T) {
	s := New()
	t.Run("parse change: replacement", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseChange("a")
		if err != nil {
			t.Fatalf("%s returned error: %s", t.Name(), err)
		}
		assert.Equal(got.replacement, "a")
		assert.Equal(got.movement, 0)
		assert.False(got.deletion)
	})
	t.Run("parse change: movement", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseChange(" @ 6 ")
		if err != nil {
			t.Fatalf("%s returned error: %s", t.Name(), err)
		}
		assert.Empty(got.replacement)
		assert.Equal(got.movement, 6)
		assert.False(got.deletion)
	})
	t.Run("parse change: negative movement", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseChange("@ -2")
		if err != nil {
			t.Fatalf("%s returned err: %s", t.Name(), err)
		}
		assert.Empty(got.replacement)
		assert.Equal(got.movement, -2)
		assert.False(got.deletion)
	})
	t.Run("parse change: replacement and movement", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseChange("b @ 13 ")
		if err != nil {
			t.Fatalf("%s returned err: %s", t.Name(), err)
		}
		assert.Equal(got.replacement, "b")
		assert.Equal(got.movement, 13)
		assert.False(got.deletion)
	})
	t.Run("parse change: replacement and negative movement", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseChange("x@-3")
		if err != nil {
			t.Fatalf("%s returned err: %s", t.Name(), err)
		}
		assert.Equal(got.replacement, "x")
		assert.Equal(got.movement, -3)
		assert.False(got.deletion)
	})
	t.Run("parse change: deletion", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseChange(" ")
		if err != nil {
			t.Fatalf("%s returned err: %s", t.Name(), err)
		}
		assert.Empty(got.replacement)
		assert.Empty(got.movement)
		assert.True(got.deletion)
	})
}
