package main

import (
	"encoding/base64"
	"fmt"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func main() {
	// encode	
	debyte := "aHR0cCUzQS8vcGFuLmJhaWR1LmNvbS9zLzFvNnlwT3ZnJTIw"
	
	// decode	
	enbyte, err := base64Decode([]byte(debyte))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(enbyte))
}
