/**
 * Created by GoLand.
 * Brief: string tool
 * User: vibrant
 * Date: 2019/03/30
 * Time: 13:10
 */
package antpath

import (
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

//EmptySpace
const EmptySpace = " "
//EmptyString
const EmptyString = ""

/**
* Tokenize the given {@code String} into a {@code String} array via a
* {@link StringTokenizer}.
* <p>The given {@code delimiters} string can consist of any number of
* delimiter characters. Each of those characters can be used to separate
* tokens. A delimiter is always a single character; for multi-character
* delimiters, consider using {@link #delimitedListToStringArray}.
* @param str the {@code String} to tokenize
* @param delimiters the delimiter characters, assembled as a {@code String}
* (each of the characters is individually considered as a delimiter)
* @param trimTokens trim the tokens via {@link String#trim()}
* @param ignoreEmptyTokens omit empty tokens from the result array
* (only applies to tokens that are empty after trimming; StringTokenizer
* will not consider subsequent delimiters as token in the first place).
* @return an array of the tokens ({@code null} if the input {@code String}
* was {@code null})
* @see java.util.StringTokenizer
* @see String#trim()
* @see #delimitedListToStringArray
*/
//TokenizeToStringArray
func TokenizeToStringArray(str,delimiters string,trimTokens,ignoreEmptyTokens bool) []*string {
	if str == EmptyString {
		return nil
	}
	tokens := make([]*string,0)
	for _,token := range strings.Split(str,delimiters)  {
		if trimTokens {
			token = strings.Trim(token,EmptySpace)
		}
		if !ignoreEmptyTokens || token != EmptyString {
			var item = token
			tokens = append(tokens,&item)
		}
	}
	return tokens
}

//TokenizeToStringArray1
func TokenizeToStringArray1(str,delimiters string) []*string {
	return TokenizeToStringArray(str,delimiters,true,true)
}

//Str2Bytes
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//Bytes2Str
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//StartsWith
func StartsWith(str,prefix string,toffset int) bool{
	ta := Str2Bytes(str)
	to := toffset
	pa := Str2Bytes(prefix)
	po := 0
	pc := utf8.RuneCountInString(prefix)
	// Note: toffset might be near -1>>>1.
	if (toffset < 0) || (toffset > utf8.RuneCountInString(str) - pc) {
		return false
	}
	for  {
		if pc--;pc >= 0 {
			if ta[to] != pa[po] {
				to++
				po++
				return false
			}
		}else{
			break
		}
	}
	return true
}

//IsBlank 判断是否存在空格
func IsBlank(source string) bool{
	if strings.EqualFold(EmptyString,source) {
		return true
	}
	for i := len(source); i > 0; {
		r, size := utf8.DecodeLastRuneInString(source[0:i])
		i -= size
		if !unicode.IsSpace(r){
			return false
		}
	}
	return true
}