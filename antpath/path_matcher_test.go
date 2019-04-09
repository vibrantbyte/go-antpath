package antpath

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestMatchLog(t *testing.T){
	t.Log(matcher.Match("tes?", "test"))
	t.Log(matcher.Match("/hotels/{hotel}", "/hotels/1"))
	t.Log(matcher.Match("tes?","tes"))
	t.Log(matcher.Match("tes?", "testt"))
	t.Log(matcher.Match("tes?", "tsst"))
}

//TestMatch
func TestMatch(t *testing.T) {
	// test exact matching
	assert.True(t,matcher.Match("test", "test"))
	assert.True(t,matcher.Match("/test", "/test"))
	assert.True(t,matcher.Match("http://example.org", "http://example.org")) // SPR-14141
	assert.False(t,matcher.Match("/test.jpg", "test.jpg"))
	assert.False(t,matcher.Match("test", "/test"))
	assert.False(t,matcher.Match("/test", "test"))
//
//	// test matching with ?'s
	assert.True(t,matcher.Match("t?st", "test"))
	assert.True(t,matcher.Match("??st", "test"))
	assert.True(t,matcher.Match("tes?", "test"))
	assert.True(t,matcher.Match("te??", "test"))
	assert.True(t,matcher.Match("?es?", "test"))
	assert.False(t,matcher.Match("tes?", "tes"))
	assert.False(t,matcher.Match("tes?", "testt"))
	assert.False(t,matcher.Match("tes?", "tsst"))
	//
	//test matching with *'s
	assert.True(t,matcher.Match("*", "test"))
	assert.True(t,matcher.Match("test*", "test"))
	assert.True(t,matcher.Match("test*", "testTest"))
	assert.True(t,matcher.Match("test/*", "test/Test"))
	assert.True(t,matcher.Match("test/*", "test/t"))
	assert.True(t,matcher.Match("test/*", "test/"))
	assert.True(t,matcher.Match("*test*", "AnothertestTest"))
	assert.True(t,matcher.Match("*test", "Anothertest"))
	assert.True(t,matcher.Match("*.*", "test."))
	assert.True(t,matcher.Match("*.*", "test.test"))
	assert.True(t,matcher.Match("*.*", "test.test.test"))
	assert.True(t,matcher.Match("test*aaa", "testblaaaa"))
	assert.False(t,matcher.Match("test*", "tst"))
	assert.False(t,matcher.Match("test*", "tsttest"))
	assert.False(t,matcher.Match("test*", "test/"))
	assert.False(t,matcher.Match("test*", "test/t"))
	assert.False(t,matcher.Match("test/*", "test"))
	assert.False(t,matcher.Match("*test*", "tsttst"))
	assert.False(t,matcher.Match("*test", "tsttst"))
	assert.False(t,matcher.Match("*.*", "tsttst"))
	assert.False(t,matcher.Match("test*aaa", "test"))
	assert.False(t,matcher.Match("test*aaa", "testblaaab"))
	//
	// test matching with ?'s and /'s
	assert.True(t,matcher.Match("/?", "/a"))
	assert.True(t,matcher.Match("/?/a", "/a/a"))
	assert.True(t,matcher.Match("/a/?", "/a/b"))
	assert.True(t,matcher.Match("/??/a", "/aa/a"))
	assert.True(t,matcher.Match("/a/??", "/a/bb"))
	assert.True(t,matcher.Match("/?", "/a"))
	//
	// test matching with **'s
	assert.True(t,matcher.Match("/**", "/testing/testing"))
	assert.True(t,matcher.Match("/*/**", "/testing/testing"))
	assert.True(t,matcher.Match("/**/*", "/testing/testing"))
	assert.True(t,matcher.Match("/bla/**/bla", "/bla/testing/testing/bla"))
	assert.True(t,matcher.Match("/bla/**/bla", "/bla/testing/testing/bla/bla"))
	assert.True(t,matcher.Match("/**/test", "/bla/bla/test"))
	assert.True(t,matcher.Match("/bla/**/**/bla", "/bla/bla/bla/bla/bla/bla"))
	assert.True(t,matcher.Match("/bla*bla/test", "/blaXXXbla/test"))
	assert.True(t,matcher.Match("/*bla/test", "/XXXbla/test"))
	assert.False(t,matcher.Match("/bla*bla/test", "/blaXXXbl/test"))
	assert.False(t,matcher.Match("/*bla/test", "XXXblab/test"))
	assert.False(t,matcher.Match("/*bla/test", "XXXbl/test"))
	//
	assert.False(t,matcher.Match("/????", "/bala/bla"))
	assert.False(t,matcher.Match("/**/*bla", "/bla/bla/bla/bbb"))
	//
	assert.True(t,matcher.Match("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing/"))
	assert.True(t,matcher.Match("/*bla*/**/bla/*", "/XXXblaXXXX/testing/testing/bla/testing"))
	assert.True(t,matcher.Match("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing"))
	assert.True(t,matcher.Match("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing.jpg"))

	assert.True(t,matcher.Match("*bla*/**/bla/**", "XXXblaXXXX/testing/testing/bla/testing/testing/"))
	assert.True(t,matcher.Match("*bla*/**/bla/*", "XXXblaXXXX/testing/testing/bla/testing"))
	assert.True(t,matcher.Match("*bla*/**/bla/**", "XXXblaXXXX/testing/testing/bla/testing/testing"))
	assert.False(t,matcher.Match("*bla*/**/bla/*", "XXXblaXXXX/testing/testing/bla/testing/testing"))

	assert.False(t,matcher.Match("/x/x/**/bla", "/x/x/x/"))

	assert.True(t,matcher.Match("/foo/bar/**", "/foo/bar"))

	assert.True(t,matcher.Match("", ""))

	assert.True(t,matcher.Match("/{bla}.*", "/testing.html"))

}

