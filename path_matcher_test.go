package antpath

import "testing"

//TestMatch
func TestMatch(t *testing.T) {
	var matcher PathMatcher
	matcher = New()

	t.Log(matcher.Match("test","test"))
	t.Log(matcher.Match("/test","/test"))
	t.Log(matcher.Match("http://example.org", "http://example.org"))
	t.Log(matcher.Match("/test.jpg", "test.jpg"))

}