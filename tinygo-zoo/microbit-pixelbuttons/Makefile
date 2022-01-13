#
# Makefile
#
PROG=microbit-pixelbuttons
usage:
	@echo "usage: make [build]"

build b:
	tinygo build -o $(PROG).hex -target microbit main.go

flash f:
	tinygo flash -target microbit main.go


