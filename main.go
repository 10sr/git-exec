package main

import (
	flag "github.com/spf13/pflag"
)

func main() {
	var (
		revisionFlag   string
		withStagedFlag bool
	)
	flag.StringVar(&revisionFlag, "revision", "", "Revision id")
	flag.StringVar(&revisionFlag, "r", "", "Revision id")
	flag.BoolVar(&withStagedFlag, "with-staged", false, "Use staged state")
	flag.BoolVar(&withStagedFlag, "s", false, "Use staged state")
	flag.Parse()
	rest := flag.Args()
	cmd := rest[0]
	args := rest[1:]
	GitExec(revisionFlag, withStagedFlag, cmd, args)
}
