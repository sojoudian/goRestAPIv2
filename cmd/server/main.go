package main

import "fmt"

// Run - is goin to be responsible for
// the instantiation and startup of the
// go application
fun Run() error {
	fmt.Println("starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REST API v2")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
