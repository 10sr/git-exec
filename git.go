package main

import (
	"log"
	"os/exec"
	"strings"
)

func gitToplevel() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

func gitHeadRevision() string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

// func gitGetOutput(arg ...string) string {
// 	cmd := exec.Command("git", arg...)
// 	out, err := cmd.Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return string(out)
// }
