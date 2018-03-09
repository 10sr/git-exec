/*
git-exec
========

Execute command in your git repository!


Install
-------

    go get -u github.com/10sr/git-exec


Usage
-----

Without any argument, just run given comand at the root directory of your git
repository (fail if you are not inside of any git repository):

    git exec <command> [<args>...]

With `-r <rev>` option, try to checkout that revision and run command.

    git exec --revision|-r <rev> <command> [<args> ...]

In this case, files are checked out into a working directory first (defaults to
`$HOME/.git-exec/$name.$id`, where `$name` is the base
name of the directory and `$id` is calculated from the directory full path)
and then the command will be invoked.

Also, you can run command with staged files:

    git exec --with-staged|-s <command> [<args> ...]
*/


package main

import (
	"fmt"
	"log"
	"os/exec"
	// "strings"
	"crypto/sha512"
	"encoding/hex"
	homedir "github.com/mitchellh/go-homedir"
	"io"
	"os"
	"path"
	"syscall"
)

// GitExec Execute command in git repository with specified revision.
func GitExec(revision string, withStaged bool, cmd string, args []string) {
	fmt.Printf("lib.Main: revision: %s\n", revision)
	fmt.Printf("lib.Main: withStaged: %v\n", withStaged)
	fmt.Printf("lib.Main: args: %v\n", args)

	if revision != "" && withStaged {
		log.Fatal("revision arg and --with-staged flags are both given at the same time")
	}

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

	if revision != "" || withStaged {
		workingDirectory, err = generateWorkingDirectoryPath(gitToplevel)
		if err != nil {
			log.Fatal(err)
		}

		if revision != "" {
			revision, err = gitRevParse(revision)
			if err != nil {
				log.Fatal(err)
			}
		} else if withStaged {
			revision, err = gitMakeCommitFromStage()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("Unreachable")
		}

		fmt.Printf("lib.Main: Checking out to %s\n", workingDirectory)
		err = gitCheckoutToByClone(gitToplevel, revision, workingDirectory)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		workingDirectory = gitToplevel
	}

	err = gitCheckStagedDiff()
	if err != nil {
		fmt.Printf("lib.Main: Staged differentials found.\n")
	}

	err = gitCheckUnstagedDiff()
	if err != nil {
		fmt.Printf("lib.Main: Unstaged differentials found.\n")
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

	sha := sha512.New()
	io.WriteString(sha, from)
	id := hex.EncodeToString(sha.Sum(nil))[0:6]

	return path.Join(home, base+"."+id), nil
}
