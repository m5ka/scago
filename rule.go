package scago

import (
	"errors"
	"regexp"
)

// Rule represents a sound change rule that can target a sound or set
// of sounds and imply a change under certain circumstances.
// The object forms part of a linked list via the next *Rule,
// which may be nil in case of being the last in the set.
type Rule struct {
	target      *Target    // target of the sound change
	change      *Change    // what's being changed to
	condition   *Condition // condition for change to take place
	exception   *Condition // exception to condition
	alternative *Change    // alternative change in case of exception
	repetition  int        // times to repeat change
	next        *Rule      // the next rule in the linked list
}

// Apply applies this rule to the given word and returns the
// resulting word. In case of an error, this is returned
// alongside an empty result.
func (r *Rule) Apply(lemma string) (string, error) {
	w, err := NewWord(lemma)
	if err != nil {
		return "", err
	}
	for w.Next() {
		// Make sure the target matches (or no target) and take note
		// of target length if so, or skip if not.
		var t int
		if r.target != nil {
			t = w.MatchTarget(r.target.pattern)
			if t < 0 {
				continue
			}
		}
		// Check conditions and skip if any do not match
		if !w.CheckConditions(r.condition, t) {
			continue
		}
		// Based on exception/alternative, decide whether to change
		// to change or exception, or whether we need to skip (exception
		// with no alternative provided)
		change := r.change
		if r.exception != nil {
			if w.CheckConditions(r.exception, t) {
				if r.alternative != nil {
					change = r.alternative
				}
			} else {
				continue
			}
		}
		// Time to carry out the change!
		// change = (*Change) change to carry out
		// t      = (int) length in word to alter/move/etc
		err := w.Change(change, t)
		if err != nil {
			return "", err
		}
	}
	return w.String(), nil
}

// HasNext returns true if the given rule is followed by another
// and thus returns false if this is the last rule in the linked list.
func (r *Rule) HasNext() bool {
	return r.next != nil
}

// Append appends a rule to the end of the linked list of rules.
func (r *Rule) Append(rule *Rule) {
	if r.HasNext() {
		r.next.Append(rule)
		return
	}
	r.next = rule
}

// AddRule creates a new rule according to the given string and
// adds it to the rules list in s.
func (s *Scago) AddRule(rule string) error {
	r, err := s.NewRule(rule)
	if err != nil {
		return err
	}
	if s.rules == nil {
		s.rules = r
	} else {
		s.rules.Append(r)
	}
	return nil
}

// NewRule returns a new Rule object according to the given rule string.
// If the rule could not be parsed, it instead returns nil and an error.
func (s *Scago) NewRule(rule string) (*Rule, error) {
	re := regexp.MustCompile(`^(.*?)>(.*?)(?:/(.*?)(?:!(.*?)(?:/(.*?))?)?)?$`)
	parts := re.FindStringSubmatch(rule)
	if parts == nil {
		return nil, errors.New("rule does not parse")
	}

	target, err := s.ParseTarget(parts[1])
	if err != nil {
		return nil, err
	}
	change, err := s.ParseChange(parts[2])
	if err != nil {
		return nil, err
	}
	condition, err := s.ParseCondition(parts[3])
	if err != nil {
		return nil, err
	}
	exception, err := s.ParseCondition(parts[4])
	if err != nil {
		return nil, err
	}
	alternative, err := s.ParseChange(parts[5])
	if err != nil {
		return nil, err
	}

	return &Rule{
		target,
		change,
		condition,
		exception,
		alternative,
		1, nil, // TODO: parse for these values too
	}, nil
}
