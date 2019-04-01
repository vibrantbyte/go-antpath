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
func (ant *AntPathMatcher) GetPatternComparator(path string) string {
	return ""
}

//@Override
//Combine 将pattern1和pattern2联合成一个新的pattern
func (ant *AntPathMatcher) Combine(pattern1,pattern2 string) string  {
	//if (!StringUtils.hasText(pattern1) && !StringUtils.hasText(pattern2)) {
	//	return "";
	//}
	//if (!StringUtils.hasText(pattern1)) {
	//	return pattern2;
	//}
	//if (!StringUtils.hasText(pattern2)) {
	//	return pattern1;
	//}
	//
	//boolean pattern1ContainsUriVar = (pattern1.indexOf('{') != -1);
	//if (!pattern1.equals(pattern2) && !pattern1ContainsUriVar && match(pattern1, pattern2)) {
	//	// /* + /hotel -> /hotel ; "/*.*" + "/*.html" -> /*.html
	//	// However /user + /user -> /usr/user ; /{foo} + /bar -> /{foo}/bar
	//	return pattern2;
	//}
	//
	//// /hotels/* + /booking -> /hotels/booking
	//// /hotels/* + booking -> /hotels/booking
	//if (pattern1.endsWith(this.pathSeparatorPatternCache.getEndsOnWildCard())) {
	//	return concat(pattern1.substring(0, pattern1.length() - 2), pattern2);
	//}
	//
	//// /hotels/** + /booking -> /hotels/**/booking
	//// /hotels/** + booking -> /hotels/**/booking
	//if (pattern1.endsWith(this.pathSeparatorPatternCache.getEndsOnDoubleWildCard())) {
	//	return concat(pattern1, pattern2);
	//}
	//
	//int starDotPos1 = pattern1.indexOf("*.");
	//if (pattern1ContainsUriVar || starDotPos1 == -1 || this.pathSeparator.equals(".")) {
	//	// simply concatenate the two patterns
	//	return concat(pattern1, pattern2);
	//}
	//
	//String ext1 = pattern1.substring(starDotPos1 + 1);
	//int dotPos2 = pattern2.indexOf('.');
	//String file2 = (dotPos2 == -1 ? pattern2 : pattern2.substring(0, dotPos2));
	//String ext2 = (dotPos2 == -1 ? "" : pattern2.substring(dotPos2));
	//boolean ext1All = (ext1.equals(".*") || ext1.isEmpty());
	//boolean ext2All = (ext2.equals(".*") || ext2.isEmpty());
	//if (!ext1All && !ext2All) {
	//	throw new IllegalArgumentException("Cannot combine patterns: " + pattern1 + " vs " + pattern2);
	//}
	//String ext = (ext1All ? ext2 : ext1);
	//return file2 + ext;
	return ""
}