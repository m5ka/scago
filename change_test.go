package scago

import "testing"

func TestParseChangeReplacement(t *testing.T) {
	testStr := "s.ParseChange('a')"
	got, err := New().ParseChange("a")
	wantReplacement := "a"
	wantMovement := 0
	if err != nil {
		t.Errorf("%s returned error: %s", testStr, err)
	}
	if got.replacement != wantReplacement {
		t.Errorf("%s = Replacement: %s; want %s", testStr, got.replacement, wantReplacement)
	}
	if got.movement != wantMovement {
		t.Errorf("%s = Movement: %d; want %d", testStr, got.movement, wantMovement)
	}
}

func TestParseChangeMovement(t *testing.T) {
	testStr := "s.ParseChange(' @ 6 ')"
	got, err := New().ParseChange(" @ 6 ")
	wantReplacement := ""
	wantMovement := 6
	if err != nil {
		t.Errorf("%s returned error: %s", testStr, err)
	}
	if got.replacement != wantReplacement {
		t.Errorf("%s = Replacement: %s; want %s", testStr, got.replacement, wantReplacement)
	}
	if got.movement != wantMovement {
		t.Errorf("%s = Movement: %d; want %d", testStr, got.movement, wantMovement)
	}
}

func TestParseChangeMovementNegative(t *testing.T) {
	testStr := "s.ParseChange('@ -2')"
	got, err := New().ParseChange("@ -2")
	wantReplacement := ""
	wantMovement := -2
	if err != nil {
		t.Errorf("%s returned error: %s", testStr, err)
	}
	if got.replacement != wantReplacement {
		t.Errorf("%s = Replacement: %s; want %s", testStr, got.replacement, wantReplacement)
	}
	if got.movement != wantMovement {
		t.Errorf("%s = Movement: %d; want %d", testStr, got.movement, wantMovement)
	}
}

func TestParseChangeReplacementAndMovement(t *testing.T) {
	testStr := "s.ParseChange(' b @ 13 ')"
	got, err := New().ParseChange(" b @ 13 ")
	wantReplacement := "b"
	wantMovement := 13
	if err != nil {
		t.Errorf("%s returned error: %s", testStr, err)
	}
	if got.replacement != wantReplacement {
		t.Errorf("%s = Replacement: %s; want %s", testStr, got.replacement, wantReplacement)
	}
	if got.movement != wantMovement {
		t.Errorf("%s = Movement: %d; want %d", testStr, got.movement, wantMovement)
	}
}

func TestParseChangeReplacementAndMovementNegative(t *testing.T) {
	testStr := "s.ParseChange('x@-3')"
	got, err := New().ParseChange("x@-3")
	wantReplacement := "x"
	wantMovement := -3
	if err != nil {
		t.Errorf("%s returned error: %s", testStr, err)
	}
	if got.replacement != wantReplacement {
		t.Errorf("%s = Replacement: %s; want %s", testStr, got.replacement, wantReplacement)
	}
	if got.movement != wantMovement {
		t.Errorf("%s = Movement: %d; want %d", testStr, got.movement, wantMovement)
	}
}

func TestParseChangeDeletion(t *testing.T) {
	testStr := "s.ParseChange(' ')"
	got, err := New().ParseChange(" ")
	if err != nil {
		t.Errorf("%s returned error: %s", testStr, err)
	}
	if !got.deletion {
		t.Errorf("%s = deletion: false; want true", testStr)
	}
}
