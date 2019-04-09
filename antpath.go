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
func IsPattern(path *C.char) bool {
	return pMatcher.IsPattern(C.GoString(path))
}

//export Match
func Match(pattern,path string) bool{
	return pMatcher.Match(pattern,path)
}

//export MatchStart
func MatchStart(pattern,path *C.char) bool{
	return pMatcher.MatchStart(C.GoString(pattern),C.GoString(path))
}

//export ExtractPathWithinPattern
func ExtractPathWithinPattern(pattern,path *C.char) *C.char {
	return C.CString(pMatcher.ExtractPathWithinPattern(C.GoString(pattern),C.GoString(path)))
}

//export ExtractUriTemplateVariables
func ExtractUriTemplateVariables(pattern,path *C.char) *map[string]string {
	return pMatcher.ExtractUriTemplateVariables(C.GoString(pattern),C.GoString(path))
}

////export GetPatternComparator
//func GetPatternComparator(path string) *AntPatternComparator {
//	return pMatcher.GetPatternComparator(path)
//}

//export Combine
func Combine(pattern1,pattern2 *C.char) *C.char {
	return C.CString(pMatcher.Combine(C.GoString(pattern1),C.GoString(pattern2)))
}

//export SetPathSeparator
func SetPathSeparator(pathSeparator *C.char){
	 pMatcher.SetPathSeparator(C.GoString(pathSeparator))
}

//export SetCaseSensitive
func SetCaseSensitive(caseSensitive bool) {
	pMatcher.SetCaseSensitive(caseSensitive)
}

//export SetTrimTokens
func SetTrimTokens(trimTokens bool){
	pMatcher.SetTrimTokens(trimTokens)
}

//export SetCachePatterns
func SetCachePatterns(cachePatterns bool){
	pMatcher.SetCachePatterns(cachePatterns)
}

//main
func main(){}