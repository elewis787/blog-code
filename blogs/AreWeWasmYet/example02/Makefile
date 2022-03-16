wasm: 
	GOOS=js GOARCH=wasm go build -o ./web/counterClient.wasm ./wasm/main_js.go
	cp -p ${GOROOT}/misc/wasm/wasm_exec.js ./web/
wasm-strip-debug: 
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o ./web/counterClient.wasm ./wasm/main_js.go
    cp -p ${GOROOT}/misc/wasm/wasm_exec.js ./web/
wasm-opt:
	GOOS=js GOARCH=wasm go build -o ./web/counterClient.wasm ./wasm/main_js.go
    cp -p ${GOROOT}/misc/wasm/wasm_exec.js ./web/
	wasm-opt -Oz -o ./web/counterClient.wasm.opt ./web/counterClient.wasm
tiny-wasm: 
	tinygo build -o ./web/counterClient.wasm.tiny ./wasm/main_js.go 
	cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./wasm_exec_tiny.js
tiny-wasm-opt:
    tinygo build -o ./web/counterClient.wasm.tiny ./wasm/main_js.go
    cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./wasm_exec_tiny.js
    wasm-opt -Oz -o ./web/counterClient.wasm.opt ./web/counterClient.wasm 
