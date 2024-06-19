package strings

import "testing"

func TestReverseWords(t *testing.T) {
	out := reverseWords("  hello world  ")
	if out != "world hello" {
		t.Log("Output: ", out)
		t.Fail()
	}
}
