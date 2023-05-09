package main

import "fmt"

func isPalindrome(name string) bool {
	for i:=0; i<=len(name)/2; i++ {
		if name[i] != name[len(name)-i-1] {
			return false
		}
	}
	return true
}

func reverseString(inputStr string) string {

	x:=[]rune(inputStr)
	i, j := 0, len(inputStr)-1
	for i<j {
		x[i], x[j] = x[j], x[i]
		i++
		j--
	}
	
	return string(x)
}

func main() {
	var name string

	name = "malayalam"

	fmt.Println(isPalindrome(name))

	reverseStr := "ihsmav erdnap"

	fmt.Println(reverseString(reverseStr))
}