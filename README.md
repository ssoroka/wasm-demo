# WebAssembly demo

`git clone git@github.com:ssoroka/wasm-demo.git $GOPATH/src/github.com/ssoroka/demo`

`cd $GOPATH/src/github.com/ssoroka/demo`

Build the client

`GOARCH=wasm GOOS=js go build -o client/public/main.wasm client/src/main.go`

Run the server

`go run server\main.go`

view it

`open http://localhost:8080`

## References

This demo was largely based off of this blog post: (Compiling Go to WebAssembly)[https://www.sitepen.com/blog/compiling-go-to-webassembly/]