package antpath

import (
	"strings"
	"testing"
)

//TestAntPathMatcher_IsPattern
func TestAntPathMatcher_IsPattern(t *testing.T) {
	matcher := &AntPathMatcher{}

	t.Log(matcher.IsPattern("http://example.org"))
	t.Log(matcher.IsPattern("http://v1/*/example.org"))
	t.Log(matcher.IsPattern("http://v1/*/t*st/example.org"))
	t.Log(matcher.IsPattern("http://v1/t?st/example.org"))
}

func TestAntPathMatcher_Match(t *testing.T) {
	t.Log(strings.HasPrefix("d111","d"))
	t.Log(strings.HasPrefix("d111","1"))
}