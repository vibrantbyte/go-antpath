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
	"unicode/utf8"
)

/**
* Actually match the given {@code path} against the given {@code pattern}.
* @param pattern the pattern to match against
* @param path the path String to test
* @param fullMatch whether a full pattern match is required (else a pattern match
* as far as the given base path goes is sufficient)
* @return {@code true} if the supplied {@code path} matched, {@code false} if it didn't
*/
//doMatch
func (ant *AntPathMatcher) doMatch(pattern,path string,fullMatch bool,uriTemplateVariables *map[string]string) bool{
	if strings.HasPrefix(path,DefaultPathSeparator) != strings.HasPrefix(pattern,DefaultPathSeparator) {
		return false
	}
	pattDirs := ant.tokenizePattern(pattern)
	if fullMatch && ant.caseSensitive && !ant.isPotentialMatch(path,pattDirs){
		return false
	}

	pathDirs := ant.tokenizePath(path)
	//define variable
	pattIdxStart := 0
	pattIdxEnd := len(pattDirs) - 1
	pathIdxStart := 0
	pathIdxEnd := len(pathDirs) - 1

	// Match all elements up to the first **
	for{
		if pattIdxStart <= pattIdxEnd && pathIdxStart <= pathIdxEnd {
			pattDir := pattDirs[pattIdxStart]
			if strings.EqualFold("**", *pattDir) {
				break
			}
			if ant.matchStrings(*pattDir,*pathDirs[pathIdxStart],uriTemplateVariables) {
				return false
			}
		}else{
			//jump out of
			break
		}
	}

	if pathIdxStart > pathIdxEnd{
		// Path is exhausted, only match if rest of pattern is * or **'s
		if pattIdxStart > pattIdxEnd {
			if strings.HasSuffix(pattern, ant.pathSeparator){
				return strings.HasSuffix(path, ant.pathSeparator)
			}else {
				return !strings.HasSuffix(path, ant.pathSeparator)
			}
		}
		if !fullMatch {
			return true
		}
		if pattIdxStart == pattIdxEnd && strings.EqualFold("*",*pattDirs[pattIdxStart]) && strings.HasSuffix(path, ant.pathSeparator) {
			return true
		}
		for i:=pattIdxStart;i<=pattIdxEnd;i++  {
			if !strings.EqualFold("**",*pattDirs[i]) {
				return false
			}
		}
		return true
	}else if pattIdxStart > pattIdxEnd {
		// String not exhausted, but pattern is. Failure.
		return false
	}else if !fullMatch && strings.EqualFold("**",*pattDirs[pattIdxStart]){
		// Path start definitely matches due to "**" part in pattern.
		return true
	}

	// up to last '**'
	for {
		if pattIdxStart <= pattIdxEnd && pathIdxStart <= pathIdxEnd {
			pattDir := pattDirs[pattIdxEnd]
			if (strings.EqualFold("**",*pattDir)) {
				break
			}
			if (!ant.matchStrings(*pattDir, *pathDirs[pathIdxEnd], uriTemplateVariables)) {
				return false
			}
			pattIdxEnd--
			pathIdxEnd--
		}else {
			break
		}
	}
	if pathIdxStart > pathIdxEnd {
		// String is exhausted
		for i:= pattIdxStart;i<=pattIdxEnd;i++{
			if !strings.EqualFold("**",*pattDirs[i]) {
				return false
			}
		}
		return true
	}

	for {
		if pattIdxStart != pattIdxEnd && pathIdxStart <= pathIdxEnd {
			patIdxTmp := -1
			for i:=pattIdxStart + 1;i<= pattIdxEnd;i++{
				if strings.EqualFold("**",*pattDirs[i]) {
					patIdxTmp = i
					break
				}
			}
			if patIdxTmp == pattIdxStart + 1 {
				// '**/**' situation, so skip one
				pattIdxStart++
				continue
			}
			// Find the pattern between padIdxStart & padIdxTmp in str between
			// strIdxStart & strIdxEnd
			patLength := patIdxTmp - pattIdxStart - 1
			strLength := pathIdxEnd - pathIdxStart + 1
			foundIdx := -1

		strLoop:
			for i:=0;i<= strLength - patLength;i++{
				for j := 0; j < patLength; j++ {
					subPat := pattDirs[pattIdxStart + j + 1]
					subStr := pathDirs[pathIdxStart + i + j]
					if !ant.matchStrings(*subPat, *subStr, uriTemplateVariables) {
						goto strLoop
					}
				}
				foundIdx = pathIdxStart + i
				break
			}

			if foundIdx == -1 {
				return false
			}

			pattIdxStart = patIdxTmp
			pathIdxStart = foundIdx + patLength
		}else {
			break
		}
	}

	for i:=pattIdxStart;i <= pattIdxEnd; i++ {
		if !strings.EqualFold ("**",*pattDirs[i]) {
			return false
		}
	}
	return false
}

