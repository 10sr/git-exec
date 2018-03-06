.PHONY: run

run:
	go run *.go -r HEAD~~ sh -c "pwd && sleep 3 &&  echo a && ls"
