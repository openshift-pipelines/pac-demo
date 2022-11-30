package main

import (
	"fmt"
	"os"
)

const foo = "bar"

func main() {
	fmt.Fprint(os.Stdout, "Hello world\n")
}
