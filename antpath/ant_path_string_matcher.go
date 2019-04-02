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

	//caseSensitive 忽略大小写（不区分大小写）
	caseSensitive bool
}

//NewDefaultStringMatcher part match
func NewDefaultStringMatcher(pattern string,caseSensitive bool)*AntPathStringMatcher{
	stringMatcher := &AntPathStringMatcher{}
	//caseSensitive
	stringMatcher.caseSensitive = caseSensitive
	//写入表达式
	reg,err := regexp.Compile(*stringMatcher.patternBuilder(pattern,false,caseSensitive))
	if err == nil {
		stringMatcher.pattern = reg
	}
	return stringMatcher
}

//NewMatchesStringMatcher full match
func NewMatchesStringMatcher(pattern string,caseSensitive bool) *AntPathStringMatcher  {
	stringMatcher := &AntPathStringMatcher{}
	//caseSensitive
	stringMatcher.caseSensitive = caseSensitive
	//写入表达式
	reg,err := regexp.Compile(*stringMatcher.patternBuilder(pattern,true,caseSensitive))
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
//MatchStrings
func (sm *AntPathStringMatcher) MatchStrings(str string,uriTemplateVariables *map[string]string) bool {
	//忽略大小写（不区分大小写）
	if sm.caseSensitive {
		str = strings.ToLower(str)
	}
	//byte
	matchBytes := Str2Bytes(str)
	findIndex := sm.pattern.FindAllIndex(matchBytes,MaxFindCount)
	if len(findIndex) > 0 {
		if uriTemplateVariables != nil{
			// SPR-8455
			if findIndex == nil && len(*uriTemplateVariables) != len(findIndex) {
				panic("The number of capturing groups in the pattern segment " +
					sm.pattern.String() + " does not match the number of URI template variables it defines, " +
					"which can occur if capturing groups are used in a URI template regex. " +
					"Use non-capturing groups instead.")
			}
			for i := 1; i <= len(findIndex); i++ {
				name := sm.variableNames[i - 1]
				//获取匹配位置
				matched := findIndex[i]
				matchedStart := matched[0]
				matchedEnd := matched[1]
				value := Bytes2Str(matchBytes[matchedStart:matchedEnd])
				(*uriTemplateVariables)[*name] = value
			}
		}
		return true
	}else{
		return false
	}
}

//quote
func (sm *AntPathStringMatcher) quote(s string,start,end int) string {
	if start == end {
		return ""
	}
	return regexp.QuoteMeta(s[start:end])
}

//patternBuilder
func (sm *AntPathStringMatcher) patternBuilder(pattern string,matches,caseSensitive bool) *string {
	//字符串拼接
	var patternBuilder string
	end := 0
	patternBytes := Str2Bytes(pattern)
	allIndex := GlobPattern.FindAllIndex(patternBytes,MaxFindCount)
	if allIndex != nil && len(allIndex) > 0 {
		for _,matched := range allIndex{
			matchedStart := matched[0]
			matchedEnd := matched[1]
			patternBuilder += sm.quote(pattern,end,matchedStart)
			//matchString
			matchstr := Bytes2Str(patternBytes[matchedStart:matchedEnd])
			if strings.EqualFold("?",matchstr){
				patternBuilder += "."
			}else if strings.EqualFold("*",matchstr){
				patternBuilder += ".*"
			}else if strings.HasPrefix(matchstr,"{") && strings.HasSuffix(matchstr,"}"){
				colonIdx := strings.Index(matchstr,":")
				if colonIdx == -1{
					patternBuilder += DefaultVariablePattern
					sm.variableNames = append(sm.variableNames,&matchstr)
				}else {
					bytes := Str2Bytes(matchstr)
					variablePattern := Bytes2Str(bytes[colonIdx + 1:len(matchstr)-1])
					patternBuilder += "("
					patternBuilder += variablePattern
					patternBuilder += ")"
					variableName :=Bytes2Str(bytes[1:colonIdx])
					sm.variableNames = append(sm.variableNames,&variableName)
				}
			}
			//向后增加end
			end = matchedEnd
		}
	}
	//patternBuilder
	patternBuilder += sm.quote(pattern,end,len(pattern))
	if caseSensitive {
		patternBuilder = strings.ToLower(patternBuilder)
	}
	//full match
	if matches {
		patternBuilder = "^" + patternBuilder + "$"
	}

	return &patternBuilder
}
