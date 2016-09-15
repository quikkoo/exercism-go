SHELL = /bin/sh

export GOPATH = $(HOME)

MODULES += anagram
MODULES += bob
MODULES += etl
MODULES += hello-world
MODULES += word-count

all:
	$(MAKE) -C anagram $@
	$(MAKE) -C bob $@
	$(MAKE) -C etl $@
	$(MAKE) -C hello-world $@
	$(MAKE) -C word-count $@

check:
	$(MAKE) -C anagram $@
	$(MAKE) -C bob $@
	$(MAKE) -C etl $@
	$(MAKE) -C hello-world $@
	$(MAKE) -C word-count $@

test:
	$(MAKE) -C anagram $@
	$(MAKE) -C bob $@
	$(MAKE) -C etl $@
	$(MAKE) -C hello-world $@
	$(MAKE) -C word-count $@

clean:
	$(MAKE) -C anagram $@
	$(MAKE) -C bob $@
	$(MAKE) -C etl $@
	$(MAKE) -C hello-world $@
	$(MAKE) -C word-count $@
