.PHONY: run

run:
	go run *.go sh -c "pwd && sleep 3 &&  echo a && ls"
