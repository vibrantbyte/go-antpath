/**
 * Created by GoLand.
 * Brief: matcher fields read/write
 * User: vibrant
 * Date: 2019/04/02
 * Time: 18:05
 */
package extend

import (
	"regexp"
)

const (
	Endanchor = 1
	Noanchor  = 0
)

/**
* An engine that performs match operations on a {@linkplain java.lang.CharSequence
* character sequence} by interpreting a {@link Pattern}.
*/
//Matcher implement MatchResult
type Matcher struct {
	/**
     * The Pattern object that created this Matcher.
     */
    //parentPattern
    parentPattern *regexp.Regexp

	/**
	 * The storage used by groups. They may contain invalid values if
	 * a group was skipped during the matching.
	 */
	//groups
	groups []int

	/**
	 * The range within the sequence that is to be matched. Anchors
	 * will match at these "hard" boundaries. Changing the region
	 * changes these values.
	 */
	//from
	from, to int

	/**
	 * Lookbehind uses this value to ensure that the subexpression
	 * match ends at the point where the lookbehind was encountered.
	 */
	//lookbehindTo
	lookbehindTo int

	/**
	 * The original string being matched.
	 */
	//text
	text string

	/**
	 * Matcher state used by the last node. Noanchor is used when a
	 * match does not have to consume all of the input. Endanchor is
	 * the mode used for matching all the input.
	 */
	//acceptMode default value = Noanchor
	acceptMode int

	/**
	 * The range of string that last matched the pattern. If the last
	 * match failed then first is -1; last initially holds 0 then it
	 * holds the index of the end of the last match (which is where the
	 * next search starts).
	 */
	//int first = -1, last = 0;
	first,last int

	/**
	 * The end index of what matched in the last match operation.
	 */
	//int oldLast = -1;
	oldLast int

	/**
	 * The index of the last position appended in a substitution.
	 */
	//int lastAppendPosition = 0;
	lastAppendPosition int

	/**
	 * Storage used by nodes to tell what repetition they are on in
	 * a pattern, and where groups begin. The nodes themselves are stateless,
	 * so they rely on this field to hold state during a match.
	 */
	//locals
	locals []int

	/**
	 * Boolean indicating whether or not more input could change
	 * the results of the last match.
	 *
	 * If hitEnd is true, and a match was found, then more input
	 * might cause a different match to be found.
	 * If hitEnd is true and a match was not found, then more
	 * input could cause a match to be found.
	 * If hitEnd is false and a match was found, then more input
	 * will not change the match.
	 * If hitEnd is false and a match was not found, then more
	 * input will not cause a match to be found.
	 */
	//hitEnd
	hitEnd bool

	/**
	 * Boolean indicating whether or not more input could change
	 * a positive match into a negative one.
	 *
	 * If requireEnd is true, and a match was found, then more
	 * input could cause the match to be lost.
	 * If requireEnd is false and a match was found, then more
	 * input might change the match but the match won't be lost.
	 * If a match was not found, then requireEnd has no meaning.
	 */
	//requireEnd
	requireEnd bool

	/**
	 * If transparentBounds is true then the boundaries of this
	 * matcher's region are transparent to lookahead, lookbehind,
	 * and boundary matching constructs that try to see beyond them.
	 */
	//transparentBounds default value is false
	transparentBounds bool

	/**
	 * If anchoringBounds is true then the boundaries of this
	 * matcher's region match anchors such as ^ and $.
	 */
	//anchoringBounds default value is true
	anchoringBounds bool
}

//NewDefaultMatcher
func NewDefaultMatcher() *Matcher{
	matcher := &Matcher{}
	matcher.acceptMode = Noanchor
	matcher.first = -1
	matcher.last = 0
	matcher.oldLast = -1
	matcher.lastAppendPosition = 0
	matcher.transparentBounds = false
	matcher.anchoringBounds = true
	return matcher
}

//NewMatcher
func NewMatcher(parent *regexp.Regexp,text string) *Matcher{
	matcher := NewDefaultMatcher()
	matcher.parentPattern = parent
	matcher.text = text

	// Allocate state storage
	matcher.groups = make([]int,0)
	matcher.locals = make([]int,0)
	matcher.reset()
	return matcher
}

//reset
func (matcher *Matcher) reset() *Matcher {
	matcher.first = -1
	matcher.last = 0
	matcher.oldLast = -1
	for i := 0; i<len(matcher.groups); i++{
		matcher.groups[i] = -1
	}
	//
	for i:= 0; i<len(matcher.locals); i++{
		matcher.locals[i] = -1
	}
	matcher.lastAppendPosition = 0
	matcher.from = 0
	matcher.to = matcher.getTextLength()
	return matcher
}

//getTextLength
func (matcher *Matcher) getTextLength() int{
	return len(matcher.text)
}

