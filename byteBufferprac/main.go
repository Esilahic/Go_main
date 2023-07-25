package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Jello ")
	fmt.Println(b.String())
	b.WriteString("Bellow ")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("Wow! ")
	fmt.Println(b.String())

	b.Write([]byte("Jam!"))
	fmt.Println(b.String())
}
