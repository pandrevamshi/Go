// Build a new String using Strings Builder
// var myStr strings.Builder
// myStr.WriteString("Hello")
// myStr.WriteRune(rune('-'))
// https://leetcode.com/problems/defanging-an-ip-address/description/
// convert inpStr to []rune(inpStr) to use it as an array which can be modified

package main

import (
	"fmt"
	"strings"
)

// using builder
func defangIPaddr(address string) string {
    var ans strings.Builder
    for i:=0; i < len(address);i++ {
        if address[i] == '.' {
            ans.WriteString("[.]")
        } else {
            ans.WriteString(string(address[i]))
			// ans.WriteRune(rune(address[i]))
        }
    }
    return ans.String()
}

// native way using dynamic runes array
func defangIPaddrRunes(address string) string {
    // x:=[]rune(address)
    y := []rune{}
    for i:=0;i<len(address);i++ {
        if address[i] == '.'{
            y = append(y, '[','.',']')
        } else {
            y = append(y, rune(address[i]))
        }
    }
    return string(y)
}
func main() {

	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddrRunes("1.1.1.1"))
	
	var test string
	test = "Hey Pandre!"
	testPlay := []rune(test)
	testPlay[0] = 'K'
	fmt.Println(string(testPlay))

}