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

//AntPathMatcher implement from PathMatcher interface
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
				if pathStarted || (segment == 0 && !strings.HasPrefix(pattern,ant.pathSeparator)) {
					builder += ant.pathSeparator
				}
				builder += *pathParts[segment]
				pathStarted = true
			}
		}
	}

	return builder
}

//@Override
//ExtractUriTemplateVariables
func (ant *AntPathMatcher) ExtractUriTemplateVariables(pattern,path string) *map[string]string{
	variables := make(map[string]string)
	result := ant.doMatch(pattern, path, true, &variables)
	if !result {
		panic("Pattern \"" + pattern + "\" is not a match for \"" + path + "\"")
	}
	return &variables
}

//@Override
//GetPatternComparator
func (ant *AntPathMatcher) GetPatternComparator(path string) *AntPatternComparator {
	return NewDefaultAntPatternComparator(path)
}

//@Override
//Combine 将pattern1和pattern2联合成一个新的pattern
func (ant *AntPathMatcher) Combine(pattern1,pattern2 string) string  {
	if !HasText(pattern1) && !HasText(pattern2) {
		return ""
	}
	if !HasText(pattern1) {
		return pattern2
	}
	if !HasText(pattern2) {
		return pattern1
	}
	//处理pattern
	pattern1ContainsUriVar := strings.Index(pattern1,"{") != -1
	if !strings.EqualFold(pattern1,pattern2) && !pattern1ContainsUriVar && ant.Match(pattern1, pattern2) {
		// /* + /hotel -> /hotel ; "/*.*" + "/*.html" -> /*.html
		// However /user + /user -> /usr/user ; /{foo} + /bar -> /{foo}/bar
		return pattern2
	}
	// /hotels/* + /booking -> /hotels/booking
	// /hotels/* + booking -> /hotels/booking
	if strings.HasSuffix(pattern1,ant.pathSeparatorPatternCache.GetEndsOnWildCard()) {
		return ant.concat(pattern1[0:len(pattern1)-2], pattern2)
	}

	// /hotels/** + /booking -> /hotels/**/booking
	// /hotels/** + booking -> /hotels/**/booking
	if strings.HasSuffix(pattern1,ant.pathSeparatorPatternCache.GetEndsOnDoubleWildCard()) {
		return ant.concat(pattern1, pattern2)
	}

	starDotPos1 := strings.Index(pattern1,"*.")
	if pattern1ContainsUriVar || starDotPos1 == -1 || strings.EqualFold(".",ant.pathSeparator) {
		// simply concatenate the two patterns
		return ant.concat(pattern1, pattern2)
	}

	ext1 := pattern1[starDotPos1+1:]
	dotPos2 := strings.Index(pattern2,".")
	file2 := EmptyString
	ext2 := EmptyString
	if dotPos2 == -1 {
		file2 = pattern2
		ext2 = ""
	}else{
		file2 = pattern2[0:dotPos2]
		ext2 = pattern2[dotPos2:]
	}
	ext1All := strings.EqualFold(".*",ext1) || strings.EqualFold(EmptyString,ext1)
	ext2All := strings.EqualFold (".*",ext2)|| strings.EqualFold(EmptyString,ext2)
	if !ext1All && !ext2All {
		panic("Cannot combine patterns: " + pattern1 + " vs " + pattern2)
	}
	//
	ext := EmptyString
	if ext1All {
		ext	= ext2
	}else{
		ext = ext1
	}
	return file2 + ext
}