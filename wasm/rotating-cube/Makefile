#
# Makefile for tinygo wasm
#
usage:
	@echo "usage: make [build]"

build b:
	tinygo build -target wasm -no-debug -o docs/wasm.wasm main.go

strip s:
	wasm2wat docs/wasm.wasm -o docs/wasm.wat
	wat2wasm docs/wasm.wat -o docs/wasm.wasm
	rm -f docs/wasm.wat



