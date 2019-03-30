/**
 * Created by GoLand.
 * Brief: apache ant path matcher implement
 * User: vibrant
 * Date: 2019/03/30
 * Time: 13:10
 */
package antpath

import "strings"

/**
* Actually match the given {@code path} against the given {@code pattern}.
* @param pattern the pattern to match against
* @param path the path String to test
* @param fullMatch whether a full pattern match is required (else a pattern match
* as far as the given base path goes is sufficient)
* @return {@code true} if the supplied {@code path} matched, {@code false} if it didn't
*/
//doMatch
func (ant *AntPathMatcher) doMatch(pattern,path string,fullMatch bool,uriTemplateVariables *map[string]string) bool{
	if strings.HasPrefix(path,DefaultPathSeparator) != strings.HasPrefix(pattern,DefaultPathSeparator) {
		return false
	}
	return false
}

/**
* Tokenize the given path pattern into parts, based on this matcher's settings.
* <p>Performs caching based on {@link #setCachePatterns}, delegating to
* {@link #tokenizePath(String)} for the actual tokenization algorithm.
* @param pattern the pattern to tokenize
* @return the tokenized pattern parts
*/
//tokenizePattern
func (ant *AntPathMatcher) tokenizePattern() []*string{
	tokenized := make([]*string,0)

	return tokenized
}