//TestMatchStart
func TestMatchStart(t *testing.T){
	
	// test exact matching
	assert.True(t,matcher.MatchStart("test", "test"))
	assert.True(t,matcher.MatchStart("/test", "/test"))
	assert.False(t,matcher.MatchStart("/test.jpg", "test.jpg"))
	assert.False(t,matcher.MatchStart("test", "/test"))
	assert.False(t,matcher.MatchStart("/test", "test"))

	// test matching with ?'s
	assert.True(t,matcher.MatchStart("t?st", "test"))
	assert.True(t,matcher.MatchStart("??st", "test"))
	assert.True(t,matcher.MatchStart("tes?", "test"))
	assert.True(t,matcher.MatchStart("te??", "test"))
	assert.True(t,matcher.MatchStart("?es?", "test"))
	assert.False(t,matcher.MatchStart("tes?", "tes"))
	assert.False(t,matcher.MatchStart("tes?", "testt"))
	assert.False(t,matcher.MatchStart("tes?", "tsst"))

	// test matching with *'s
	assert.True(t,matcher.MatchStart("*", "test"))
	assert.True(t,matcher.MatchStart("test*", "test"))
	assert.True(t,matcher.MatchStart("test*", "testTest"))
	assert.True(t,matcher.MatchStart("test/*", "test/Test"))
	assert.True(t,matcher.MatchStart("test/*", "test/t"))
	assert.True(t,matcher.MatchStart("test/*", "test/"))
	assert.True(t,matcher.MatchStart("*test*", "AnothertestTest"))
	assert.True(t,matcher.MatchStart("*test", "Anothertest"))
	assert.True(t,matcher.MatchStart("*.*", "test."))
	assert.True(t,matcher.MatchStart("*.*", "test.test"))
	assert.True(t,matcher.MatchStart("*.*", "test.test.test"))
	assert.True(t,matcher.MatchStart("test*aaa", "testblaaaa"))
	assert.False(t,matcher.MatchStart("test*", "tst"))
	assert.False(t,matcher.MatchStart("test*", "test/"))
	assert.False(t,matcher.MatchStart("test*", "tsttest"))
	assert.False(t,matcher.MatchStart("test*", "test/"))
	assert.False(t,matcher.MatchStart("test*", "test/t"))
	assert.True(t,matcher.MatchStart("test/*", "test"))
	assert.True(t,matcher.MatchStart("test/t*.txt", "test"))
	assert.False(t,matcher.MatchStart("*test*", "tsttst"))
	assert.False(t,matcher.MatchStart("*test", "tsttst"))
	assert.False(t,matcher.MatchStart("*.*", "tsttst"))
	assert.False(t,matcher.MatchStart("test*aaa", "test"))
	assert.False(t,matcher.MatchStart("test*aaa", "testblaaab"))

	// test matching with ?'s and /'s
	assert.True(t,matcher.MatchStart("/?", "/a"))
	assert.True(t,matcher.MatchStart("/?/a", "/a/a"))
	assert.True(t,matcher.MatchStart("/a/?", "/a/b"))
	assert.True(t,matcher.MatchStart("/??/a", "/aa/a"))
	assert.True(t,matcher.MatchStart("/a/??", "/a/bb"))
	assert.True(t,matcher.MatchStart("/?", "/a"))

	// test matching with **'s
	assert.True(t,matcher.MatchStart("/**", "/testing/testing"))
	assert.True(t,matcher.MatchStart("/*/**", "/testing/testing"))
	assert.True(t,matcher.MatchStart("/**/*", "/testing/testing"))
	assert.True(t,matcher.MatchStart("test*/**", "test/"))
	assert.True(t,matcher.MatchStart("test*/**", "test/t"))
	assert.True(t,matcher.MatchStart("/bla/**/bla", "/bla/testing/testing/bla"))
	assert.True(t,matcher.MatchStart("/bla/**/bla", "/bla/testing/testing/bla/bla"))
	assert.True(t,matcher.MatchStart("/**/test", "/bla/bla/test"))
	assert.True(t,matcher.MatchStart("/bla/**/**/bla", "/bla/bla/bla/bla/bla/bla"))
	assert.True(t,matcher.MatchStart("/bla*bla/test", "/blaXXXbla/test"))
	assert.True(t,matcher.MatchStart("/*bla/test", "/XXXbla/test"))
	assert.False(t,matcher.MatchStart("/bla*bla/test", "/blaXXXbl/test"))
	assert.False(t,matcher.MatchStart("/*bla/test", "XXXblab/test"))
	assert.False(t,matcher.MatchStart("/*bla/test", "XXXbl/test"))

	assert.False(t,matcher.MatchStart("/????", "/bala/bla"))
	assert.True(t,matcher.MatchStart("/**/*bla", "/bla/bla/bla/bbb"))

	assert.True(t,matcher.MatchStart("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing/"))
	assert.True(t,matcher.MatchStart("/*bla*/**/bla/*", "/XXXblaXXXX/testing/testing/bla/testing"))
	assert.True(t,matcher.MatchStart("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing"))
	assert.True(t,matcher.MatchStart("/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing.jpg"))

	assert.True(t,matcher.MatchStart("*bla*/**/bla/**", "XXXblaXXXX/testing/testing/bla/testing/testing/"))
	assert.True(t,matcher.MatchStart("*bla*/**/bla/*", "XXXblaXXXX/testing/testing/bla/testing"))
	assert.True(t,matcher.MatchStart("*bla*/**/bla/**", "XXXblaXXXX/testing/testing/bla/testing/testing"))
	assert.True(t,matcher.MatchStart("*bla*/**/bla/*", "XXXblaXXXX/testing/testing/bla/testing/testing"))

	assert.True(t,matcher.MatchStart("/x/x/**/bla", "/x/x/x/"))

	assert.True(t,matcher.MatchStart("", ""))
}

