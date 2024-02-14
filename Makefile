#!/usr/bin/make -f

VERSION ?= $(shell echo $(shell git describe --tags `git rev-list --tags="v*" --max-count=1`) | sed 's/^v//')
TMVERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
GALACTICA_BINARY = galacticad
BUILDDIR ?= $(CURDIR)/build


# RocksDB is a native dependency, so we don't assume the library is installed.
# Instead, it must be explicitly enabled and we warn when it is not.
ENABLE_ROCKSDB ?= false

# Default target executed when no arguments are given to make.
default_target: all

.PHONY: default_target build

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/Galactica-corp/galactica/version.Name=galactica \
		  -X github.com/Galactica-corp/galactica/version.AppName=$(GALACTICA_BINARY) \
		  -X github.com/Galactica-corp/galactica/version.Version=$(VERSION) \
		  -X github.com/Galactica-corp/galactica/version.Commit=$(COMMIT) \
			-X "github.com/Galactica-corp/galactica/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/evmos/ethermint/version.TMCoreSemVer=$(TMVERSION)

ifeq ($(ENABLE_ROCKSDB),true)
  BUILD_TAGS += rocksdb_build
  test_tags += rocksdb_build
endif

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += gcc
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  ifneq ($(ENABLE_ROCKSDB),true)
    $(error Cannot use RocksDB backend unless ENABLE_ROCKSDB=true)
  endif
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
endif

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# check if no optimization option is passed
# used for remote debugging
ifneq (,$(findstring nooptimization,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -gcflags "all=-N -l"
endif

###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/
build-linux:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

build-alpine:
	@echo "Building galacticad for alpine..."
	@GOOS=linux CGO_ENABLED=1 go build -ldflags="-w -s" -o ./build/galacticad ./cmd/galacticad

localnet-build:
	@make build
	@echo "Build docker image for localnet..."
	@./localnet/build-docker.sh

localnet-init:
	@echo "Init configs for localnet..."
	@./localnet/init-configs-localnet.sh

localnet-start:
	@echo "Start localnet..."
	@./localnet/run-validators.sh
