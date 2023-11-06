package argon2

import (
	"fmt"
	"testing"
)

func Test_Encode(t *testing.T) {
	data := []byte("0000")
	encode := Encode(data, nil)
	fmt.Println(string(encode))

	// input := []byte("1234")
	// enter := Compare(input, encode, nil)
	// fmt.Println("isEnter:", enter)
}
