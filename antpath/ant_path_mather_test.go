package antpath

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
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

func TestRune(t *testing.T){
	t.Log(utf8.RuneCountInString("Hello, 世界"))

	t.Log(utf8.RuneCountInString("*"))

	t.Log(utf8.RuneCountInString("?"))

	t.Log(utf8.DecodeLastRuneInString("{"))

	for _,char := range WildcardChars {

		r,_ :=utf8.DecodeLastRuneInString("{")

		if char == r {
			t.Log("test name")
		}
		t.Log(char)
	}

	b := make([]byte, utf8.UTFMax)

	n := utf8.EncodeRune(b, '*')
	fmt.Printf("%v：%v\n", b, n) // [229 165 189 0]：3

	r, n := utf8.DecodeRune(b)
	fmt.Printf("%c：%v\n", r, n) // 好：3

	t.Log("----------")
	t.Log(reflect.TypeOf("22"[0]))

	t.Log("----------1")
	t.Log(string("\""[0]))


	t.Log(rune("**"[0]))



	t.Log("字符串长度测试")
	str1 := "Hello, 世界"
	t.Log(len(str1))
	t.Log(utf8.RuneCountInString(str1))


}

//TestSkipSeparator
func TestSkipSeparator(t *testing.T){
	t.Log(strings.HasPrefix("/vv/mm/dd/ii","/"))
}

//TestStartsWith
func TestStartsWith(t *testing.T)  {
	t.Log(StartsWith("/vv/mm/dd/ii","/",0))
	t.Log(StartsWith("/vv/mm/dd/ii","/",3))
	t.Log(StartsWith("/vv/mm/dd/ii","/",6))
	t.Log(StartsWith("/vv/mm/dd/ii","/",9))
	t.Log(StartsWith("/vv/mm/dd/ii","/",5))
}