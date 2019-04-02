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


	reg = regexp.MustCompile("tes.*aa")
	t.Log(reg.MatchString("test"))
	t.Log(reg.MatchString("testaa"))
	t.Log(reg.MatchString("testteaab"))
	//rxp := &syntax.Regexp{}


}
