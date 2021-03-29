package scago

import "errors"

type ScagoInstance struct {
	rules      *Rule     // a pointer to the first rule in the list
	categories *Category // a pointer to the first category in the list
}

// Applies the ScagoInstance's ruleset to the given word, returning
// the changed word and whether there were any errors. In case of
// errors, the returned string may be empty.
// TODO: implement this functionally
func (s *ScagoInstance) Apply(lemma string) (string, error) {
	return "", errors.New("function not yet implemented")
}

// Returns a new ScagoInstance that can be used to
// start doing sound changes!
func New() *ScagoInstance {
	return &ScagoInstance{}
}
