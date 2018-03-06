package main

import (
	"os/exec"
	"strings"
	"fmt"
	"os"
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

func gitRevParse(rev string) (string, error) {
	cmd := exec.Command("git", "rev-parse", rev)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func gitCheckoutToByClone(repository string, revision string, workingDirectory string) error {
	var err error

	// Clone repository
	if _, err = os.Stat(workingDirectory); os.IsNotExist(err) {
		clone := exec.Command("git", "clone", "--shared", repository, workingDirectory)
		out, err := clone.Output()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", out)
	}

	{
		checkout := exec.Command("git", "checkout", "--force", revision)
		checkout.Dir = workingDirectory
		out, err := checkout.Output()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", out)
	}

	{
		checkout := exec.Command("git", "reset", "--hard", "HEAD")
		checkout.Dir = workingDirectory
		out, err := checkout.Output()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", out)
	}

	{
		submodule := exec.Command("git", "submodule", "update", "--init")
		submodule.Dir = workingDirectory
		out, err := submodule.Output()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", out)
	}
	return nil
}

func gitCheckoutTo(repository string, revision string, workingDirectory string) error {
	var err error
	err = os.MkdirAll(workingDirectory, 0755)
	if err != nil {
		return err
	}

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