//TestExtractPathWithinPattern
func TestExtractPathWithinPattern(t *testing.T){
	
	assert.Equal(t,"", matcher.ExtractPathWithinPattern("/docs/commit.html", "/docs/commit.html"))

	assert.Equal(t,"cvs/commit", matcher.ExtractPathWithinPattern("/docs/*", "/docs/cvs/commit"))
	assert.Equal(t,"commit.html", matcher.ExtractPathWithinPattern("/docs/cvs/*.html", "/docs/cvs/commit.html"))
	assert.Equal(t,"cvs/commit", matcher.ExtractPathWithinPattern("/docs/**", "/docs/cvs/commit"))
	assert.Equal(t,"cvs/commit.html", matcher.ExtractPathWithinPattern("/docs/**/*.html", "/docs/cvs/commit.html"))
	assert.Equal(t,"commit.html", matcher.ExtractPathWithinPattern("/docs/**/*.html", "/docs/commit.html"))
	assert.Equal(t,"commit.html", matcher.ExtractPathWithinPattern("/*.html", "/commit.html"))
	assert.Equal(t,"docs/commit.html", matcher.ExtractPathWithinPattern("/*.html", "/docs/commit.html"))
	assert.Equal(t,"/commit.html", matcher.ExtractPathWithinPattern("*.html", "/commit.html"))
	assert.Equal(t,"/docs/commit.html", matcher.ExtractPathWithinPattern("*.html", "/docs/commit.html"))
	assert.Equal(t,"/docs/commit.html", matcher.ExtractPathWithinPattern("**/*.*", "/docs/commit.html"))
	assert.Equal(t,"/docs/commit.html", matcher.ExtractPathWithinPattern("*", "/docs/commit.html"))
	// SPR-10515
	assert.Equal(t,"/docs/cvs/other/commit.html", matcher.ExtractPathWithinPattern("**/commit.html", "/docs/cvs/other/commit.html"))
	assert.Equal(t,"cvs/other/commit.html", matcher.ExtractPathWithinPattern("/docs/**/commit.html", "/docs/cvs/other/commit.html"))
	assert.Equal(t,"cvs/other/commit.html", matcher.ExtractPathWithinPattern("/docs/**/**/**/**", "/docs/cvs/other/commit.html"))

	assert.Equal(t,"docs/cvs/commit", matcher.ExtractPathWithinPattern("/d?cs/*", "/docs/cvs/commit"))
	assert.Equal(t,"cvs/commit.html", matcher.ExtractPathWithinPattern("/docs/c?s/*.html", "/docs/cvs/commit.html"))
	assert.Equal(t,"docs/cvs/commit", matcher.ExtractPathWithinPattern("/d?cs/**", "/docs/cvs/commit"))
	assert.Equal(t,"docs/cvs/commit.html",matcher.ExtractPathWithinPattern("/d?cs/**/*.html", "/docs/cvs/commit.html"))
}

