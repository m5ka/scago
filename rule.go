package scago

import "regexp"

// The Rule object represents a sound change rule that can target
// a sound or set of sounds and imply a change under certain
// circumstances. The object forms part of a linked list via the
// next *Rule (which may be nil in case of being the last in the set).
type Rule struct {
	target      *Target    // target of the sound change
	change      *Change    // what's being changed to
	condition   *Condition // condition for change to take place
	exception   *Condition // exception to condition
	alternative *Change    // alternative change in case of exception
	repetition  int        // times to repeat change
	next        *Rule      // the next rule in the linked list
}

// Returns whether a rule is followed by another,
// i.e false if this is the last rule in the linked list.
func (r *Rule) HasNext() bool {
	return r.next != nil
}

// Appends a rule to the end of the linked list of rules.
func (r *Rule) Append(rule *Rule) {
	if r.HasNext() {
		r.next.Append(rule)
		return
	}
	r.next = rule
}

// Creates a new rule according to the given string and
// adds it to the rules of this ScagoInstance.
func (s *ScagoInstance) AddRule(rule string) {
	r := NewRule(rule)
	s.rules.Append(r)
}

// Return a new Rule object according to the given rule string.
func NewRule(rule string) *Rule {
	re := regexp.MustCompile(`^(.*?)>(.*?)(?:/(.*?)(?:!(.*?)(?:/(.*?))?)?)?$`)
	parts := re.FindStringSubmatch(rule)

	return &Rule{
		ParseTarget(parts[1]),
		ParseChange(parts[2]),
		ParseCondition(parts[3]),
		ParseCondition(parts[4]),
		ParseChange(parts[5]),
		1, nil, // TODO: parse for these values too
	}
}
