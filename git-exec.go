package main

import (
	"fmt"
	"log"
	"os/exec"
	// "strings"
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
}


func execCommand(pwd string, cmd string, args []string){
	fmt.Printf("lib.Main: cmd: %v\n", cmd)
	fmt.Printf("lib.Main: args: %v\n", args)

	cmdPath, pathErr := exec.LookPath(cmd)
	if pathErr != nil {
		log.Fatal(pathErr)
	}

	args = append([]string{cmd}, args...)

	chdirErr := os.Chdir(pwd)
	if chdirErr != nil {
		log.Fatal(chdirErr)
	}

	env := os.Environ()
	execErr := syscall.Exec(cmdPath, args, env)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
