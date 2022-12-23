.PHONY: upgrade test

upgrade:
	go get -u all

test:
	go test ./...
