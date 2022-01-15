#
# Makefile
#
PROG=esp32-ws2802
usage:
	@echo "usage: make [build]"

build b:
	tinygo build -o $(PROG).uf2 -target=esp32-mini32 main.go

flash f:
	tinygo flash -target=esp32-mini32 main.go


