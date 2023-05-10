// https://leetcode.com/problems/rotate-array/
// Arrays are pass by reference by default
// fixed size array
// dynamic array using slice

package main

import "fmt"

//using memory
func rotateArray(nums []int, k int) []int{

	n := len(nums)
	k = k%n
	if k == 0 {
		return nums
	}
	
	res:=make([]int, n)

	for i:=0;i<n;i++ {
		res[(i+k)%n] =  nums[i]
	}

	return res
}
func reverseArray(nums []int, start int, end int) {

	for start<end {
		nums[start], nums[end] = nums[end], nums[start]
		start++ 
		end--
	}
	
}
//in-place
func rotateArrayInplace(nums []int, k int)  {
	n := len(nums)
    k %= n
	if k==0 || len(nums)<=1{
		return 
	}
	reverseArray(nums, 0, n-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, n-1)
}

func main() {

	//fixed size array
	// var nums [5]int 

	//dynamic array using slice
	variableArray:=[]int{1,2,3,4,5}

	// add new elements to slice
	variableArray = append(variableArray, 6, 7)

	// fmt.Println(variableArray)

	// fmt.Println(variableArray, rotateArray(variableArray, 2))
	// variableArray = []int {1,2,3,4,5,6,7}


	fmt.Println("Given Array: ", variableArray, "k:", 3)
	rotateArrayInplace(variableArray, 3)
	fmt.Println("Result: ", variableArray)

}