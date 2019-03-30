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
	//tokenize pattern cache(thread safe cache)
	tokenizedPatternCache *sync.Map

	//caseSensitive default value
	caseSensitive bool
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
	ant.caseSensitive = true
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
