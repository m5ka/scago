package scago

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Word represents a word as can be manipulated by sound changes,
// allowing functionality such as look-ahead and target matching.
// It also handles word-initial and word-final checks by appending
// and prepending the # character to the word in Word's internal
// representation. Word also has an internal counter to keep track
// of how far along the word checking/compilation is.
type Word struct {
	internal []string
	index    int
}

// CheckConditions loops through a linked list of conditions
// and checks if they apply, returning true if so and false if not.
func (w *Word) CheckConditions(c *Condition, length int) bool {
	for ; c != nil; c = c.next {
		if c.global && !w.MatchGlobal(c.pattern) {
			return false
		} else {
			if c.pre != nil && !w.MatchPre(c.pre) {
				return false
			}
			if c.post != nil && !w.MatchPost(c.post, length) {
				return false
			}
		}
	}
	return true
}

// Change changes w according to the given parameters. The Change
// determines the type of change and the length int determines how
// many characters from the current index in the word needs to be
// altered.
func (w *Word) Change(change *Change, length int) error {
	original := make([]string, len(w.internal))
	copy(original, w.internal)
	if change.deletion {
		// Delete length amount of characters from current index in word
		w.internal = original[:w.index]
		w.internal = append(w.internal, original[w.index+length:]...)
		// if we don't do the below, the next Next() will skip past the
		// first character after deletion
		w.index--
	} else {
		// Trim movement if we're too close to the end of the word
		movement := change.movement
		if movement != 0 && (w.index+movement < 1 || w.index+movement+1 >= len(w.internal)) {
			if movement < 0 {
				movement = (-w.index) + 1
			} else {
				movement = len(w.internal) - w.index - 2
			}
		}
		// Find original target for movement, or replacement for replacement
		var replacement string
		if change.replacement != "" {
			replacement = change.replacement
		} else {
			replacement = strings.Join(w.internal[w.index:w.index+length], "")
		}
		// Rebuild the internal representation from our copy of its previous state,
		// based on the movement and replacements that need to happen for this Change.
		if movement < 0 {
			w.internal = nil
			w.internal = append(w.internal, original[:w.index+movement]...)
			w.internal = append(w.internal, replacement)
			w.internal = append(w.internal, original[w.index+movement:w.index+movement+length]...)
			w.internal = append(w.internal, original[w.index+length:]...)
			fmt.Println("internal (after) =", strings.Join(w.internal, ""))
		} else {
			w.internal = nil
			w.internal = append(w.internal, original[:w.index]...)
			w.internal = append(w.internal, original[w.index+length:w.index+movement+length]...)
			w.internal = append(w.internal, replacement)
			w.internal = append(w.internal, original[w.index+length+movement:]...)
			// Stop it from unintentionally moving this again by finding it next iteration
			// NB: this does mean that in e.g "apopiiii" with p>@4, the second 'p' will be
			// ignored. This is not good but can be fixed in future - it seems like a fairly
			// slim use case.
			w.index += movement
		}
	}
	return nil
}

// MatchTarget checks whether the given regexp expression matches
// against the current subsection of the word (without checking
// boundary markers). If it does, returns the length of the match
// and if not returns -1.
func (w *Word) MatchTarget(re *regexp.Regexp) int {
	match := re.FindStringIndex(w.Substring())
	if match == nil {
		return -1
	} else if match[0] == 0 {
		return match[1]
	}
	return -1
}

// MatchGlobal checks whether the given regexp expression matches
// the entire word. Returns true if so, and false if not.
func (w *Word) MatchGlobal(re *regexp.Regexp) bool {
	return re.MatchString(w.BoundaryString())
}

// MatchPre checks whether the given regexp expression matches the
// portion of the word (including boundary markers) prior to the
// current index.
func (w *Word) MatchPre(re *regexp.Regexp) bool {
	pre := w.PreString()
	if pre == "" {
		return false
	}
	return re.MatchString(pre)
}

// MatchPost checks whether the given regexp expression matches the
// portion of the word (including boundary markers) after the current
// index.
func (w *Word) MatchPost(re *regexp.Regexp, length int) bool {
	post := w.PostString(length)
	if post == "" {
		return false
	}
	return re.MatchString(post)
}

// Next increments w's internal index and returns a bool which is
// true if the resulting current first character of the internal
// word is valid (i.e not the end of the word or a word boundary)
func (w *Word) Next() bool {
	w.index++
	if w.index >= len(w.internal) || w.internal[w.index] == "#" {
		return false
	}
	return true
}

// Substring returns the Word as a substring, starting from the current
// index and ending before the word boundary. Returns an empty string
// if the internal index has passed all characters in the word.
func (w *Word) Substring() string {
	if w.index < len(w.internal)-1 {
		return strings.Join(w.internal[w.index:len(w.internal)-1], "")
	}
	return ""
}

// PreString returns the Word as a substring (with boundary markers)
// prior to and not including the current index.
func (w *Word) PreString() string {
	if w.index >= len(w.internal) {
		return ""
	}
	return strings.Join(w.internal[:w.index], "")
}

// PostString returns the Word as a substring (with boundary markers)
// after and not including the current index
func (w *Word) PostString(length int) string {
	if w.index+length >= len(w.internal) {
		return ""
	}
	return strings.Join(w.internal[w.index+length:], "")
}

// String returns the Word as a full string, without the word-boundary
// markers (#). In an unchanged word, this is the equivalent of
// accessing the original string given to the constructor/
func (w *Word) String() string {
	return strings.Join(w.internal[1:len(w.internal)-1], "")
}

// BoundaryString returns the entire word including boundary markers (#)
// as a string.
func (w *Word) BoundaryString() string {
	return strings.Join(w.internal, "")
}

// NewWord returns a new Word object based on the given word as a
// string. It automatically prepends and appends the # marker.
func NewWord(lemma string) (*Word, error) {
	lemma = strings.TrimSpace(lemma)
	if lemma == "" {
		return nil, errors.New("empty word given")
	}
	var sb strings.Builder
	sb.WriteString("#")
	sb.WriteString(lemma)
	sb.WriteString("#")
	return &Word{strings.Split(sb.String(), ""), 0}, nil
}
