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

	var err error

	headRevision, err := gitHeadRevision()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lib.Main: %s\n", headRevision)

	gitToplevel, err := gitToplevel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lib.Main: %s\n", gitToplevel)

	err = execCommand(gitToplevel, cmd, args)
	if err != nil {
		log.Fatal(err)
	}
}


func execCommand(pwd string, cmd string, args []string) error {
	fmt.Printf("lib.Main: cmd: %v\n", cmd)
	fmt.Printf("lib.Main: args: %v\n", args)

	var err error

	cmdPath, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}

	args = append([]string{cmd}, args...)

	err = os.Chdir(pwd)
	if err != nil {
		return err
	}

	env := os.Environ()
	err = syscall.Exec(cmdPath, args, env)
	// Unreachable!
	return err
}
