# Python parameters
COMPOSECMD=docker-compose

.PHONY: test

all: | build run

setup:	;
build:
		$(COMPOSECMD) build
test:	
		$(COMPOSECMD) up --abort-on-container-exit
clean: ;
run:
		$(COMPOSECMD) up