//TestExtractUriTemplateVariables
func TestExtractUriTemplateVariables(t *testing.T)  {
	
	result := matcher.ExtractUriTemplateVariables("/hotels/{hotel}", "/hotels/1")
	assert.Equal(t,"1", (*result)["hotel"])

	result = matcher.ExtractUriTemplateVariables("/h?tels/{hotel}", "/hotels/1")
	assert.Equal(t,"1", (*result)["hotel"])

	result = matcher.ExtractUriTemplateVariables("/hotels/{hotel}/bookings/{booking}", "/hotels/1/bookings/2")
	assert.Equal(t,"1", (*result)["hotel"])
	assert.Equal(t,"2", (*result)["booking"])

	result = matcher.ExtractUriTemplateVariables("/**/hotels/**/{hotel}", "/foo/hotels/bar/1")
	assert.Equal(t,"1", (*result)["hotel"])

	result = matcher.ExtractUriTemplateVariables("/{page}.html", "/42.html")
	assert.Equal(t,"42", (*result)["page"])

	result = matcher.ExtractUriTemplateVariables("/{page}.*", "/42.html")
	assert.Equal(t,"42", (*result)["page"])

	result = matcher.ExtractUriTemplateVariables("/A-{B}-C", "/A-b-C")
	assert.Equal(t,"b", (*result)["B"])

	result = matcher.ExtractUriTemplateVariables("/{name}.{extension}", "/test.html")
	assert.Equal(t,"test", (*result)["name"])
	assert.Equal(t,"html", (*result)["extension"])
}

