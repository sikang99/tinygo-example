#
# Makefile
#
PROG=arduino-blink
TARGET=arduino-nano
FORMAT=uf2
#
usage:
	@echo "usage: make [build|flash] of $(PROG) for $(TARGET)"

build b:
	tinygo build -o $(PROG).cw$(FORMAT) -target=$(TARGET) main.go

flash f:
	tinygo flash -target=$(TARGET) main.go

git-update gu:
	cd ..; make gu

