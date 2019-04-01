package antpath

import (
	"fmt"
	"testing"
)

//TestTokenizeToStringArray
func TestTokenizeToStringArray(t *testing.T) {
	tokens := TokenizeToStringArray("/bla /**/**/bla/","/",false,false)
	for _,item :=range tokens  {
		fmt.Println(*item)
	}
}

func TestIsBlank(t *testing.T) {
	t.Log(IsBlank(""))
	t.Log(IsBlank(" "))
	t.Log(IsBlank("		"))
	t.Log(IsBlank(" t"))
	t.Log(IsBlank("t "))
	t.Log(IsBlank("t t"))
	t.Log(IsBlank("tt"))

}
