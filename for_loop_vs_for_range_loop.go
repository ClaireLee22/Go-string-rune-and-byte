package main

import (
	"fmt"
)

func main() {
	str := "Hello GO!"

	// index a string yields bytes
	fmt.Println("for loop")
	for i := 0; i < len(str); i++ {
		fmt.Printf("index: %d, value %v, type: %T\n", i, str[i], str[i])
	}

	/* output:
	index: 0, value 72, type: uint8
	index: 1, value 101, type: uint8
	index: 2, value 108, type: uint8
	index: 3, value 108, type: uint8
	index: 4, value 111, type: uint8
	index: 5, value 32, type: uint8
	index: 6, value 71, type: uint8
	index: 7, value 79, type: uint8
	index: 8, value 33, type: uint8
	*/

	// for range loop decodes one UTF-8-encoded rune on each iteration
	fmt.Println("for range loop")
	for index, value := range str {
		fmt.Printf("index: %d, value %v, type: %T\n", index, value, value)
	}

	/* output:
	index: 0, value 72, type: int32
	index: 1, value 101, type: int32
	index: 2, value 108, type: int32
	index: 3, value 108, type: int32
	index: 4, value 111, type: int32
	index: 5, value 32, type: int32
	index: 6, value 71, type: int32
	index: 7, value 79, type: int32
	index: 8, value 33, type: int32
	*/

	// convert rune to byte
	fmt.Println("convert rune to byte")
	for index, value := range str {
		convertValueFromRuneToByte := byte(value)
		fmt.Printf("index: %d, value %v, type: %T\n", index, convertValueFromRuneToByte, convertValueFromRuneToByte)
	}

	/* output:
	index: 1, value 101, type: uint8
	index: 2, value 108, type: uint8
	index: 3, value 108, type: uint8
	index: 4, value 111, type: uint8
	index: 5, value 32, type: uint8
	index: 6, value 71, type: uint8
	index: 7, value 79, type: uint8
	index: 8, value 33, type: uint8
	*/
}
