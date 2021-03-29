package scago

import (
	"errors"
	"regexp"
	"strings"
)

// Condition represents a word's environment that can be
// checked against in order to conditionally perform
// sound changes on words. This can be used, for example,
// to perform sound changes only on sounds that appear
// with certain adjacent sounds or other word-environmental
// factors.
type Condition struct {
	global  bool           // true if condition is global (whole word pattern)
	pre     *regexp.Regexp // if local, pattern to check for before index
	post    *regexp.Regexp // if local, pattern to check for after index
	pattern *regexp.Regexp // if global, pattern to check for
	next    *Condition     // next condition in the linked list
}

func (c *Condition) HasNext() bool {
	return c.next != nil
}

func (c *Condition) Append(condition *Condition) {
	if c.HasNext() {
		c.next.Append(condition)
	} else {
		c.next = condition
	}
}

func (s *Scago) ExpandPatternToRegex(pattern string) (*regexp.Regexp, error) {
	pattern = strings.TrimSpace(pattern)
	chars := strings.Split(pattern, "")
	sb := &strings.Builder{}
	for _, c := range chars {
		c = strings.TrimSpace(c)
		if c == "" {
			continue
		}
		if cat := s.GetCategory(c); cat != nil {
			sb.WriteString(cat.pattern)
		} else {
			sb.WriteString(c)
		}
	}
	if sb.Len() == 0 {
		return nil, nil
	}
	if re, err := regexp.Compile(sb.String()); err != nil {
		return nil, err
	} else {
		return re, nil
	}
}

// ParseCondition returns a Condition based on the given input
// string, corresponding to the condition string as would be
// written in the scago sound change notation.
// Returns nil if there is no condition, or an error if the
// condition could not be parsed.
func (s *Scago) ParseCondition(input string) (*Condition, error) {
	// Trim and return no condition if blank
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}
	// Split conditions by comma for multiple and loop through,
	// making a chain of conditions
	split := strings.Split(input, ",")
	var conditions *Condition
	for _, cond := range split {
		c := &Condition{}
		cond = strings.TrimSpace(cond)
		// Ignore blank conditions
		if cond == "" {
			continue
		}
		// Determine global or local condition
		condSplit := strings.Split(cond, "_")
		if len(condSplit) == 1 {
			c.global = true
			pattern, err := s.ExpandPatternToRegex(cond)
			if err != nil {
				return nil, err
			}
			c.pattern = pattern
		} else if len(condSplit) == 2 {
			c.global = false
			re, err := s.ExpandPatternToRegex(condSplit[0])
			if err != nil {
				return nil, err
			}
			c.pre = re
			re, err = s.ExpandPatternToRegex(condSplit[1])
			if err != nil {
				return nil, err
			}
			c.post = re
		} else {
			return nil, errors.New("invalid condition")
		}
		if conditions == nil {
			conditions = c
		} else {
			conditions.Append(c)
		}
	}
	return conditions, nil
}
