package scago

import "strings"

// Target represents the target of a sound change, that is
// the sounds or broad categories of sounds that are to be
// targeted by a sound change.
type Target struct {
}

// ParseTarget returns a Target object based on a given input
// string, corresponding to the target string as would be
// written in the scago sound change notation.
// Returns nil if there is no target or returns an error if
// the target could not be parsed.
func ParseTarget(input string) (*Target, error) {
	// Trim spaces from the input string and return
	// nil if there is no target to parse
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}
	// TODO: parse target instead of blank-returning nil always
	return nil, nil
}
