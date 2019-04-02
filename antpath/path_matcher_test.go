package antpath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//matchers
var matcher PathMatcher

func init(){
	matcher = New()
}

//TestIsPattern
func TestIsPattern(t *testing.T){
	/**
	 规则：* | ** | ？
	 */
	// ---true---
	t.Log(matcher.IsPattern(""))
	t.Log(matcher.IsPattern("http://example.org"))
	t.Log(matcher.IsPattern("{"))

	// ---false---
	t.Log(matcher.IsPattern("*"))
	t.Log(matcher.IsPattern("?"))
	t.Log(matcher.IsPattern("*?"))
	t.Log(matcher.IsPattern("\\*"))
	//ant 标准语法 ？ * **
	t.Log(matcher.IsPattern("http://example.org?name=chao"))
	t.Log(matcher.IsPattern("/app/*.x"))
	t.Log(matcher.IsPattern("/app/p?ttern"))
	t.Log(matcher.IsPattern("/**/example"))
	t.Log(matcher.IsPattern("/app/**/dir/file."))
	t.Log(matcher.IsPattern("/**/*.jsp"))
}

//TestMatch
func TestMatch(t *testing.T) {
	// test exact matching
	//assert.True(t,matcher.Match("test", "test"))
	//assert.True(t,matcher.Match("/test", "/test"))
	//assert.True(t,matcher.Match("http://example.org", "http://example.org")) // SPR-14141
	//assert.False(t,matcher.Match("/test.jpg", "test.jpg"))
	//assert.False(t,matcher.Match("test", "/test"))
	//assert.False(t,matcher.Match("/test", "test"))
	//
	//// test matching with ?'s
	//assert.True(t,matcher.Match("t?st", "test"))
	//assert.True(t,matcher.Match("??st", "test"))
	//assert.True(t,matcher.Match("tes?", "test"))
	//assert.True(t,matcher.Match("te??", "test"))
	//assert.True(t,matcher.Match("?es?", "test"))
	//assert.False(t,matcher.Match("tes?", "tes"))
	assert.False(t,matcher.Match("tes?", "testt"))
	assert.False(t,matcher.Match("tes?", "tsst"))
	//
	//// test matching with *'s
	//assert.True(t,matcher.Match("*", "test"))
	//assert.True(t,matcher.Match("test*", "test"))
	//assert.True(t,matcher.Match("test*", "testTest"))
	//assert.True(t,matcher.Match("test/*", "test/Test"))
	//assert.True(t,matcher.Match("test/*", "test/t"))
	//assert.True(t,matcher.Match("test/*", "test/"))
	//assert.True(t,matcher.Match("*test*", "AnothertestTest"))
	//assert.True(t,matcher.Match("*test", "Anothertest"))
	//assert.True(t,matcher.Match("*.*", "test."))
	//assert.True(t,matcher.Match("*.*", "test.test"))
	//assert.True(t,matcher.Match("*.*", "test.test.test"))
	//assert.True(t,matcher.Match("test*aaa", "testblaaaa"))
	//assert.False(t,matcher.Match("test*", "tst"))
	//assert.False(t,matcher.Match("test*", "tsttest"))
	//assert.False(t,matcher.Match("test*", "test/"))
	//assert.False(t,matcher.Match("test*", "test/t"))
	//assert.False(t,matcher.Match("test/*", "test"))
	//assert.False(t,matcher.Match("*test*", "tsttst"))
	//assert.False(t,matcher.Match("*test", "tsttst"))
	//assert.False(t,matcher.Match("*.*", "tsttst"))
	//assert.False(t,matcher.Match("test*aaa", "test"))
	//assert.False(t,matcher.Match("test*aaa", "testblaaab"))
	//
	//// test matching with ?'s and /'s
	//assert.True(t,matcher.Match("/?", "/a"))
	//assert.True(t,matcher.Match("/?/a", "/a/a"))
	//assert.True(t,matcher.Match("/a/?", "/a/b"))
	//assert.True(t,matcher.Match("/??/a", "/aa/a"))
	//assert.True(t,matcher.Match("/a/??", "/a/bb"))
	//assert.True(t,matcher.Match("/?", "/a"))
	//
	//// test matching with **'s
	//assert.True(t,matcher.Match("/**", "/testing/testing"))
	//assert.True(t,matcher.Match("/*/**", "/testing/testing"))
	//assert.True(t,matcher.Match("/**/*", "/testing/testing"))
	//assert.True(t,matcher.Match("/bla/**/bla", "/bla/testing/testing/bla"))
	//assert.True(t,matcher.Match("/bla/**/bla", "/bla/testing/testing/bla/bla"))
	//assert.True(t,matcher.Match("/**/test", "/bla/bla/test"))
	//assert.True(t,matcher.Match("/bla/**/**/bla", "/bla/bla/bla/bla/bla/bla"))
	//assert.True(t,matcher.Match("/bla*bla/test", "/blaXXXbla/test"))
	//assert.True(t,matcher.Match("/*bla/test", "/XXXbla/test"))
	//assert.False(t,matcher.Match("/bla*bla/test", "/blaXXXbl/test"))
	//assert.False(t,matcher.Match("/*bla/test", "XXXblab/test"))
	//assert.False(t,matcher.Match("/*bla/test", "XXXbl/test"))
	//
	//assert.False(t,matcher.Match("/????", "/bala/bla"))
	//assert.False(t,matcher.Match("/**/*bla", "/bla/bla/bla/bbb"))
	//
	//assert.True(t,matcher.Match("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing/"))
	//assert.True(t,matcher.Match("/*bla*/**/bla/*", "/XXXblaXXXX/testing/testing/bla/testing"))
	//assert.True(t,matcher.Match("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing"))
	//assert.True(t,matcher.Match("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing.jpg"))
	//
	//assert.True(t,matcher.Match("*bla*/**/bla/**", "XXXblaXXXX/testing/testing/bla/testing/testing/"))
	//assert.True(t,matcher.Match("*bla*/**/bla/*", "XXXblaXXXX/testing/testing/bla/testing"))
	//assert.True(t,matcher.Match("*bla*/**/bla/**", "XXXblaXXXX/testing/testing/bla/testing/testing"))
	//assert.False(t,matcher.Match("*bla*/**/bla/*", "XXXblaXXXX/testing/testing/bla/testing/testing"))
	//
	//assert.False(t,matcher.Match("/x/x/**/bla", "/x/x/x/"))
	//
	//assert.True(t,matcher.Match("/foo/bar/**", "/foo/bar"))
	//
	//assert.True(t,matcher.Match("", ""))
	//
	//assert.True(t,matcher.Match("/{bla}.*", "/testing.html"))

}

