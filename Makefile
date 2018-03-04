BINARY_HTTP=bin
BINARY_BACK=bin
BINARY_CONF=configure

CURRENTDIR=$(shell pwd)

BINARY_DIR=$(CURRENTDIR)/bin

SRC_HTTP=./http
SRC_BACK=./background
SRC_CONF=./configure



PATH_HTTP=$(BINARY_DIR)/http/
PATH_BACK=$(BINARY_DIR)/back/
PATH_CONF=$(BINARY_DIR)/

PATH_CONFIG_HTTP=$(PATH_HTTP)/config

MKDIR_P=mkdir -p

GO=$(shell which go)
GOBUILD=$(GO) build
GOBUILDTAGS=$(GOBUILD) -tags netgo

all:
	$(MKDIR_P) $(PATH_CONFIG_HTTP) $(PATH_HTTP)
	$(GOBUILDTAGS) -o $(PATH_HTTP)/$(BINARY_HTTP)  $(SRC_HTTP)
	$(GOBUILDTAGS) -o $(PATH_BACK)/$(BINARY_BACK)  $(SRC_BACK)
	$(GOBUILDTAGS) -o $(PATH_CONF)/$(BINARY_CONF)  $(SRC_CONF)
