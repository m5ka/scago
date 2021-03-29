package scago

import (
	"regexp"
	"strings"
)

// Target represents the target of a sound change, that is
// the sounds or broad categories of sounds that are to be
// targeted by a sound change.
// TODO: add more features to target e.g indexing (nth instance of target)
type Target struct {
	pattern *regexp.Regexp // the pattern represented by the target
}

// ParseTarget returns a Target object based on a given input
// string, corresponding to the target string as would be
// written in the scago sound change notation.
// Returns nil if there is no target or returns an error if
// the target could not be parsed.
func (s *Scago) ParseTarget(input string) (*Target, error) {
	// Trim spaces from the input string and return
	// nil if there is no target to parse
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}

	sb := &strings.Builder{}
	targets := strings.Split(input, ",")
	sb.WriteString("^(")
	for i, target := range targets {
		if i != 0 {
			sb.WriteString("|")
		}
		target = strings.TrimSpace(target)
		// Append the category's pattern to the string if the
		// target is a category identifier, otherwise just append
		// the target.
		c := s.GetCategory(target)
		if c != nil {
			sb.WriteString(c.pattern)
		} else {
			sb.WriteString(target)
		}
	}
	sb.WriteString(")")
	// Check the pattern compiles and return it as a Target if so
	pattern := sb.String()
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &Target{re}, nil
}
