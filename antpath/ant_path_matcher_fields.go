/**
 * Created by GoLand.
 * Brief: matcher fields read/write
 * User: vibrant
 * Date: 2019/04/01
 * Time: 14:57
 */
package antpath

import "strings"

/**
Set the path separator to use for pattern parsing.
 */
//SetPathSeparator The default is "/",as in ant.
func (ant *AntPathMatcher) SetPathSeparator(pathSeparator string)  {
	if !strings.EqualFold(EmptyString,pathSeparator) {
		ant.pathSeparator = pathSeparator
		ant.pathSeparatorPatternCache = NewDefaultPathSeparatorPatternCache(pathSeparator)
	}
}

/**
Specify whether to perform pattern matching in a case-sensitive fashion.
<p>Default is {@code true}. Switch this to {@code false} for case-insensitive matching.
 */
//SetCaseSensitive 是否忽略大小写 The default is false
func (ant *AntPathMatcher) SetCaseSensitive(caseSensitive bool){
	ant.caseSensitive = caseSensitive
}

/**
Specify whether to trim tokenized paths and patterns.
 */
//SetTrimTokens 是否去除空格 The default is false
func (ant *AntPathMatcher) SetTrimTokens(trimTokens bool){
	ant.trimTokens = trimTokens
}

/**
* Specify whether to cache parsed pattern metadata for patterns passed
* into this matcher's {@link #match} method. A value of {@code true}
* activates an unlimited pattern cache; a value of {@code false} turns
* the pattern cache off completely.
* <p>Default is for the cache to be on, but with the variant to automatically
* turn it off when encountering too many patterns to cache at runtime
* (the threshold is 65536), assuming that arbitrary permutations of patterns
* are coming in, with little chance for encountering a recurring pattern.
*/
//SetCachePatterns
func (ant *AntPathMatcher) SetCachePatterns(cachePatterns bool){
	ant.cachePatterns = cachePatterns
}