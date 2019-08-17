goroot=$(shell go env GOROOT)

build-minimal:
	GOOS=js GOARCH=wasm go build -o minimal/main.wasm ./minimal
build-hellogo:
	GOOS=js GOARCH=wasm go build -o hellogo/main.wasm ./hellogo
	cp ${goroot}/misc/wasm/wasm_exec.js ./hellogo/wasm_exec.js
build-syscall:
	GOOS=js GOARCH=wasm go build -o syscall/main.wasm ./syscall
	cp ${goroot}/misc/wasm/wasm_exec.js ./syscall/wasm_exec.js
build-square:
	GOOS=js GOARCH=wasm go build -o square/main.wasm ./square
	cp ${goroot}/misc/wasm/wasm_exec.js ./square/wasm_exec.js
build-bindings:
	GOOS=js GOARCH=wasm go build -o bindings/main.wasm ./bindings
	cp ${goroot}/misc/wasm/wasm_exec.js ./bindings/wasm_exec.js
build-server:
	go build -o bin/server ./server
