package sputter

import (
	"regexp"
	"testing"
)

func TestLiteral(t *testing.T) {
	testSputHundredEmoji(t, `abc`)
}

func TestCharClassSingle(t *testing.T) {
	testSputHundredEmoji(t, `[A]`)
}

func TestCharClassSingleRange(t *testing.T) {
	testSputHundredEmoji(t, `[0-9]`)
}

func TestCharClassMultipleRange(t *testing.T) {
	testSputHundredEmoji(t, `[A-Z0-9]`)
}

func TestConcat(t *testing.T) {
	testSputHundredEmoji(t, `A[0-9]`)
}

func TestRange(t *testing.T) {
	testSputHundredEmoji(t, `A{1,5}`)
}

func TestAlternate(t *testing.T) {
	testSputHundredEmoji(t, `A|B`)
}

func TestPlus(t *testing.T) {
	testSputHundredEmoji(t, `A+`)
}

func TestStar(t *testing.T) {
	testSputHundredEmoji(t, `A*`)
}

func TestQuest(t *testing.T) {
	testSputHundredEmoji(t, `A+`)
}

func TestAnyCharNotNL(t *testing.T) {
	testSputHundredEmoji(t, ".")
}

func TestBeginLine(t *testing.T) {
	// testSputHundredEmoji(t, "a^a")
}

func TestEndLine(t *testing.T) {
	// testSputHundredEmoji(t, "a$a")
}

func testSputHundredEmoji(t *testing.T, exp string) {
	for i := 0; i < 100; i++ {
		s, err := Gen(exp)
		if err != nil {
			t.Error(`error from Gen:`, err)
		}

		match, err := regexp.Match(exp, []byte(s))
		if !match {
			t.Errorf(`output string "%s" does not match expression "%s"`, s, exp)
		}

		if err != nil {
			t.Error(err)
		}
	}
}
