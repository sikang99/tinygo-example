#
# Makefile
#
PROG=pico-led
TARGET=pico
usage:
	@echo "usage: make [build|flash] of $(PROG) for $(TARGET)"

build b:
	tinygo build -o $(PROG).uf2 -target=$(TARGET) main.go

flash f:
	tinygo flash -target=$(TARGET) main.go


git-update gu:
	cd ..; make gu