//TestMatchStart
func TestMatchStart(t *testing.T){

}

//TestExtractPathWithinPattern
func TestExtractPathWithinPattern(t *testing.T){

}

//TestExtractUriTemplateVariables
func TestExtractUriTemplateVariables(t *testing.T)  {

}

//TestGetPatternComparator
func TestGetPatternComparator(t *testing.T){

}

//TestCombine
func TestCombine(t *testing.T){
	t.Log("TestCombine beginning...")

	assert.Equal(t,"",matcher.Combine("", ""))
	assert.Equal(t,"/hotels", matcher.Combine("/hotels", ""))
	assert.Equal(t,"/hotels", matcher.Combine("", "/hotels"))
	assert.Equal(t,"/hotels/booking", matcher.Combine("/hotels/*", "booking"))
	assert.Equal(t,"/hotels/booking", matcher.Combine("/hotels/*", "/booking"))
	assert.Equal(t,"/hotels/**/booking", matcher.Combine("/hotels/**", "booking"))
	assert.Equal(t,"/hotels/**/booking", matcher.Combine("/hotels/**", "/booking"))
	assert.Equal(t,"/hotels/booking", matcher.Combine("/hotels", "/booking"))
	assert.Equal(t,"/hotels/booking", matcher.Combine("/hotels", "booking"))
	assert.Equal(t,"/hotels/booking", matcher.Combine("/hotels/", "booking"))
	assert.Equal(t,"/hotels/{hotel}", matcher.Combine("/hotels/*", "{hotel}"))
	assert.Equal(t,"/hotels/**/{hotel}", matcher.Combine("/hotels/**", "{hotel}"))
	assert.Equal(t,"/hotels/{hotel}", matcher.Combine("/hotels", "{hotel}"))
	assert.Equal(t,"/hotels/{hotel}.*", matcher.Combine("/hotels", "{hotel}.*"))
	assert.Equal(t,"/hotels/*/booking/{booking}", matcher.Combine("/hotels/*/booking", "{booking}"))
	//
	assert.Equal(t,"/hotel.html", matcher.Combine("/*.html", "/hotel.html"))
	assert.Equal(t,"/hotel.html", matcher.Combine("/*.html", "/hotel"))
	assert.Equal(t,"/hotel.html", matcher.Combine("/*.html", "/hotel.*"))
	assert.Equal(t,"/*.html", matcher.Combine("/**", "/*.html"))
	assert.Equal(t,"/*.html", matcher.Combine("/*", "/*.html"))
	assert.Equal(t,"/*.html", matcher.Combine("/*.*", "/*.html"))
	assert.Equal(t,"/{foo}/bar", matcher.Combine("/{foo}", "/bar"))    // SPR-8858
	assert.Equal(t,"/user/user", matcher.Combine("/user", "/user"))    // SPR-7970
	assert.Equal(t,"/{foo:.*[^0-9].*}/edit/", matcher.Combine("/{foo:.*[^0-9].*}", "/edit/")) // SPR-10062
	assert.Equal(t,"/1.0/foo/test", matcher.Combine("/1.0", "/foo/test")) // SPR-10554
	assert.Equal(t,"/hotel", matcher.Combine("/", "/hotel")) // SPR-12975
	assert.Equal(t,"/hotel/booking", matcher.Combine("/hotel/", "/booking")) // SPR-12975

	t.Log("TestCombine ended...")

}