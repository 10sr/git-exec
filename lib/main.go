package lib

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Main(num *int){
	head_revision := gitHeadRevision()
	fmt.Printf("lib.Main: %s\n", head_revision)

	git_toplevel := gitToplevel()
	fmt.Printf("lib.Main: %s\n", git_toplevel)

	cmd_tgt := exec.Command("pwd")
	cmd_tgt.Dir = git_toplevel
	out_tgt, err_tgt := cmd_tgt.Output()
	if err_tgt != nil {
		log.Fatal(err_tgt)
	}
	fmt.Printf("lib.Main: %s\n", strings.TrimSpace(string(out_tgt)))
}
