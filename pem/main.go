package main

import (
	"bytes"
	"encoding/pem"
	"fmt"
)

func main() {
	block := &pem.Block{
		Type:  "Meles Block",
		Bytes: []byte("this is exmaple of encoded data."),
	}

	var buf bytes.Buffer

	err := pem.Encode(&buf, block)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())

	data := []byte(`-----BEGIN Meles Block-----
dGhpcyBpcyBleG1hcGxlIG9mIGVuY29kZWQgZGF0YS4=
-----END Meles Block-----`)

	// block, rest := pem.Decode(data)
	// if block != nil {
	//         fmt.Printf("Block Type: %s\n", block.Type)
	//         fmt.Printf("Block Bytes: %s\n", block.Bytes)
	// } else {
	//         fmt.Println("No PEM block found.")
	// }

	// secondBlock, rest2 := pem.Decode(rest)

	// if secondBlock != nil {
	//         fmt.Printf("Second block type: %s\n", secondBlock.Type)
	//         fmt.Printf("Second block data: %s\n", secondBlock.Bytes)
	// }

	// fmt.Printf("Remaining data: %s\n", rest2)

	block, rest := pem.Decode(data)
	if block != nil {
		fmt.Printf("Block Type: %s\n", block.Type)
		fmt.Printf("Block Bytes:%s\n", block.Bytes)
	} else {
		fmt.Println("Block is not found.")
	}

	block, rest = pem.Decode(rest)
	if block != nil {
		fmt.Printf("Block type:%s\n", block.Type)
		fmt.Printf("Block bytes:%s\n", block.Bytes)
	}
	fmt.Println("Remaining data:", rest)
}
