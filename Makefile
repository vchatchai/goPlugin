all: run


hello_world: 
	go build -buildmode=plugin -o eng.so ./plugins/eng/engHelloWorld.go
	go build -buildmode=plugin -o thai.so ./plugins/thai/thaiHelloWorld.go

run:  hello_world
	go run $(GOPATH)/src/github.com/vchatchai/goPlugin/main.go

clean:
	rm -f *.so


