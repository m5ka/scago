package scago

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWord(t *testing.T) {
	assert := assert.New(t)
	got, err := NewWord(" pineapple")
	want := []string{"#", "p", "i", "n", "e", "a", "p", "p", "l", "e", "#"}
	if err != nil {
		t.Fatalf("NewWord('pineapple') returned error: %s", err)
	}
	assert.Equal(got.index, 0)
	assert.Equal(len(got.internal), len(want))
	for i := range want {
		if got.internal[i] != want[i] {
			t.Errorf("internal[%d] = %s; want %s", i, got.internal[i], want[i])
			return
		}
	}
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	assert.Equal(w.String(), "pineapple")
}

func TestBoundaryString(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	assert.Equal(w.BoundaryString(), "#pineapple#")
}

func TestSubstring(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	assert.True(w.Next())
	assert.Equal(w.Substring(), "pineapple")
	assert.True(w.Next())
	assert.Equal(w.Substring(), "ineapple")
	assert.True(w.Next())
	assert.Equal(w.Substring(), "neapple")
}

func TestPreString(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	assert.True(w.Next()) // #_pineapple
	assert.Equal(w.PreString(), "#")
	assert.True(w.Next()) // #p_ineapple
	assert.Equal(w.PreString(), "#p")
	assert.True(w.Next()) // #pi_neapple
	assert.Equal(w.PreString(), "#pi")
}

func TestPostString(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	assert.True(w.Next()) // #_ineapple
	assert.Equal(w.PostString(1), "ineapple#")
	assert.True(w.Next()) // #p_neapple
	assert.Equal(w.PostString(1), "neapple#")
	assert.True(w.Next()) // #pi_eapple
	assert.Equal(w.PostString(1), "eapple#")
	assert.True(w.Next()) // #pin_apple
	assert.Equal(w.PostString(1), "apple#")
}

func TestNext(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	for i := 1; w.Next(); i++ {
		assert.Equal(w.index, i)
	}
	assert.False(w.Next())
}

func TestMatchTarget(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	assert.True(w.Next()) // pineapple
	re1 := regexp.MustCompile(`^((o|p|q|r))`)
	re2 := regexp.MustCompile(`^((s|t|u|v))`)
	re3 := regexp.MustCompile(`^(pin)`)
	assert.Equal(w.MatchTarget(re1), 1)
	assert.Equal(w.MatchTarget(re2), -1)
	assert.Equal(w.MatchTarget(re3), 3)
	assert.True(w.Next()) // ineapple
	re4 := regexp.MustCompile(`^(in)`)
	re5 := regexp.MustCompile(`^((x|y|z))`)
	assert.Equal(w.MatchTarget(re4), 2)
	assert.Equal(w.MatchTarget(re5), -1)
}

func TestMatchGlobal(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	// #pineapple#
	re1 := regexp.MustCompile(`#pi`)
	re2 := regexp.MustCompile(`a`)
	re3 := regexp.MustCompile(`(x|y|z)`)
	re4 := regexp.MustCompile(`(a|b|c)pple#`)
	re5 := regexp.MustCompile(`(a|b|c)p(p|q|r)m`)
	assert.True(w.MatchGlobal(re1))
	assert.True(w.MatchGlobal(re2))
	assert.False(w.MatchGlobal(re3))
	assert.True(w.MatchGlobal(re4))
	assert.False(w.MatchGlobal(re5))
}

func TestMatchPost(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	assert.True(w.Next()) // #_ineapple#
	assert.True(w.Next()) // #p_neapple#
	re1 := regexp.MustCompile(`^ne`)
	re2 := regexp.MustCompile(`^une`)
	re3 := regexp.MustCompile(`^inea`)
	re4 := regexp.MustCompile(`^(l|m|n|o)e(a|e|i|o|u)pple#`)
	re5 := regexp.MustCompile(`^e#`)
	assert.True(w.MatchPost(re1, 1))
	assert.False(w.MatchPost(re2, 1))
	assert.False(w.MatchPost(re3, 1))
	assert.True(w.MatchPost(re4, 1))
	assert.False(w.MatchPost(re5, 1))
}

func TestMatchPre(t *testing.T) {
	assert := assert.New(t)
	w, err := NewWord("pineapple")
	if err != nil {
		t.Fatalf("NewWord returned error: %s", err)
	}
	assert.True(w.Next()) // #_pineapple#
	assert.True(w.Next()) // #p_ineapple#
	assert.True(w.Next()) // #pi_neapple#
	assert.True(w.Next()) // #pin_eapple#
	re1 := regexp.MustCompile(`in$`)
	re2 := regexp.MustCompile(`#pin$`)
	re3 := regexp.MustCompile(`#pim$`)
	re4 := regexp.MustCompile(`(a|e|i|o|u)n$`)
	assert.True(w.MatchPre(re1))
	assert.True(w.MatchPre(re2))
	assert.False(w.MatchPre(re3))
	assert.True(w.MatchPre(re4))
}
