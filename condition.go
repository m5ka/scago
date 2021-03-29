package scago

import "strings"

// Condition represents a word's environment that can be
// checked against in order to conditionally perform
// sound changes on words. This can be used, for example,
// to perform sound changes only on sounds that appear
// with certain adjacent sounds or other word-environmental
// factors.
type Condition struct {
}

// ParseCondition returns a Condition based on the given input
// string, corresponding to the condition string as would be
// written in the scago sound change notation.
// Returns nil if there is no condition, or an error if the
// condition could not be parsed.
func (s *Scago) ParseCondition(input string) (*Condition, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}
	// TODO: parse condition instead of blank-returning nil always
	return nil, nil
}
