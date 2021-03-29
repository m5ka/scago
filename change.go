package scago

import (
	"errors"
	"strconv"
	"strings"
)

// Change represents the set of sounds or categories
// that are to be changed to within a sound change rule.
// This often takes the form of a replacement but may also
// represent a movement.
type Change struct {
	replacement string
	movement    int
	deletion    bool
}

// ParseChange returns a Change object based on a given input
// string, corresponding to the change string as would be
// written in the scago sound change notation.
// Returns nil if there is no change, or an error if the
// change could not be parsed.
func (s *ScagoInstance) ParseChange(input string) (*Change, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return &Change{"", 0, true}, nil
	}
	change := &Change{}
	split := strings.Split(input, "@")
	if len(split) == 1 {
		change.replacement = input
	} else if len(split) >= 2 {
		change.replacement = strings.TrimSpace(split[0])
		i, err := strconv.Atoi(strings.TrimSpace(split[1]))
		if err != nil {
			return nil, err
		}
		change.movement = i
	} else {
		return nil, errors.New("too many '@' operators in change")
	}
	return change, nil
}
