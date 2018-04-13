# Python parameters
DOCKERCMD=docker
PORT?=8888
TAG?=research
WORKINGDIR?= $(PWD)/notebooks

.PHONY: test

all: | build run

setup:	;
build:
		$(DOCKERCMD) build -t $(TAG)
test:	;
clean: ;
run:
		$(DOCKERCMD) run -itv $(WORKINGDIR):/home/jovyan/work -p $(PORT):8888 $(TAG)