package antpath

import (
	"fmt"
	"regexp"
	"testing"
)

//TestNewStringMatcher
func TestNewStringMatcher(t *testing.T) {
	t.Log(regexp.MatchString("\\d","12122admin"))
	t.Log(regexp.MatchString("\\d","admin"))


	matchedList := GlobPattern.FindAllString("{bane}/women/{名称}/{ddd}",10)
	t.Log(fmt.Sprintf("%v",matchedList))


	matchedListNew := GlobPattern.FindStringSubmatch("{bane}/women/{名称}/{ddd}")
	t.Log(matchedListNew)

	t.Log(1<<5)
	//应该输出结果
	//源：{bane}/women/{名称:测试}/{ddd}
	//目：(.*)\Q/women/\E(测试)\Q/\E(.*)

	/**
	{bane} -> (.*)
	/women/ -> \Q/women/\E
	women -> women
	{名称:测试} -> (测试)
	/ -> \Q/\E
	{ddd} -> (.*)
	 */

	NewMatchesStringMatcher("{bane}/women/{名称}/{ddd}",false)
	NewMatchesStringMatcher("{bane}/women/{name:value}/{ddd}",false)
	NewMatchesStringMatcher("{bane}/women/{名称}/{ddd:111}",false)
}

func TestQuote(t *testing.T){
	t.Log(regexp.QuoteMeta("11wew333我们"))
}
