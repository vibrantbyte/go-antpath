/**
 * Created by GoLand.
 * Brief: apache ant path matcher implement
 * User: vibrant
 * Date: 2019/03/30
 * Time: 13:10
 */
package antpath

import "strings"

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

	if (pathIdxStart > pathIdxEnd){
		// Path is exhausted, only match if rest of pattern is * or **'s
		if (pattIdxStart > pattIdxEnd) {
			if strings.HasSuffix(pattern,ant.pathSeparator){
				return strings.HasSuffix(path,ant.pathSeparator)
			}else {
				return !strings.HasSuffix(path,ant.pathSeparator)
			}
		}
		if (!fullMatch) {
			return true
		}
		if (pattIdxStart == pattIdxEnd && strings.EqualFold("*",*pattDirs[pattIdxStart]) && strings.HasSuffix(path,ant.pathSeparator)) {
			return true
		}
		for i:=pattIdxStart;i<=pattIdxEnd;i++  {
			if (!strings.EqualFold("**",*pattDirs[i])) {
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
	return TokenizeToStringArray1(path, ant.pathSeparator)
}

//isPotentialMatch
func (ant *AntPathMatcher) isPotentialMatch(path string,pattDirs []*string) bool{
	return true
}

//matchStrings
func (ant *AntPathMatcher) matchStrings(pattern,str string,uriTemplateVariables *map[string]string) bool{
	return true
}