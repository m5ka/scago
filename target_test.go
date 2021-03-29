package scago

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTarget(t *testing.T) {
	s := New()
	err := s.AddCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Fatalf("error encountered when adding category")
	}
	err = s.AddCategory("M", []string{"x", "y", "z"})
	if err != nil {
		t.Fatalf("error encountered when adding category")
	}
	t.Run("Single literal target", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseTarget("a ")
		assert.NoError(err)
		assert.Equal(got.pattern.String(), "^(a)")
	})
	t.Run("Multiple literal targets", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseTarget("a , b,c ")
		assert.NoError(err)
		assert.Equal(got.pattern.String(), "^(a|b|c)")
	})
	t.Run("Single category target", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseTarget("K")
		assert.NoError(err)
		assert.Equal(got.pattern.String(), "^((a|b|c))")
	})
	t.Run("Multiple category targets", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseTarget("M, K")
		assert.NoError(err)
		assert.Equal(got.pattern.String(), "^((x|y|z)|(a|b|c))")
	})
	t.Run("Multiple mixed targets", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseTarget(" K, a , b,c,d,M, e")
		assert.NoError(err)
		assert.Equal(got.pattern.String(), "^((a|b|c)|a|b|c|d|(x|y|z)|e)")
	})
}
