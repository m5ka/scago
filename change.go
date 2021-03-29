package scago

import "strings"

// Change represents the set of sounds or categories
// that are to be changed to within a sound change rule.
// This often takes the form of a replacement but may also
// represent a movement.
type Change struct {
}

// ParseChange returns a Change object based on a given input
// string, corresponding to the change string as would be
// written in the scago sound change notation.
// Returns nil if there is no change, or an error if the
// change could not be parsed.
func (s *ScagoInstance) ParseChange(input string) (*Change, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}
	// TODO: parse target instead of blank-returning nil always
	return nil, nil
}
