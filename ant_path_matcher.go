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

	//-------private---------//
	// 忽略大小写（不区分大小写）
	//caseSensitive default value = true
	caseSensitive bool
	//trimTokens default value = false
	trimTokens bool
	//cachePatterns default value = true
	cachePatterns bool
}

//New
func New() *AntPathMatcher {
	ant := NewS(DefaultPathSeparator)
	return ant
}

//NewS
func NewS(separator string) *AntPathMatcher{
	if separator == "" {
		return nil
	}
	ant := &AntPathMatcher{}
	ant.pathSeparator = separator
	ant.tokenizedPatternCache = new(sync.Map)
	ant.stringMatcherCache = new(sync.Map)
	ant.caseSensitive = true
	ant.trimTokens = false
	ant.cachePatterns = true
	return ant
}

//IsPattern
func (ant *AntPathMatcher) IsPattern(path string) bool{
	return strings.Index(path,"*")!=-1 || strings.Index(path,"?")!=-1
}

//Match
func (ant *AntPathMatcher) Match(pattern,path string) bool{
	return false
}
