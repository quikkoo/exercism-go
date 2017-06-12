SHELL = /bin/sh

export GOPATH = $(HOME)

MODULES = \
	anagram     \
	bob         \
	etl         \
	hello-world \
	word-count

all:
	@for MODULE in $(MODULES);  \
	do                          \
		$(MAKE) -C $$MODULE $@; \
	done

check:
	@for MODULE in $(MODULES);  \
	do                          \
		$(MAKE) -C $$MODULE $@; \
	done

test:
	@for MODULE in $(MODULES);  \
	do                          \
		$(MAKE) -C $$MODULE $@; \
	done

clean:
	@for MODULE in $(MODULES);  \
	do                          \
		$(MAKE) -C $$MODULE $@; \
	done
