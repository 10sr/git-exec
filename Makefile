.PHONY: run

run:
	go run *.go -r HEAD~~ sh -c "pwd && sleep 3 &&  echo a && ls && git show | cat"
	go run *.go -s sh -c "pwd && git show | cat"
