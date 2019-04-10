package main

import (
	"C"
	. "github.com/vibrantbyte/go-antpath/antpath"
)

//pMatcher
var pMatcher PathMatcher

//init
func init(){
	pMatcher = New()
}

//export Version
func Version() *C.char{
	return C.CString("v1.0")
}

//export IsPattern
func IsPattern(path string) bool {
	return pMatcher.IsPattern(path)
}

//export Match
func Match(pattern,path string) bool{
	return pMatcher.Match(pattern,path)
}

//export MatchStart
func MatchStart(pattern,path string) bool{
	return pMatcher.MatchStart(pattern,path)
}

//export ExtractPathWithinPattern
func ExtractPathWithinPattern(pattern,path string) *C.char {
	result := pMatcher.ExtractPathWithinPattern(pattern,path)
	return C.CString(result)
}

//export ExtractUriTemplateVariables
func ExtractUriTemplateVariables(pattern,path string) *map[string]string {
	return pMatcher.ExtractUriTemplateVariables(pattern,path)
}

////export GetPatternComparator
//func GetPatternComparator(path string) *AntPatternComparator {
//	return pMatcher.GetPatternComparator(path)
//}

//export Combine
func Combine(pattern1,pattern2 string) *C.char {
	result := pMatcher.Combine(pattern1,pattern2)
	return C.CString(result)
}

//export SetPathSeparator
func SetPathSeparator(pathSeparator string) bool{
	 pMatcher.SetPathSeparator(pathSeparator)
	 return true
}

//export SetCaseSensitive
func SetCaseSensitive(caseSensitive bool) bool{
	pMatcher.SetCaseSensitive(caseSensitive)
	return true
}

//export SetTrimTokens
func SetTrimTokens(trimTokens bool) bool{
	pMatcher.SetTrimTokens(trimTokens)
	return true
}

//export SetCachePatterns
func SetCachePatterns(cachePatterns bool) bool{
	pMatcher.SetCachePatterns(cachePatterns)
	return true
}

//main
func main(){}