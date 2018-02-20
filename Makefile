.PHONY: run

run:
	sed -e 's:github.com/10sr/git-exec/lib:./lib:' git-exec.go >git-exec.local.go
	go run git-exec.local.go sh -c "pwd && sleep 10 &&  echo a"
