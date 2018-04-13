[![forthebadge](https://forthebadge.com/images/badges/you-didnt-ask-for-this.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/fo-shizzle.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/certified-snoop-lion.svg)](https://forthebadge.com)
# cookiecutter-monorepo
[![Build Status](https://travis-ci.org/nlittlepoole/cookiecutter-monorepo.svg?branch=master)](https://travis-ci.org/nlittlepoole/cookiecutter-monorepo)
[![codecov](https://codecov.io/gh/nlittlepoole/cookiecutter-monorepo/branch/master/graph/badge.svg)](https://codecov.io/gh/nlittlepoole/cookiecutter-monorepo)

## Overview
This reposiotry is a template, meant to be the foundation of a polyglot monrepo. The core pillars of this setup are [Makefiles](https://gist.github.com/isaacs/62a2d1825d04437c6f08), [Docker](https://www.docker.com/), [TravisCI](https://travis-ci.org/), [CodeCov](https://codecov.io). This structure is built around the idea of (Micro)Service oriented architecture. An example service is provided in the `guessing` directory. It implements a naive version of the [21 Questions](http://web.stanford.edu/class/archive/cs/cs106x/cs106x.1174/assn/twentyOneQuestions.html) game and provides a simple html front end to access the game. 

## Service vs Micro-service
The debate about the difference between Microservice Architecture and Service Oriented Archietecture is still [ongoing](https://stackoverflow.com/questions/25501098/difference-between-microservices-architecture-and-soa). I'm going to take a page from [Martin Fowler](https://youtu.be/2yko4TbC8cI?t=15m53s), and define micro services as subsets of services. Services and Micro-services should follow the

- Boundaries are explicit
- Services are autonomous
- Services share schema and contract, not class
- Service compatibility is based on policy

![](https://i1.wp.com/swaggerhub.com/wp-content/uploads/2017/04/MicroservicesIntro.png?resize=1076%2C517)

The above demonstrates how I conceptualize services and micro-services in a mock ridesharing app. The individual rectangles such as `Billing`, `Notification`, etc are micro-services. A Service is an abstraction on top of a combination of micro-services. The green diamond around them represents a service, in this case a `User Service`. There may be another Service, like a `Dispatch Service`, that has its own micro services such as `Pricing`, `Routing`, etc. Boundaries between the `User Service` and `Dispatch Service` are explicit. The micro-services in one can't arbitrarily call the micro-services in another. Only a subset of endpoints will be exposed, usually through an API gateway, to each other. 

## How to Add a Service to the Repo

### Add a Service Directory
The following is the directory structure expected of any service in the repository

```
service
│   README.md
│   Makefile
│   docker-compose.yml
│
└───micro-service-1
│   │   README.md
│   │   Makefile
│   │   Dockerfile
│   │
│   └───doc
│       │   some-file.md
│       │   some-file.html
│       │   ...
│   
└───micro-service-2
│   │   README.md
│   │   Makefile
│   │   Dockerfile
│   │
│   └───doc
│       │   some-file.png
│       │   some-file.txt
│       │   ...
│  
│   └───src
│       │   SomeClass.go
│       │   AnotherClass.go
│       │   ...
```
Micro services are children of other services. Each is expected to have a doc folder, README, Makefile and Dockerfile. Otherwise developers can add any other directories needed in convention with what is idiomatic for the language of the service. 

### Makefile

```cmake

all: ;

travis: | setup build test coverage

setup:	;
build:	;
test:	;
coverage:	;
		

clean: ;
run:   ;
```

The above is an example of the structure of the Makefile that is expected in every service and micro-service directory. Travis has been configured to use these make files.

- `setup`: Any dependencies the base system needs to install, for example pip installing requirements.txt.
- `build`: Compiling the source code. For interpreted languages like Ruby or Python this isn't necessary.
- `test`: Run all the associated tests. 
- `clean`: Clean the local machine to enable a fresh build
- `run`: Run the (micro)service
- `coverage`: Output a CodeCov compatible coverage file to the `/shared/` directory. Used by travis
- `travis`: this should almost always be as I've setup above. This is the command travis will actually run to test the code. 
- `all`: Whatever you think is useful here. Not required.

### Add Each Micro Service to .travis.yml
Add each micro service to the environment section of the Travis Config file. Specifically add each to the env list. Travis' `Build Matrix` feature is utilized to run tests on each service provided.

So for example if I wanted to add a Service, `face-detection`, with 2 micro services; `api`, `neural-network`. I would make my Travis file look like the following

```yaml
env:
	- MICRO_SERVICE=guessing/game
	- MICRO_SERVICE=face-detection/api
	- MICRO_SERVICE=face-detection/neural-network
```
