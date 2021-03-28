package scago

import "testing"

// TODO: actually test these, at the moment it's just to manually
// see what they're outputting
func TestNewRule(*testing.T) {
	NewRule("a > b / c")
	NewRule("x>y/z")
	NewRule(">b / d")
	NewRule("c>/_d")
}
