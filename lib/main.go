package lib

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Main(revision string, withStaged bool, command string, args []string){
	fmt.Printf("lib.Main: revision: %s\n", revision)
	fmt.Printf("lib.Main: revision: %v\n", withStaged)
	fmt.Printf("lib.Main: revision: %v\n", args)

	head_revision := gitHeadRevision()
	fmt.Printf("lib.Main: %s\n", head_revision)

	git_toplevel := gitToplevel()
	fmt.Printf("lib.Main: %s\n", git_toplevel)

	cmd_tgt := exec.Command(command, args...)
	cmd_tgt.Dir = git_toplevel
	out_tgt, err_tgt := cmd_tgt.Output()
	if err_tgt != nil {
		log.Fatal(err_tgt)
	}
	fmt.Printf("lib.Main: %s\n", strings.TrimSpace(string(out_tgt)))
}
