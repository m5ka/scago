package scago

import "testing"

func TestExpandPatternToRegex(t *testing.T) {
	s := New()
	err := s.AddCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Errorf("error encountered when adding category")
	}
	testStr := "s.ExpandPatternToRegex('xyKz')"
	got, err := s.ExpandPatternToRegex("xyKz")
	want := "xy(a|b|c)z"
	if err != nil {
		t.Errorf("%s returned an error: %s", testStr, err)
	} else if got.String() != want {
		t.Errorf("%s = Regexp{%s}; want Regexp{%s}", testStr, got.String(), want)
	}
}
