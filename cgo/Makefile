#
# Makefile for tinygo cgo
#
PROG=test
usage:
	@echo "usage: make [edit|build-run]"

edit e:
	vi main.go

build-run br:
	@echo "> make (build-run) [go|tinygo]"

build-run-go brg:
	go build -o $(PROG) && ./$(PROG) 10000000
	rm -f $(PROG)

build-run-tinygo brt:
	tinygo build -o $(PROG) && ./$(PROG) 10000000
	rm -f $(PROG)

#--------------------------------------------------------
git-update gu:
	cd ..; make gu
