package lib

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Main(num *int){
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lib.Main: %s\n", strings.TrimSpace(string(out)));

	cmd_getroot := exec.Command("git", "rev-parse", "--show-toplevel")
	out_getroot, err_getroot := cmd_getroot.Output()
	if err_getroot != nil {
		log.Fatal(err_getroot)
	}
	fmt.Printf("lib.Main: %s\n", strings.TrimSpace(string(out_getroot)))
	git_toplevel := strings.TrimSpace(string(out_getroot))

	cmd_tgt := exec.Command("pwd")
	cmd_tgt.Dir = git_toplevel
	out_tgt, err_tgt := cmd_tgt.Output()
	if err_tgt != nil {
		log.Fatal(err_tgt)
	}
	fmt.Printf("lib.Main: %s\n", strings.TrimSpace(string(out_tgt)))
}
