#
# Makefile
#
PROG=arduino-blink
TARGET=arduino-nano
FORMAT=uf2
#
usage:
	@echo "usage: make [build|flash] of $(PROG) for $(TARGET)"

edit e:
	vi main.go

build b:
	tinygo build -o $(PROG).$(FORMAT) -target=$(TARGET) main.go

flash f:
	tinygo flash -target=$(TARGET) main.go

clean:
	rm -rf *.uf2 *.hex

git-update gu:
	cd ..; make gu