//TestExtractUriTemplateVariablesRegex
func TestExtractUriTemplateVariablesRegex(t *testing.T) {
	

	result := matcher.ExtractUriTemplateVariables("{symbolicName:[\\w\\.]+}-{version:[\\w\\.]+}.jar", "com.example-1.0.0.jar")
	assert.Equal(t,"com.example",(*result)["symbolicName"])
	assert.Equal(t,"1.0.0",(*result)["version"])

	result = matcher.ExtractUriTemplateVariables("{symbolicName:[\\w\\.]+}-sources-{version:[\\w\\.]+}.jar",
	"com.example-sources-1.0.0.jar")
	assert.Equal(t,"com.example", (*result)["symbolicName"])
	assert.Equal(t,"1.0.0", (*result)["version"])
}

/**
* SPR-7787
*/
//TestExtractUriTemplateVarsRegexQualifiers
func TestExtractUriTemplateVarsRegexQualifiers(t *testing.T) {
	

	result := matcher.ExtractUriTemplateVariables("{symbolicName:[\\p{L}\\.]+}-sources-{version:[\\p{N}\\.]+}.jar", "com.example-sources-1.0.0.jar")
	assert.Equal(t,"com.example", (*result)["symbolicName"])
	assert.Equal(t,"1.0.0", (*result)["version"])

	result = matcher.ExtractUriTemplateVariables(
	"{symbolicName:[\\w\\.]+}-sources-{version:[\\d\\.]+}-{year:\\d{4}}{month:\\d{2}}{day:\\d{2}}.jar",
	"com.example-sources-1.0.0-20100220.jar")
	assert.Equal(t,"com.example", (*result)["symbolicName"])
	assert.Equal(t,"1.0.0", (*result)["version"])
	assert.Equal(t,"2010", (*result)["year"])
	assert.Equal(t,"02", (*result)["month"])
	assert.Equal(t,"20", (*result)["day"])

	result = matcher.ExtractUriTemplateVariables(
	"{symbolicName:[\\p{L}\\.]+}-sources-{version:[\\p{N}\\.\\{\\}]+}.jar",
	"com.example-sources-1.0.0.{12}.jar")
	assert.Equal(t,"com.example", (*result)["symbolicName"])
	assert.Equal(t,"1.0.0.{12}", (*result)["version"])
}

/**
	 * SPR-8455
	 */
//TestExtractUriTemplateVarsRegexCapturingGroups
func TestExtractUriTemplateVarsRegexCapturingGroups(t *testing.T) {
	result := matcher.ExtractUriTemplateVariables("/web/{id:foo(bar)?}", "/web/foobar")
	assert.Equal(t,"foobar", (*result)["id"])
}

