package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {

	for {
		n, err := r13.r.Read(b)
		if err != nil {
			return 0, io.EOF
		}
		for i := 0; i < n; i++ {
			if b[i] == ' ' || b[i] == '!' {
				b[i] = b[i]
			} else if b[i] < 'l' {
				b[i] = b[i] + 13
			} else {
				b[i] = b[i] - 13
			}
		}
		return n, nil
	}

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
