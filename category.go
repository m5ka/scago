package scago

// Represents a category that can be used by the sound change applier
// to make changes based on broad categories of sounds, e.g changing
// any consonant (category in target) or implementing a sound change
// when adjacent to a plosive (category in condition).
type Category struct {
	identifier string    // the identifying name of the category
	sounds     []string  // a list of sounds belonging to the category
	next       *Category // the next category in the linked list
}

// Returns whether the category is followed by another,
// i.e false if this is the last category in the linked list.
func (c *Category) HasNext() bool {
	return c.next != nil
}

// Appends a category to the end of the linked list of
// categories.
func (c *Category) Append(category *Category) {
	if c.HasNext() {
		c.next.Append(category)
		return
	}
	c.next = category
}

// Adds a new category with the given identifier and
// sounds to the scago instance.
func (s *ScagoInstance) AddCategory(identifier string, sounds []string) {
	category := NewCategory(identifier, sounds)
	s.categories.Append(category)
}

// Returns a new category with the given identifier and
// sounds.
func NewCategory(identifier string, sounds []string) *Category {
	return &Category{identifier, sounds, nil}
}
