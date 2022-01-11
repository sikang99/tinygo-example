#
# Makefile for gst-go
#
PROG=tinygo-example
BUILD=0.0.0.1
#----------------------------------------------------------------------------------
usage:
	@echo "usage: make [git] for $(PROG)"
#----------------------------------------------------------------------------------
MSG="modifying"
git g:
	@echo "> make (git:g) [tag|update|store]"

git-tag gt:
	git tag $(BUILD)

git-update gu:
	git add .
	git commit -a -m "$(BUILD),$(USER)-$(MSG)"
	git push

git-store gs:
	git config credential.helper store
#----------------------------------------------------------------------------------
