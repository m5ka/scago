package scago

type ScagoInstance struct {
	rules      *Rule     // a pointer to the first rule in the list
	categories *Category // a pointer to the first category in the list
}

// Returns a new ScagoInstance that can be used to
// start doing sound changes!
func New() *ScagoInstance {
	return &ScagoInstance{}
}
