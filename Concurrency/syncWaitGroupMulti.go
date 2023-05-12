// Provides LOCKS and Syncronization primitives
// Add -> number of Goroutines we wish to wait for. Also, add panic when the inner counter becomes negative
// Done -> Decrements the inner counter by 1 and should be used when a goroutine finishes its work.
// Wait -> The wait blocks the goroutine from which it is invoked until the counter reaches zero. 

package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	// Add - Just a counter -> No of Go Routines you expect to have. Adding 2 here makes you to write Done twice.
	waitGroup.Add(1)
	firstName(&waitGroup)
	// Done - Decrements the waitGroup Counter
	// waitGroup.Done()
	// Wait - Prompts error if the inner counter is non zero
	waitGroup.Wait()
	// lastName()
}


//WaitGroup needs to be passed by reference. By default it is passed by value and defaults the inner counter to '0'
func firstName(waitGroup *sync.WaitGroup) {
	fmt.Println("Hero Name: Vamshi")
	// Decrements the waitGroup Counter
	waitGroup.Add(1)
	defer waitGroup.Done()
	lastName(waitGroup)
}

func lastName(waitGroup *sync.WaitGroup) {
	fmt.Println("Family Name : Pandre")
	defer waitGroup.Done()
}