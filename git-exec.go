package main

import (
	"flag"
	"github.com/10sr/git-exec/lib"
)

var num *int = flag.Int("n", 0, "Number")

func main(){
	flag.Parse()
	lib.Main(num)
}
