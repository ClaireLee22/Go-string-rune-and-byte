package main

import (
	"fmt"
)

func main() {
	str := "Hello GO!"

	// 1. convert string to bytes
	bytes := []byte(str)
	fmt.Println("convert str to bytes: ", bytes)

	// 2. convert bytes to string
	strFromBytes := string(bytes)
	fmt.Println("convert bytes to str: ", strFromBytes)

	if str != strFromBytes {
		fmt.Println("strings are not matched")
	} else {
		fmt.Println("strings are matched")
	}

	/* output:
	convert str to bytes:  [72 101 108 108 111 32 71 79 33]
	convert bytes to str:  Hello GO!
	strings are matched
	*/
}
