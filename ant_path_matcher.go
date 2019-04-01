/**
 * Created by GoLand.
 * Brief: apache ant path matcher implement
 * User: vibrant
 * Date: 2019/03/30
 * Time: 13:10
 */
package antpath

import (
	"strings"
	"sync"
)

//AntPathMatcher
type AntPathMatcher struct {
	//path separator
	pathSeparator string
	//tokenize pattern cache(thread safe)
	tokenizedPatternCache *sync.Map
	//stringMatcherCache string-matcher cache (thread safe)
	stringMatcherCache *sync.Map
	//pathSeparatorPatternCache
	pathSeparatorPatternCache *PathSeparatorPatternCache

	//-------private---------//
	// 忽略大小写（不区分大小写）
	//caseSensitive default value = true
	caseSensitive bool
	//trimTokens default value = false
	trimTokens bool
	//cachePatterns default value = true
	cachePatterns bool
}

/**
* Create a new instance with the {@link #DEFAULT_PATH_SEPARATOR}.
*/
//New
func New() *AntPathMatcher {
	ant := NewS(DefaultPathSeparator)
	return ant
}

/**
* Create a new instance with the {@link #DEFAULT_PATH_SEPARATOR}.
*/
//NewS
func NewS(separator string) *AntPathMatcher{
	if strings.EqualFold(EmptyString,separator) {
		separator = DefaultPathSeparator
	}
	ant := &AntPathMatcher{}
	//
	ant.pathSeparator = separator
	ant.tokenizedPatternCache = new(sync.Map)
	ant.stringMatcherCache = new(sync.Map)
	ant.pathSeparatorPatternCache = NewDefaultPathSeparatorPatternCache(separator)

	//filed
	ant.caseSensitive = true
	ant.trimTokens = false
	ant.cachePatterns = true
	return ant
}

//@Override
//IsPattern
func (ant *AntPathMatcher) IsPattern(path string) bool{
	return strings.Index(path,"*")!=-1 || strings.Index(path,"?")!=-1
}

//@Override
//Match
func (ant *AntPathMatcher) Match(pattern,path string) bool{
	return ant.doMatch(pattern, path, true, nil)
}

//@Override
//MatchStart
func (ant *AntPathMatcher) MatchStart(pattern,path string) bool{
	return ant.doMatch(pattern, path, false, nil)
}

//@Override
//ExtractPathWithinPattern
func (ant *AntPathMatcher) ExtractPathWithinPattern(pattern,path string) string{
	patternParts := TokenizeToStringArray(pattern, ant.pathSeparator, ant.trimTokens, true)
	pathParts := TokenizeToStringArray(path, ant.pathSeparator, ant.trimTokens, true)
	builder := EmptyString
	pathStarted := false
	for segment := 0; segment < len(patternParts); segment++ {
		patternPart := patternParts[segment]
		if strings.Index(*patternPart,"*") > -1 || strings.Index(*patternPart,"?") > -1 {
			for ;segment < len(pathParts); segment++ {
				if pathStarted || (segment == 0 && !strings.HasSuffix(pattern,ant.pathSeparator)) {
					builder += ant.pathSeparator
				}
				builder += *pathParts[segment]
				pathStarted = true
			}
		}
	}

	return builder
}