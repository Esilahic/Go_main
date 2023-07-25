package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {

	p := person{
		first: "Jane",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
}

// b := bytes.NewBufferString("Jello ")
// fmt.Println(b.String())
// b.WriteString("Bellow ")
// fmt.Println(b.String())
// b.Reset()
// b.WriteString("Wow! ")
// fmt.Println(b.String())

// b.Write([]byte("Jam!"))
// fmt.Println(b.String())
