package antpath

import "testing"

//matchers
var matcher PathMatcher

func init(){
	matcher = New()
}

//TestIsPattern
func TestIsPattern(t *testing.T){
	t.Log(matcher.IsPattern(""))
	t.Log(matcher.IsPattern("*"))
	t.Log(matcher.IsPattern("?"))
	t.Log(matcher.IsPattern("{"))
	t.Log(matcher.IsPattern("*?"))
	t.Log(matcher.IsPattern("\\*"))
	t.Log(matcher.IsPattern("http://example.org"))
	//ant 标准语法
	t.Log(matcher.IsPattern("http://example.org?name=chao"))
}

//TestMatch
func TestMatch(t *testing.T) {
	t.Log(matcher.Match("test","test"))
	t.Log(matcher.Match("/test","/test"))
	t.Log(matcher.Match("http://example.org", "http://example.org"))
	t.Log(matcher.Match("/test.jpg", "test.jpg"))

}