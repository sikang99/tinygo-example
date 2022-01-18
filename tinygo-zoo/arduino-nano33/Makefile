#
# Makefile
#
PROG=arduino-nano33
usage:
	@echo "usage: make [build]"

build b:
	tinygo build -o $(PROG).hex -target arduino-nano33 main.go

flash f:
	tinygo flash -target arduino-nano33 main.go

clean:
	rm -f *.hex




