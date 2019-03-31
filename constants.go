/**
 * Created by GoLand.
 * Brief: constant use
 * User: vibrant
 * Date: 2019/03/30
 * Time: 13:10
 */
package antpath

const(

	//DefaultPathSeparator Default path separator: "/"
	DefaultPathSeparator = "/"

	//CacheTurnoffThreshold
	CacheTurnoffThreshold = 65536

)

//WildcardChars
var WildcardChars []rune

//initial
func init(){
	//WildcardChars initial '*', '?', '{'
	WildcardChars = []rune{'\u002a','\u003f','\u007b'}
}