/**
* Tokenize the given path pattern into parts, based on this matcher's settings.
* <p>Performs caching based on {@link #setCachePatterns}, delegating to
* {@link #tokenizePath(String)} for the actual tokenization algorithm.
* @param pattern the pattern to tokenize
* @return the tokenized pattern parts
*/
//tokenizePattern default use cache
func (ant *AntPathMatcher) tokenizePattern(pattern string) []*string{
	tokenized := make([]*string,0)
	//The first step is to fetch from the cache map.
	value,ok := ant.tokenizedPatternCache.Load(pattern)
	if ok {
		tokenized = value.([]*string)
	}else{
		//No records was fetched from the cache map.
		tokenized = ant.tokenizePath(pattern)
		//add
		if tokenized != nil {
			ant.tokenizedPatternCache.Store(pattern,tokenized)
		}
	}
	return tokenized
}

//tokenizePath
func (ant *AntPathMatcher) tokenizePath(path string) []*string{
	return TokenizeToStringArray(path, ant.pathSeparator,ant.trimTokens,true)
}

//isPotentialMatch
func (ant *AntPathMatcher) isPotentialMatch(path string,pattDirs []*string) bool{
	if !ant.trimTokens {
		pos := 0
		for _,pattDir := range pattDirs {
			skipped := ant.skipSeparator(path, pos, ant.pathSeparator)
			pos += skipped
			skipped = ant.skipSegment(path, pos, *pattDir)
			if skipped < utf8.RuneCountInString(*pattDir) {
				tempPattDir := rune((*pattDir)[0])
				return skipped > 0 || utf8.RuneCountInString(*pattDir) > 0 && ant.isWildcardChar(tempPattDir)
			}
			pos += skipped
		}
	}
	return true
}

//skipSegment
func (ant *AntPathMatcher) skipSegment(path string,pos int,prefix string) int {
	skipped := 0
	for i := 0; i < utf8.RuneCountInString(prefix); i++ {
		c := rune(prefix[i])
		if ant.isWildcardChar(c) {
			return skipped
		}
		currPos := pos + skipped
		if currPos >= utf8.RuneCountInString(path) {
			return 0
		}
		if c == rune(path[currPos]) {
			skipped++
		}
	}
	return skipped
}

//skipSeparator
func (ant *AntPathMatcher) skipSeparator(path string,pos int,separator string) int{
	skipped := 0
	for {
		if StartsWith (path,separator,pos + skipped) {
			skipped += utf8.RuneCountInString(separator)
		}else {
			break
		}
	}
	return skipped
}

//isWildcardChar
func (ant *AntPathMatcher) isWildcardChar(c rune) bool{
	for _,candidate := range WildcardChars {
		if c == candidate {
			return true
		}
	}
	return false
}

/**
* Test whether or not a string matches against a pattern.
*
* @param pattern the pattern to match against (never {@code null})
* @param str     the String which must be matched against the pattern (never {@code null})
* @return {@code true} if the string matches against the pattern, or {@code false} otherwise
*/
//matchStrings
func (ant *AntPathMatcher) matchStrings(pattern,str string,uriTemplateVariables *map[string]string) bool{
	return ant.getStringMatcher(pattern).matchStrings(str, uriTemplateVariables)
}

/**
* Build or retrieve an {@link AntPathStringMatcher} for the given pattern.
* <p>The default implementation checks this AntPathMatcher's internal cache
* (see {@link #setCachePatterns}), creating a new AntPathStringMatcher instance
* if no cached copy is found.
* <p>When encountering too many patterns to cache at runtime (the threshold is 65536),
* it turns the default cache off, assuming that arbitrary permutations of patterns
* are coming in, with little chance for encountering a recurring pattern.
* <p>This method may be overridden to implement a custom cache strategy.
*
* @param pattern the pattern to match against (never {@code null})
* @return a corresponding AntPathStringMatcher (never {@code null})
* @see #setCachePatterns
*/
//getStringMatcher
func (ant *AntPathMatcher) getStringMatcher(pattern string) *AntPathStringMatcher{
	var matcher *AntPathStringMatcher
	cachePatterns := ant.cachePatterns
	if cachePatterns{
		value,ok := ant.stringMatcherCache.Load(pattern)
		if ok && value != nil {
			matcher = value.(*AntPathStringMatcher)
		}
	}
	if matcher == nil {
		matcher = NewStringMatcher(pattern)
		if cachePatterns {
			ant.stringMatcherCache.Store(pattern, matcher)
		}
	}
	return matcher
}