package main

import (
	"fmt"
	"os"

	"github.com/thebaer/punctuation"
)

func main() {
	o := punctuation.Extract(os.Stdin)
	fmt.Println(o)
}
