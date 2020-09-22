package main

import(
	"encoding/base64"
	"fmt"
)
func main() {
	str := "bagus kurniawan"
	var encodedString = base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encodedString)

	var decodedByte, _ =base64.StdEncoding.DecodeString(encodedString)
	fmt.Println("decoded:", string(decodedByte))
}

