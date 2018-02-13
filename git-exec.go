package main

import (
	"flag"
	"./lib"
)

var num *int = flag.Int("n", 0, "Number")

func main(){
	flag.Parse()
	lib.Main(num)
}
