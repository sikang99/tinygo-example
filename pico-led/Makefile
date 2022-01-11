#
# Makefile
#
PROG=pico-led
usage:
	@echo "usage: make [build]"

build b:
	tinygo build -o $(PROG).bin -target=pico main.go

flash f:
	tinygo flash -target=pico main.go