//TestGetPatternComparator
func TestGetPatternComparator(t *testing.T){
	

	comparator := matcher.GetPatternComparator("/hotels/new")

	assert.Equal(t,0, comparator.Compare("", ""))
	assert.Equal(t,1, comparator.Compare("", "/hotels/new"))
	assert.Equal(t,-1, comparator.Compare("/hotels/new", ""))

	assert.Equal(t,0, comparator.Compare("/hotels/new", "/hotels/new"))

	assert.Equal(t,-1, comparator.Compare("/hotels/new", "/hotels/*"))
	assert.Equal(t,1, comparator.Compare("/hotels/*", "/hotels/new"))
	assert.Equal(t,0, comparator.Compare("/hotels/*", "/hotels/*"))

	assert.Equal(t,-1, comparator.Compare("/hotels/new", "/hotels/{hotel}"))
	assert.Equal(t,1, comparator.Compare("/hotels/{hotel}", "/hotels/new"))
	assert.Equal(t,0, comparator.Compare("/hotels/{hotel}", "/hotels/{hotel}"))
	assert.Equal(t,-1, comparator.Compare("/hotels/{hotel}/booking", "/hotels/{hotel}/bookings/{booking}"))
	assert.Equal(t,1, comparator.Compare("/hotels/{hotel}/bookings/{booking}", "/hotels/{hotel}/booking"))

	// SPR-10550
	assert.Equal(t,-1, comparator.Compare("/hotels/{hotel}/bookings/{booking}/cutomers/{customer}", "/**"))
	assert.Equal(t,1, comparator.Compare("/**", "/hotels/{hotel}/bookings/{booking}/cutomers/{customer}"))
	assert.Equal(t,0, comparator.Compare("/**", "/**"))

	assert.Equal(t,-1, comparator.Compare("/hotels/{hotel}", "/hotels/*"))
	assert.Equal(t,1, comparator.Compare("/hotels/*", "/hotels/{hotel}"))

	assert.Equal(t,-1, comparator.Compare("/hotels/*", "/hotels/*/**"))
	assert.Equal(t,1, comparator.Compare("/hotels/*/**", "/hotels/*"))

	assert.Equal(t,-1, comparator.Compare("/hotels/new", "/hotels/new.*"))
	assert.Equal(t,2, comparator.Compare("/hotels/{hotel}", "/hotels/{hotel}.*"))

	// SPR-6741
	assert.Equal(t,-1, comparator.Compare("/hotels/{hotel}/bookings/{booking}/cutomers/{customer}", "/hotels/**"))
	assert.Equal(t,1, comparator.Compare("/hotels/**", "/hotels/{hotel}/bookings/{booking}/cutomers/{customer}"))
	assert.Equal(t,1, comparator.Compare("/hotels/foo/bar/**", "/hotels/{hotel}"))
	assert.Equal(t,-1, comparator.Compare("/hotels/{hotel}", "/hotels/foo/bar/**"))
	assert.Equal(t,2, comparator.Compare("/hotels/**/bookings/**", "/hotels/**"))
	assert.Equal(t,-2, comparator.Compare("/hotels/**", "/hotels/**/bookings/**"))

	// SPR-8683
	assert.Equal(t,1, comparator.Compare("/**", "/hotels/{hotel}"))

	// longer is better
	assert.Equal(t,1, comparator.Compare("/hotels", "/hotels2"))

	// SPR-13139
	assert.Equal(t,-1, comparator.Compare("*", "*/**"))
	assert.Equal(t,1, comparator.Compare("*/**", "*"))
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

// SPR-14247
//TestMatchWithTrimTokensEnabled
func TestMatchWithTrimTokensEnabled(t *testing.T){
	matcher.SetTrimTokens(true)
	assert.True(t,matcher.Match("/foo/bar", "/foo /bar"))
}

//TestUniqueDeliminator
func TestUniqueDeliminator(t *testing.T){
	matcher.SetPathSeparator(".")

	// test exact matching
	assert.True(t,matcher.Match("test", "test"))
	assert.True(t,matcher.Match(".test", ".test"))
	assert.False(t,matcher.Match(".test/jpg", "test/jpg"))
	assert.False(t,matcher.Match("test", ".test"))
	assert.False(t,matcher.Match(".test", "test"))

	// test matching with ?'s
	assert.True(t,matcher.Match("t?st", "test"))
	assert.True(t,matcher.Match("??st", "test"))
	assert.True(t,matcher.Match("tes?", "test"))
	assert.True(t,matcher.Match("te??", "test"))
	assert.True(t,matcher.Match("?es?", "test"))
	assert.False(t,matcher.Match("tes?", "tes"))
	assert.False(t,matcher.Match("tes?", "testt"))
	assert.False(t,matcher.Match("tes?", "tsst"))

	// test matching with *'s
	assert.True(t,matcher.Match("*", "test"))
	assert.True(t,matcher.Match("test*", "test"))
	assert.True(t,matcher.Match("test*", "testTest"))
	assert.True(t,matcher.Match("*test*", "AnothertestTest"))
	assert.True(t,matcher.Match("*test", "Anothertest"))
	assert.True(t,matcher.Match("*/*", "test/"))
	assert.True(t,matcher.Match("*/*", "test/test"))
	assert.True(t,matcher.Match("*/*", "test/test/test"))
	assert.True(t,matcher.Match("test*aaa", "testblaaaa"))
	assert.False(t,matcher.Match("test*", "tst"))
	assert.False(t,matcher.Match("test*", "tsttest"))
	assert.False(t,matcher.Match("*test*", "tsttst"))
	assert.False(t,matcher.Match("*test", "tsttst"))
	assert.False(t,matcher.Match("*/*", "tsttst"))
	assert.False(t,matcher.Match("test*aaa", "test"))
	assert.False(t,matcher.Match("test*aaa", "testblaaab"))

	// test matching with ?'s and .'s
	assert.True(t,matcher.Match(".?", ".a"))
	assert.True(t,matcher.Match(".?.a", ".a.a"))
	assert.True(t,matcher.Match(".a.?", ".a.b"))
	assert.True(t,matcher.Match(".??.a", ".aa.a"))
	assert.True(t,matcher.Match(".a.??", ".a.bb"))
	assert.True(t,matcher.Match(".?", ".a"))

	// test matching with **'s
	assert.True(t,matcher.Match(".**", ".testing.testing"))
	assert.True(t,matcher.Match(".*.**", ".testing.testing"))
	assert.True(t,matcher.Match(".**.*", ".testing.testing"))
	assert.True(t,matcher.Match(".bla.**.bla", ".bla.testing.testing.bla"))
	assert.True(t,matcher.Match(".bla.**.bla", ".bla.testing.testing.bla.bla"))
	assert.True(t,matcher.Match(".**.test", ".bla.bla.test"))
	assert.True(t,matcher.Match(".bla.**.**.bla", ".bla.bla.bla.bla.bla.bla"))
	assert.True(t,matcher.Match(".bla*bla.test", ".blaXXXbla.test"))
	assert.True(t,matcher.Match(".*bla.test", ".XXXbla.test"))
	assert.False(t,matcher.Match(".bla*bla.test", ".blaXXXbl.test"))
	assert.False(t,matcher.Match(".*bla.test", "XXXblab.test"))
	assert.False(t,matcher.Match(".*bla.test", "XXXbl.test"))
}

// SPR-8687
//TestTrimTokensOff
func TestTrimTokensOff(t *testing.T) {
	matcher.SetTrimTokens(false)
	assert.True(t,matcher.Match("/group/{groupName}/members", "/group/sales/members"))
	assert.True(t,matcher.Match("/group/{groupName}/members", "/group/  sales/members"))
	assert.False(t,matcher.Match("/group/{groupName}/members", "/Group/  Sales/Members"))
}

//TestDefaultCacheSetting
func TestDefaultCacheSetting(t *testing.T) {
	t1 := time.Now().Nanosecond()
	for i := 0; i < 65536; i++ {
		matcher.Match(fmt.Sprint("test",i), fmt.Sprint("test",i))
	}
	// Cache turned off because it went beyond the threshold
	assert.True(t,matcher.PatternCacheSize()<= 0)
	t.Log(time.Now().Nanosecond() - t1)
}

//TestCachePatternsSetToFalse
func TestCachePatternsSetToFalse(t *testing.T) {
	matcherFalse := New()

	matcherFalse.SetCachePatterns(false)
	for i := 0; i < 10; i++ {
		matcherFalse.Match(fmt.Sprint("test",i), fmt.Sprint("test",i))
	}
	t.Log(matcherFalse.PatternCacheSize())
	assert.True(t,matcherFalse.PatternCacheSize()<= 0)
}

// SPR-13286
//TestCaseInsensitive
func TestCaseInsensitive(t *testing.T) {
	
	matcher.SetCaseSensitive(false)

	assert.True(t,matcher.Match("/group/{groupName}/members", "/group/sales/members"))
	assert.True(t,matcher.Match("/group/{groupName}/members", "/Group/Sales/Members"))
	assert.True(t,matcher.Match("/Group/{groupName}/Members", "/group/Sales/members"))
}
