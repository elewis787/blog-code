clean: 
	find ./ -name '*.wasm*' -delete 

go-wasm: 
	GOOS=js GOARCH=wasm go build -o ./web/sha256.wasm ./wasm/main_js.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./web/

go-wasm-strip-debug: 
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o ./web/sha256.wasm ./wasm/main_js.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./web/

go-wasm-opt:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o ./web/sha256.wasm ./wasm/main_js.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./web/
	wasm-opt -Oz -o ./web/sha256.wasm.opt ./web/sha256.wasm

tiny-wasm: 
	tinygo build -target wasm -o ./web/sha256.wasm.tiny ./wasm/main_js.go 
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./web/wasm_exec_tiny.js

tiny-wasm-opt:
	tinygo build -target wasm -o ./web/sha256.wasm.tiny ./wasm/main_js.go
	cp "$(shell tinygo env TINYGOROOT)/targets/wasm_exec.js" ./web/wasm_exec_tiny.js
	wasm-opt -Oz -o ./web/sha256.wasm.tiny.opt ./web/sha256.wasm.tiny 
