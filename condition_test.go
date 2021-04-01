package scago

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandPatternToRegex(t *testing.T) {
	s := New()
	err := s.AddCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Fatalf("error when adding category")
	}
	t.Run("Expand xyKz non-initial non-final", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ExpandPatternToRegex("xyKz", false, false)
		if !assert.NoError(err) {
			return
		}
		assert.Equal(got.String(), "xy(a|b|c)z")
	})
	t.Run("Expand xyKz initial non-final", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ExpandPatternToRegex("xyKz", true, false)
		if !assert.NoError(err) {
			return
		}
		assert.Equal(got.String(), "^xy(a|b|c)z")
	})
	t.Run("Expand xyKz non-initial final", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ExpandPatternToRegex("xyKz", false, true)
		if !assert.NoError(err) {
			return
		}
		assert.Equal(got.String(), "xy(a|b|c)z$")
	})
	t.Run("Expand xyKz initial final", func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ExpandPatternToRegex("xyKz", true, true)
		if !assert.NoError(err) {
			return
		}
		assert.Equal(got.String(), "^xy(a|b|c)z$")
	})
}

func TestParseCondition(t *testing.T) {
	s := New()
	err := s.AddCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Errorf("error encountered when adding category")
	}
	err = s.AddCategory("M", []string{"x", "y", "z"})
	if err != nil {
		t.Errorf("error encountered when adding category")
	}
	t.Run(`ParseCondition("K_M, _Mn,ab")`, func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseCondition("K_M, _Mn,ab")
		if !assert.NoError(err) {
			return
		}
		assert.False(got.global)
		assert.Empty(got.pattern)
		assert.Equal(got.pre.String(), "(a|b|c)$")
		assert.Equal(got.post.String(), "^(x|y|z)")
		if !assert.NotNil(got.next) {
			return
		}
		got = got.next
		assert.False(got.global)
		assert.Empty(got.pattern)
		assert.Nil(got.pre)
		assert.Equal(got.post.String(), "^(x|y|z)n")
		if !assert.NotNil(got.next) {
			return
		}
		got = got.next
		assert.True(got.global)
		assert.Nil(got.pre)
		assert.Nil(got.post)
		assert.Equal(got.pattern.String(), "ab")
		assert.Nil(got.next)
	})
	t.Run(`ParseCondition(" MMaa,#p_t")`, func(t *testing.T) {
		assert := assert.New(t)
		got, err := s.ParseCondition(" MMaa,#p_t")
		if !assert.NoError(err) {
			return
		}
		assert.True(got.global)
		assert.Nil(got.pre)
		assert.Nil(got.post)
		assert.Equal(got.pattern.String(), "(x|y|z)(x|y|z)aa")
		if !assert.NotNil(got.next) {
			return
		}
		got = got.next
		assert.False(got.global)
		assert.Nil(got.pattern)
		assert.Equal(got.pre.String(), "#p$")
		assert.Equal(got.post.String(), "^t")
		assert.Nil(got.next)
	})
}
