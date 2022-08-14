package main

import (
	"fmt"
)

func main() {
	strWithSpecialCharacter := "Héllô GÕ!"
	bytesStrWithSpecialCharacter := []byte(strWithSpecialCharacter)
	runesStrWithSpecialCharacter := []rune(strWithSpecialCharacter)
	fmt.Println("strWithSpecialCharacter", strWithSpecialCharacter)
	fmt.Println("bytesStrWithSpecialCharacter", bytesStrWithSpecialCharacter)
	fmt.Println("runesStrWithSpecialCharacter", runesStrWithSpecialCharacter)

	/* output:
	strWithSpecialCharacter Héllô GÕ!
	bytesStrWithSpecialCharacter [72 195 169 108 108 195 180 32 71 195 149 33]
	runesStrWithSpecialCharacter [72 233 108 108 244 32 71 213 33]
	*/
}
