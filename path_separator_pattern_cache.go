/**
 * Created by GoLand.
 * Brief: matcher fields read/write
 * User: vibrant
 * Date: 2019/04/01
 * Time: 14:57
 */
package antpath

/**
* A simple cache for patterns that depend on the configured path separator.
*/
type PathSeparatorPatternCache struct {
	//"*"
	endsOnWildCard string
	//"**"
	endsOnDoubleWildCard string
}

//NewDefaultPathSeparatorPatternCache 构造函数
func NewDefaultPathSeparatorPatternCache(pathSeparator string) *PathSeparatorPatternCache {
	patternCache := &PathSeparatorPatternCache{}
	patternCache.endsOnWildCard = pathSeparator + "*"
	patternCache.endsOnDoubleWildCard = pathSeparator + "**"
	return patternCache
}

//GetEndsOnWildCard 返回 "*"
func (patternCache *PathSeparatorPatternCache) GetEndsOnWildCard() string {
	return patternCache.endsOnWildCard
}

//GetEndsOnDoubleWildCard 返回 "**"
func (patternCache *PathSeparatorPatternCache) GetEndsOnDoubleWildCard() string{
	return patternCache.endsOnDoubleWildCard
}