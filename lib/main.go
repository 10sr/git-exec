package lib

import (
	"fmt"
	"log"
	"os/exec"
)

func Main(num *int){
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lib.Main: %s\n", out);
}
