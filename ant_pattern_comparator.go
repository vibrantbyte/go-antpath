package antpath

import "strings"

/**
* The default {@link Comparator} implementation returned by
* {@link #getPatternComparator(String)}.
* <p>In order, the most "generic" pattern is determined by the following:
* <ul>
* <li>if it's null or a capture all pattern (i.e. it is equal to "/**")</li>
* <li>if the other pattern is an actual match</li>
* <li>if it's a catch-all pattern (i.e. it ends with "**"</li>
* <li>if it's got more "*" than the other pattern</li>
* <li>if it's got more "{foo}" than the other pattern</li>
* <li>if it's shorter than the other pattern</li>
* </ul>
*/
//AntPatternComparator
type AntPatternComparator struct {
	//not null
	path string
}

//NewDefaultAntPatternComparator
func NewDefaultAntPatternComparator(path string) *AntPatternComparator {
	comparator := &AntPatternComparator{}
	comparator.path = path
	return comparator
}

//Compare
func (comparator *AntPatternComparator) Compare(pattern1,pattern2 string) int{
	info1 := NewDefaultPatternInfo(pattern1)
	info2 := NewDefaultPatternInfo(pattern2)

	if info1.IsLeastSpecific() && info2.IsLeastSpecific() {
		return 0
	} else if info1.IsLeastSpecific() {
		return 1
	} else if info2.IsLeastSpecific() {
		return -1
	}

	pattern1EqualsPath := strings.EqualFold(comparator.path,pattern1)
	pattern2EqualsPath := strings.EqualFold(comparator.path,pattern2)
	if pattern1EqualsPath && pattern2EqualsPath {
		return 0
	} else if pattern1EqualsPath {
		return -1
	} else if pattern2EqualsPath {
		return 1
	}

	if info1.IsPrefixPattern() && info2.GetDoubleWildcards() == 0 {
		return 1
	} else if info2.IsPrefixPattern() && info1.GetDoubleWildcards() == 0 {
		return -1
	}

	if info1.GetTotalCount() != info2.GetTotalCount() {
		return info1.GetTotalCount() - info2.GetTotalCount()
	}

	if info1.GetLength() != info2.GetLength() {
		return info2.GetLength() - info1.GetLength()
	}

	if info1.GetSingleWildcards() < info2.GetSingleWildcards() {
		return -1
	} else if info2.GetSingleWildcards() < info1.GetSingleWildcards() {
		return 1
	}

	if info1.GetUriVars() < info2.GetUriVars() {
		return -1
	} else if info2.GetUriVars() < info1.GetUriVars() {
		return 1
	}

	return 0
}