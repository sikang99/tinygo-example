#
# Makefile
#
PROG=esp32-blink
TARGET=esp32-coreboard-v2
usage:
	@echo "usage: make [build|flash] of $(PROG) for $(TARGET)"

build b:
	tinygo build -o $(PROG).uf2 -target=$(TARGET)  main.go

flash f:
	echo $(TARGET)
	tinygo flash -target=$(TARGET) -port /dev/tty.usbserial-0001 -x main.go

git-update gu:
	cd ..; make gu



