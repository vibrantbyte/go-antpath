package antpath

import (
	"regexp"
	"testing"
)

//TestNewStringMatcher
func TestNewStringMatcher(t *testing.T) {
	t.Log(regexp.MatchString("\\d","12122admin"))
	t.Log(regexp.MatchString("\\d","admin"))

	r := regexp.MustCompilePOSIX("123admin")
	r.
	t.Log(regexp.MustCompilePOSIX("123admin"))
}
