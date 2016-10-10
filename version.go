package main

import "fmt"

func getVersion() string {
	return "0.5"
}

func cmdVersion() {
	fmt.Println(getVersion())
}
