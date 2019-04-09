/**
 * Created by GoLand.
 * Brief: pMatcher fields read/write
 * User: vibrant
 * Date: 2019/04/02
 * Time: 14:57
 */
package main

import "fmt"
import . "github.com/vibrantbyte/go-antpath/antpath"


//matchers
var matcher PathMatcher

func init(){
	matcher = New()
}


func main(){
	fmt.Println(matcher.Match("test","test"))
	fmt.Println(matcher.Match("test*aaa", "testblaaaa"))
	fmt.Println(matcher.Match("t?st", "test"))
	fmt.Println(matcher.Match("/{bla}.*", "/testing.html"))
}