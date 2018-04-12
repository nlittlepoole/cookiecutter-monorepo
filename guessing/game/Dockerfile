FROM golang:alpine

RUN apk add --update make
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ARG app_env
ENV APP_ENV $app_env
ENV AMPUP_PORT=80

COPY . /go/src/github.com/nlittlepoole/cookiecutter-monorepo/guessing/game
WORKDIR /go/src/github.com/nlittlepoole/cookiecutter-monorepo/guessing/game

ENV PATH="$PATH:$GOROOT/bin:$GOPATH/bin"

EXPOSE 80