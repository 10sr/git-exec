package main

import (
	"fmt"
	"log"
	"os/exec"
	// "strings"
	"os"
	"path"
	"syscall"
	homedir "github.com/mitchellh/go-homedir"
)

func GitExec(revision string, withStaged bool, cmd string, args []string){
	fmt.Printf("lib.Main: revision: %s\n", revision)
	fmt.Printf("lib.Main: withStaged: %v\n", withStaged)
	fmt.Printf("lib.Main: args: %v\n", args)

	var err error
	var workingDirectory string

	// headRevision, err := gitHeadRevision()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("lib.Main: %s\n", headRevision)

	gitToplevel, err := gitToplevel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lib.Main: %s\n", gitToplevel)

	if revision != "" {
		workingDirectory, err = generateWorkingDirectoryPath(gitToplevel)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("lib.Main: Checking out to %s\n", workingDirectory)
		err = gitCheckoutTo(gitToplevel, revision, workingDirectory)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		workingDirectory = gitToplevel
	}

	err = gitCheckStagedDiff()
	if err != nil {
		fmt.Printf("lib.Main: differentials found.\n")
	}

	err = gitCheckUnstagedDiff()
	if err != nil {
		fmt.Printf("lib.Main: differentials found.\n")
	}

	err = execCommand(workingDirectory, cmd, args)
	if err != nil {
		log.Fatal(err)
	}
}


func execCommand(pwd string, cmd string, args []string) error {
	fmt.Printf("lib.Main: pwd: %v\n", pwd)
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
	return err
}


func generateWorkingDirectoryPath(from string) (string, error) {
	home, err := homedir.Expand("~/.git-exec")
	if err != nil {
		return "", err
	}
	base := path.Base(from)

	return path.Join(home, base), nil
}
