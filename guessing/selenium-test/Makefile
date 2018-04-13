# Python parameters
PYTHONCMD=python3
PIP=pip3


.PHONY: test

all: | setup run

travis: | setup build test coverage

setup:
	$(PIP) install -r requirements.txt
build:	;
test:	;
coverage:	;
		

clean: ;
run:
		$(PYTHONCMD) test.py