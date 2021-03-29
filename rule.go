package scago

import "regexp"

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
// adds it to the rules of this ScagoInstance.
func (s *ScagoInstance) AddRule(rule string) error {
	r, err := s.NewRule(rule)
	if err != nil {
		return err
	}
	s.rules.Append(r)
	return nil
}

// NewRule returns a new Rule object according to the given rule string.
// If the rule could not be parsed, it instead returns nil and an error.
func (s *ScagoInstance) NewRule(rule string) (*Rule, error) {
	re := regexp.MustCompile(`^(.*?)>(.*?)(?:/(.*?)(?:!(.*?)(?:/(.*?))?)?)?$`)
	parts := re.FindStringSubmatch(rule)

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
