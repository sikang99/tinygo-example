#
# Makefile
#
PROG=pico-led
usage:
	@echo "usage: make [build]"

build b:
	tinygo build -o $(PROG).uf2 -target=pca10040 main.go

flash f:
	tinygo flash -target=pca10040 main.go


