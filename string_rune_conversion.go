package main

import (
	"fmt"
)

func main() {
	str := "Hello GO!"

	// 1. convert string to runes
	runes := []rune(str)
	fmt.Println("convert string to runes: ", runes)

	// 2. convert runes to string
	strFromRunes := string(runes)
	fmt.Println("convert runes to string: ", strFromRunes)

	if str != strFromRunes {
		fmt.Println("strings are not matched")
	} else {
		fmt.Println("strings are matched")
	}

	/* output:
	convert string to runes:  [72 101 108 108 111 32 71 79 33]
	convert runes to string:  Hello GO!
	strings are matched
	*/
}
