/**
 * Created by GoLand.
 * Brief: matcher fields read/write
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
	fmt.Println("")
	fmt.Println(matcher.Match("test","test"))
}