package scago

import "testing"

func TestParseSimpleTarget(t *testing.T) {
	s := New()
	testStr := "s.ParseTarget('a ')"
	got, err := s.ParseTarget("a ")
	want := "^(a)"
	if err != nil {
		t.Errorf("%s returned error: %s", testStr, err)
	} else if got.pattern.String() != want {
		t.Errorf("%s = Target{%s}; want Target{%s}", testStr, got.pattern.String(), want)
	}
}

func TestParseMultipleSimpleTargets(t *testing.T) {
	s := New()
	testStr := "s.ParseTarget('a , b,c ')"
	got, err := s.ParseTarget("a , b,c ")
	want := "^(a|b|c)"
	if err != nil {
		t.Errorf("%s returned an error: %s", testStr, err)
	} else if got.pattern.String() != want {
		t.Errorf("%s = Target{%s}; want Target{%s}", testStr, got.pattern.String(), want)
	}
}

func TestParseCategoryTarget(t *testing.T) {
	s := New()
	err := s.AddCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Errorf("Error encountered when adding category: %s", err)
	}
	testStr := "(K = a|b|c) s.ParseTarget('K')"
	got, err := s.ParseTarget("K")
	want := "^((a|b|c))"
	if err != nil {
		t.Errorf("%s returned an error: %s", testStr, err)
	} else if got.pattern.String() != want {
		t.Errorf("%s = Target{%s}; want Target{%s}", testStr, got.pattern.String(), want)
	}
}

func TestParseMultipleCategoryTargets(t *testing.T) {
	s := New()
	err := s.AddCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Errorf("Error encountered when adding category: %s", err)
	}
	err = s.AddCategory("M", []string{"x", "y", "z"})
	if err != nil {
		t.Errorf("Error encountered when adding category: %s", err)
	}
	testStr := "(K = a|b|c; M = x|y|z) s.ParseTarget('M, K')"
	got, err := s.ParseTarget("M, K")
	want := "^((x|y|z)|(a|b|c))"
	if err != nil {
		t.Errorf("%s returned an error: %s", testStr, err)
	} else if got.pattern.String() != want {
		t.Errorf("%s = Target{%s}; want Target{%s}", testStr, got.pattern.String(), want)
	}
}

func TestParseMultipleMixedTargets(t *testing.T) {
	s := New()
	err := s.AddCategory("K", []string{"a", "b", "c"})
	if err != nil {
		t.Errorf("Error encountered when adding category: %s", err)
	}
	err = s.AddCategory("M", []string{"x", "y", "z"})
	if err != nil {
		t.Errorf("Error encountered when adding category: %s", err)
	}
	testStr := "(K = a|b|c; M = x|y|z) s.ParseTarget(' K, a , b,c,d,M, e')"
	got, err := s.ParseTarget(" K, a , b,c,d,M, e")
	want := "^((a|b|c)|a|b|c|d|(x|y|z)|e)"
	if err != nil {
		t.Errorf("%s returned an error: %s", testStr, err)
	} else if got.pattern.String() != want {
		t.Errorf("%s = Target{%s}; want Target{%s}", testStr, got.pattern.String(), want)
	}
}
