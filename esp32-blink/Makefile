#
# Makefile
#
PROG=esp32-blink
TARGET=teensy40
usage:
	@echo "usage: make [build|flash] of $(PROG) for $(TARGET)"

build b:
	tinygo build -o $(PROG).uf2 -target=$(TARGET) main.go

flash f:
	@echo "> make (flash) [1|2]"

f1:
	tinygo flash -target=esp32-coreboard-v2 -port /dev/tty.usbserial-0001 -x main.go
f2:
	tinygo flash -target=esp32 -x main.go
f3:
	tinygo flash -target=esp32c3 -x main.go
f4:
	tinygo flash -target=arduino -x main.go
f5:
	tinygo flash -target=teensy40 -x main.go

git-update gu:
	cd ..; make gu



