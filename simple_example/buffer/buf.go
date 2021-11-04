package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Fprint(&b, " World !")
	io.Copy(os.Stdout, &b)
}
