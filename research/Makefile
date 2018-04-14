# Python parameters
DOCKERCMD=docker
PORT?=8888
TAG?=research
WORKINGDIR?= $(PWD)/notebooks
ENV_FILE?=.env_file

.PHONY: test

all: | build run

setup:
		touch .env_file
build:
		$(DOCKERCMD) build -t $(TAG) ./
test:	;
clean: ;
run:
		$(DOCKERCMD) run -itv $(WORKINGDIR):/home/jovyan/work -p $(PORT):8888 --env-file $(ENV_FILE) $(TAG)
