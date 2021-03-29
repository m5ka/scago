package scago

import "testing"

func TestNewCategory(t *testing.T) {
	testStr := "NewCategory('K', []string{'a', 'b', 'c'})"
	got, err := NewCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Errorf("%s returned error: %s", testStr, err)
	} else if got.identifier != "K" || got.pattern != "(a|b|c)" {
		t.Errorf("%s = {%s, %s}; want {K, '(a|b|c)'}", testStr, got.identifier, got.pattern)
	}
}
