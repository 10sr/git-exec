package main

import (
	"os/exec"
	"strings"
	"fmt"
)

func gitToplevel() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func gitHeadRevision() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func gitCheckoutTo(repository string, revision string, workingDirectory string) error {
	args := []string {"--work-tree=" + workingDirectory, "checkout", revision, "--", "."}
	cmd := exec.Command("git", args...)
	cmd.Dir = repository
	out, err := cmd.Output()
	fmt.Printf("%s\n", out)
	return err
}

func gitCheckStagedDiff() error {
	var err error
	staged := exec.Command("git", "diff", "--cached", "--quiet")
	_, err = staged.Output()
	if err != nil {
		return err
	}

	return nil
}

func gitCheckUnstagedDiff() error {
	var err error
	unstaged := exec.Command("git", "diff", "--quiet")
	_, err = unstaged.Output()
	if err != nil {
		return err
	}

	return nil
}
