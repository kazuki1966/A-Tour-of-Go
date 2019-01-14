package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func convertRot13(char byte, min byte, max byte) byte {
	if (min < char) && (char < max) {
		char += 13
		if max < char {
			char -= 26
		}
	}
	return char
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)

	for i := range b {
		b[i] = convertRot13(convertRot13(b[i], 'A', 'Z'), 'a', 'z')
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
