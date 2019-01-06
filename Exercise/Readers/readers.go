package main

import "golang.org/x/tour/reader"

// MyReader :: Reader 構造体（Read メソッド実装対象）
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
