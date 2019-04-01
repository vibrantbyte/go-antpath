/**
 * Created by GoLand.
 * Brief: constant use
 * User: vibrant
 * Date: 2019/03/30
 * Time: 13:10
 */
package antpath

import "regexp"

const(

	//DefaultPathSeparator Default path separator: "/"
	DefaultPathSeparator = "/"

	//CacheTurnoffThreshold
	CacheTurnoffThreshold = 65536

)

//WildcardChars
var WildcardChars []rune
//*
var Asterisk rune
//?
var QuestionMark rune
//{
var Brackets rune

//pattern
var VariablePattern *regexp.Regexp

//initial
func init(){
	Asterisk = '\u002a'
	QuestionMark = '\u003f'
	Brackets = '\u007b'
	//WildcardChars initial '*', '?', '{'
	WildcardChars = []rune{Asterisk,QuestionMark,Brackets}

	//pattern
	reg,err := regexp.Compile("\\{[^/]+?\\}")
	if err == nil {
		VariablePattern = reg
	}
}