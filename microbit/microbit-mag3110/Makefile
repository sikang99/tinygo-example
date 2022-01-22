#
# Makefile
#
PROG=microbit-mag3110
usage:
	@echo "usage: make [build]"

build b:
	tinygo build -o $(PROG).hex -target microbit main.go

flash f:
	tinygo flash -target microbit main.go

clean:
	rm -f *.hex



