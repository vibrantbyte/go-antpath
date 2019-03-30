/**
 * Created by GoLand.
 * Brief: string tool
 * User: vibrant
 * Date: 2019/03/30
 * Time: 13:10
 */
package antpath

import "strings"

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

//TokenizeToStringArray
func TokenizeToStringArray1(str,delimiters string) []*string {
	return TokenizeToStringArray(str,delimiters,true,true)
}
