UNAME := $(shell uname)
TEST := $(shell pip)


build:
ifeq ($(TEST),)
ifeq ($(UNAME), Darwin)
	brew install python
endif
ifeq ($(UNAME), Linux)
	sudo apt-get install python-pip python-dev build-essential 
	sudo pip install --upgrade pip 
endif
endif
	pip install pylint
	pip install cpplint

install:
	go install .