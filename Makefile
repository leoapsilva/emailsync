ifndef VERBOSE
.SILENT:
endif

BINARY = emailsync
PRODUCTION_DIR=${BINARY}-build
CURRENT_DIR = $(shell pwd)
GOARCH = 386
GOOS = windows
LDFLAGS = -ldflags "-s -w -X main.build=DEV"
VERSION?=?
ENV_FILE = .env

all: build deploy

install: 
	go mod tidy

build:
	go build ${LDFLAGS} -o ${BINARY}-${GOARCH}-${GOARCH}.exe

deploy:
	mkdir -p ./${PRODUCTION_DIR} && \
	cp ./.env ./${PRODUCTION_DIR} && \
	cp ./${BINARY}-${GOARCH}-${GOARCH}.exe ./${PRODUCTION_DIR} && \
	cd ./doc && \
	make

