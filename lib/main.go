package lib

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"os"
	"syscall"
)

func Main(revision string, withStaged bool, cmd string, args []string){
	fmt.Printf("lib.Main: revision: %s\n", revision)
	fmt.Printf("lib.Main: withStaged: %v\n", withStaged)
	fmt.Printf("lib.Main: args: %v\n", args)

	head_revision := gitHeadRevision()
	fmt.Printf("lib.Main: %s\n", head_revision)

	git_toplevel := gitToplevel()
	fmt.Printf("lib.Main: %s\n", git_toplevel)

	execCommand(git_toplevel, cmd, args)
	cmd_tgt := exec.Command(cmd, args...)
	cmd_tgt.Dir = git_toplevel
	out_tgt, err_tgt := cmd_tgt.Output()
	if err_tgt != nil {
		log.Fatal(err_tgt)
	}
	fmt.Printf("lib.Main: %s\n", strings.TrimSpace(string(out_tgt)))
}


func execCommand(pwd string, cmd string, args []string){
	fmt.Printf("lib.Main: cmd: %v\n", cmd)
	fmt.Printf("lib.Main: args: %v\n", args)
	cmdPath, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatal(err)
	}

	env := os.Environ()
	exec_err := syscall.Exec(cmdPath, args, env)
	if exec_err != nil {
		log.Fatal(exec_err)
	}
}
