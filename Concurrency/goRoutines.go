// usage of go 
// go -> creates a new route of execution
// time.Sleep -> waits until t time
// try executing below code without time.Sleep (line 16)

package main

import (
	"fmt"
	"time"
)

func main() {

	go firstName()
	time.Sleep(1*time.Second)
	lastName()
}

func firstName() {
	fmt.Println("Hero Name: Vamshi")
}

func lastName() {
	fmt.Println("Family Name : Pandre")
}