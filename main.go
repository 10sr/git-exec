package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"os"
	"syscall"
)

func GitExec(revision string, withStaged bool, cmd string, args []string){
	fmt.Printf("lib.Main: revision: %s\n", revision)
	fmt.Printf("lib.Main: withStaged: %v\n", withStaged)
	fmt.Printf("lib.Main: args: %v\n", args)

	headRevision := gitHeadRevision()
	fmt.Printf("lib.Main: %s\n", headRevision)

	gitToplevel := gitToplevel()
	fmt.Printf("lib.Main: %s\n", gitToplevel)

	execCommand(gitToplevel, cmd, args)
	cmdTgt := exec.Command(cmd, args...)
	cmdTgt.Dir = gitToplevel
	outTgt, errTgt := cmdTgt.Output()
	if errTgt != nil {
		log.Fatal(errTgt)
	}
	fmt.Printf("lib.Main: %s\n", strings.TrimSpace(string(outTgt)))
}


func execCommand(pwd string, cmd string, args []string){
	fmt.Printf("lib.Main: cmd: %v\n", cmd)
	fmt.Printf("lib.Main: args: %v\n", args)
	cmdPath, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatal(err)
	}

	args = append([]string{cmd}, args...)

	env := os.Environ()
	execErr := syscall.Exec(cmdPath, args, env)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
