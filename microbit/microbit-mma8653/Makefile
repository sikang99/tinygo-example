#
# Makefile
#
PROG=microbit-mma8653
usage:
	@echo "usage: make [build]"

edit e:
	vi main.go

build b:
	tinygo build -o $(PROG).hex -target microbit main.go

flash f:
	tinygo flash -target microbit main.go

clean:
	rm -f *.hex *.uf2

git-update gu:
	cd ..; make gu


