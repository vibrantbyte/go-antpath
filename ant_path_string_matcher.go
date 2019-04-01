/**
 * Created by GoLand.
 * Brief: apache ant path matcher implement
 * User: vibrant
 * Date: 2019/03/31
 * Time: 13:51
 */
package antpath

import (
	"regexp"
	"strings"
)

/**
// QuoteMeta 将字符串 s 中的“特殊字符”转换为其“转义格式”
// 例如，QuoteMeta（`[foo]`）返回`\[foo\]`。
// 特殊字符有：\.+*?()|[]{}^$
// 这些字符用于实现正则语法，所以当作普通字符使用时需要转换
func QuoteMeta
// 通过 Complite、CompilePOSIX、MustCompile、MustCompilePOSIX
// 四个函数可以创建一个 Regexp 对象
struct Regexp
// Compile 用来解析正则表达式 expr 是否合法，如果合法，则返回一个 Regexp 对象
// Regexp 对象可以在任意文本上执行需要的操作
func Compile

// 在 s 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
// 同时返回子表达式匹配的内容
// {完整匹配项, 子匹配项, 子匹配项, ...}
func FindStringSubmatch
// 在 b 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
// 同时返回子表达式匹配的内容
// {{完整匹配项}, {子匹配项}, {子匹配项}, ...}
func FindSubmatch
 */

//DefaultVariablePattern
const DefaultVariablePattern = "(.*)"
//MaxFindCount default value = 32
const MaxFindCount = 1<<5

//GlobPattern
var GlobPattern *regexp.Regexp

//initial
func init() {
	//reg
	reg,err := regexp.Compile("\\?|\\*|\\{((?:\\{[^/]+?\\}|[^/{}]|\\\\[{}])+?)\\}")
	if err == nil {
		GlobPattern = reg
	}

}

/**
* Tests whether or not a string matches against a pattern via a {@link Pattern}.
* <p>The pattern may contain special characters: '*' means zero or more characters; '?' means one and
* only one character; '{' and '}' indicate a URI template pattern. For example <tt>/users/{user}</tt>.
*/
//AntPathStringMatcher
type AntPathStringMatcher struct {
	//variableNames
	variableNames []*string
	//pattern
	pattern *regexp.Regexp
}

//NewStringMatcher-目前是regexp.Compile 不支持多参数（FoldCase  Flags = 1 << iota // case-insensitive match）
func NewStringMatcher(pattern string) *AntPathStringMatcher  {
	stringMatcher := &AntPathStringMatcher{}
	//字符串拼接
	var patternBuilder string
	end := 0
	matchedList := GlobPattern.FindAllString(pattern,MaxFindCount)
	if matchedList != nil && len(matchedList) > 0 {
		for index,matched := range matchedList{
			patternBuilder += stringMatcher.quote(pattern,end,index)
			if strings.EqualFold("?",matched){
				patternBuilder += "."
			}else if strings.EqualFold("*",matched){
				patternBuilder += ".*"
			}else if strings.HasPrefix(matched,"{") && strings.HasSuffix(matched,"}"){
				colonIdx := strings.Index(matched,":")
				if colonIdx == -1{
					patternBuilder += DefaultVariablePattern
				}else {
					bytes := Str2Bytes(matched)
					variablePattern := Bytes2Str(bytes[colonIdx + 1:len(matched)-1])
					patternBuilder += "("
					patternBuilder += variablePattern
					patternBuilder += ")"
					variableName :=Bytes2Str(bytes[1:colonIdx])
					stringMatcher.variableNames = append(stringMatcher.variableNames,&variableName)
				}
			}
			//向后增加end

		}
	}
	//patternBuilder
	patternBuilder += stringMatcher.quote(pattern,end,len(pattern))
	reg,err := regexp.Compile(patternBuilder)
	if err == nil {
		stringMatcher.pattern = reg
	}
	return stringMatcher
}

/**
* Main entry point.
*
* @return {@code true} if the string matches against the pattern, or {@code false} otherwise.
*/
//matchStrings
func (sm *AntPathStringMatcher) matchStrings(str string,uriTemplateVariables *map[string]string) bool {
	//可移植性操作系统接口
	//matcher := regexp.MustCompile(str)
	if uriTemplateVariables != nil{
		// SPR-8455
		//matcher.
		//
		//if len(sm.variableNames) !=  matcher .groupCount()) {
		//	throw new IllegalArgumentException("The number of capturing groups in the pattern segment " +
		//		this.pattern + " does not match the number of URI template variables it defines, " +
		//		"which can occur if capturing groups are used in a URI template regex. " +
		//		"Use non-capturing groups instead.");
		//}
		//for (int i = 1; i <= matcher.groupCount(); i++) {
		//	String name = this.variableNames.get(i - 1);
		//	String value = matcher.group(i);
		//	uriTemplateVariables.put(name, value);
		//}
	}
	return true
}

//quote
func (sm *AntPathStringMatcher) quote(s string,start,end int) string {
	if start == end {
		return ""
	}
	return regexp.QuoteMeta(s[start:end])
}
