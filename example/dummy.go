// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// p := Pair{A: 1, B: true}
	// _ = Pair{1, 2, 3}
	// var f os.file
	// _ = bytes.Buffer{nil, 0}
}

func main1() {
	var x io.ReadWriteCloser = os.Stdin
	f := io.ReadWriteCloser.Close
	f(x)
	fmt.Printf("%T\n", f)
}