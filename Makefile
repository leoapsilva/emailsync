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
	go build ${LDFLAGS} -o ${BINARY}-${GOOS}-${GOARCH}.exe

deploy:
	mkdir -p ./${PRODUCTION_DIR} && \
	cp ./.env ./${PRODUCTION_DIR} && \
	cp ./${BINARY}-${GOOS}-${GOARCH}.exe ./${PRODUCTION_DIR} && \
	cd ./doc && \
	make && \
	cd -

clean:
	rm -rf ./${PRODUCTION_DIR} && \
	rm -rf ${BINARY}-${GOOS}-${GOARCH}.exe && \
	cd ./doc && \
	make clean