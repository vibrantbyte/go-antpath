/**
 * Created by GoLand.
 * Brief: matcher fields read/write
 * User: vibrant
 * Date: 2019/04/01
 * Time: 18:05
 */
package antpath

import (
	"strings"
)

/**
* Value class that holds information about the pattern, e.g. number of
* occurrences of "*", "**", and "{" pattern elements.
*/
//PatternInfo
type PatternInfo struct {
	//not null
	pattern string
	uriVars int
	singleWildcards int
	doubleWildcards int
	catchAllPattern bool
	prefixPattern bool
	//not null
	length int

}

//NewDefaultPatternInfo
func NewDefaultPatternInfo(pattern string) *PatternInfo {
	hastext := HasText(pattern)
	// 实例化
	pi := &PatternInfo{}
	pi.pattern = pattern
	if hastext {
		pi.initCounters()
		pi.catchAllPattern = strings.EqualFold("/**",pattern)
		pi.prefixPattern =  pi.catchAllPattern && strings.HasSuffix(pi.pattern,"/**")
	}
	if pi.uriVars == 0 {
		if hastext{
			pi.length = len(pattern)
		}else {
			pi.length = 0
		}
	}
	return pi
}

//initCounters
func (pi *PatternInfo) initCounters(){
	pos := 0
	if HasText(pi.pattern) {
		for  {
			if pos < len(pi.pattern) {
				if rune(pi.pattern[pos]) == Brackets {
					pi.uriVars ++
					pos++
				}else if rune(pi.pattern[pos]) == Asterisk{
					if pos + 1 < len(pi.pattern) && rune(pi.pattern[pos + 1]) == Asterisk {
						pi.doubleWildcards++
						pos += 2
					} else if pos > 0 && !strings.EqualFold(".*",pi.pattern[pos - 1:]) {
						pi.singleWildcards++
						pos++
					} else {
						pos++
					}
				}else {
					pos++
				}
			}else{
				break
			}
		}
	}
}

//GetUriVars
func (pi *PatternInfo) GetUriVars() int {
	return pi.uriVars
}

//GetSingleWildcards
func (pi *PatternInfo) GetSingleWildcards() int {
	return pi.singleWildcards
}

//GetDoubleWildcards
func (pi *PatternInfo) GetDoubleWildcards() int{
	return pi.doubleWildcards
}

//IsLeastSpecific
func (pi *PatternInfo) IsLeastSpecific() bool {
	return IsBlank(pi.pattern) || pi.catchAllPattern
}

//IsPrefixPattern
func (pi *PatternInfo) IsPrefixPattern() bool {
	return pi.prefixPattern
}

//GetTotalCount
func (pi *PatternInfo) GetTotalCount() int {
	return pi.uriVars + pi.singleWildcards + (2 * pi.doubleWildcards)
}

/**
 * Returns the length of the given pattern, where template variables are considered to be 1 long.
 */
func (pi *PatternInfo) GetLength() int {
	if pi.length == 0 {
		if HasText(pi.pattern)  {
			target := VariablePattern.ReplaceAllString(pi.pattern,"#")
			pi.length = len(target)
		}
	}
	return pi.length
}