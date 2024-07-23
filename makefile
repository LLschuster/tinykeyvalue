
run: build
	./main $(PORT)

build:
	go build main.go utils.go server.go
