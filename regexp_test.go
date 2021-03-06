package main

import (
	"fmt"
	"regexp"
	"strings"
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

	i := 0
	to:
	for ;i< 10 ; i ++  {
		fmt.Println(i)
		if i%2 == 1{
			continue to
		}
	}


}

func TestMatcher01(t *testing.T)  {
	reg := regexp.MustCompile("hotels")
	t.Log(string(reg.Find([]byte("hotels"))))

	t.Log(reg.FindStringSubmatch("hotels"))

	t.Log(strings.Trim(`{name}`,"{}"))

	reg = regexp.MustCompile("/(.*).*")
	t.Log(reg.FindString("/42.html"))

	reg = regexp.MustCompile("^([\\p{L}\\.]+)&")
	t.Log(reg.FindString("/A-b-C"))


	reg = regexp.MustCompile("/A-(.*)-C")
	indexs := reg.FindSubmatch([]byte("/A-b-C"))
	for index := range indexs  {
		t.Log(index)
		t.Log(fmt.Sprint(string(indexs[index])))
	}

}
