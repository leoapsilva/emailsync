ifndef VERBOSE
.SILENT:
endif

BINARY = emailsync
PRODUCTION_DIR=${BINARY}-build
CURRENT_DIR = $(shell pwd)

all: build deploy

install: 
	npm install swagger-cli

build:
	swagger-cli bundle ${BINARY}.yml -o ${BINARY}.json;

deploy:
	mkdir -p ../${PRODUCTION_DIR} && \
	mkdir -p ../${PRODUCTION_DIR}/doc && \
	cp -R ./dist ../${PRODUCTION_DIR}/doc && \
	cp ./swagger-initializer.js ../${PRODUCTION_DIR}/doc && \
	cp ./index.html ../${PRODUCTION_DIR}/doc && \
	cp ./${BINARY}.json ../${PRODUCTION_DIR}/doc

clean: 
	rm -rf ${BINARY}.json