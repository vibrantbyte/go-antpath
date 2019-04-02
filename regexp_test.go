package main

import (
	"regexp"
	"testing"
)

//TestMatcher
func TestMatcher(t *testing.T)  {


	reg := regexp.MustCompile("tes.")
	t.Log(reg.MatchString("test"))
	t.Log(reg.MatchString("testt"))
	t.Log(reg.MatchString("testtee"))


	reg = regexp.MustCompile("t.e.*s.*$")
	t.Log(reg.FindAllIndex([]byte("testteaabtestteaa"),2))
	t.Log(reg.MatchString("testaa"))
	t.Log(reg.MatchString("testteaab"))
	//rxp := &syntax.Regexp{}
}
