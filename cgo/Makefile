#
# Makefile for tinygo cgo
#
usage:
	@echo "usage: make [build|run]"

build-run br:
	@echo "> make (build-run) [go|tinygo]"

build-run-go brg:
	go build -o test && ./test 10000000
	rm -f test

build-run-tinygo brt:
	tinygo build -o test && ./test 10000000
	rm -f test

#--------------------------------------------------------
git-update gu:
	cd ..; make gu
