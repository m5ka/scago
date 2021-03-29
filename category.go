package scago

import (
	"regexp"
	"strings"
)

// Category represents a category that can be used by the sound change applier
// to make changes based on broad categories of sounds, e.g changing
// any consonant (category in target) or implementing a sound change
// when adjacent to a plosive (category in condition).
type Category struct {
	identifier string    // the identifying name of the category
	pattern    string    // the sounds in the category as a regexp string
	next       *Category // the next category in the linked list
}

// HasNext returns true if c is followed by another category,
// thus false if this is the last category in the linked list.
func (c *Category) HasNext() bool {
	return c.next != nil
}

// Append appends a category to the end of the linked list of
// categories.
func (c *Category) Append(category *Category) {
	if c.HasNext() {
		c.next.Append(category)
		return
	}
	c.next = category
}

// GetCategory returns a Category that belongs to s ScagoInstance,
// and corresponds to the given category string, or nil if no such
// category exists.
func (s *ScagoInstance) GetCategory(identifier string) *Category {
	// Return nil in case of no categories at all
	if s.categories == nil {
		return nil
	}
	// Keep looping until we find the category we want
	category := s.categories
	for {
		if category.identifier == identifier {
			return category
		}
		if category.HasNext() {
			category = category.next
		} else {
			break
		}
	}
	return nil
}

// AddCategory creates a new category with the given identifier and sounds and
// adds it to s.
// Returns an error if an error was encountered.
func (s *ScagoInstance) AddCategory(identifier string, sounds []string) error {
	category, err := NewCategory(identifier, sounds)
	if err != nil {
		return err
	}
	s.categories.Append(category)
	return nil
}

// NewCategory returns a new category with the given identifier and sounds.
// If it encounters an error, the error is returned and the Category is nil.
func NewCategory(identifier string, sounds []string) (*Category, error) {
	// Construct the sounds as a regexp pattern (a|b|c|d|etc)
	sb := &strings.Builder{}
	sb.WriteString("(")
	for i, sound := range sounds {
		if i != 0 {
			sb.WriteString("|")
		}
		sb.WriteString(sound)
	}
	sb.WriteString(")")
	exp := sb.String()
	// Test if string compiles and return error if not
	_, err := regexp.Compile(exp)
	if err != nil {
		return nil, err
	}
	// Return category with the given regexp pattern
	return &Category{identifier, exp, nil}, nil
}
