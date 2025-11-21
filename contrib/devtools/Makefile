###
# Find OS and Go environment
# GO contains the Go binary
# FS contains the OS file separator
###
ifeq ($(OS),Windows_NT)
  GO := $(shell where go.exe 2> NUL)
  FS := "\\"
else
  GO := $(shell command -v go 2> /dev/null)
  FS := "/"
endif

ifeq ($(GO),)
  $(error could not find go. Is it in PATH? $(GO))
endif


###############################################################################
###                                 Tools                                   ###
###############################################################################

PREFIX ?= /usr/local
BIN ?= $(PREFIX)/bin
UNAME_S ?= $(shell uname -s)
UNAME_M ?= $(shell uname -m)

GOPATH ?= $(shell $(GO) env GOPATH)
GITHUBDIR := $(GOPATH)$(FS)src$(FS)github.com

BUF_VERSION ?= 0.11.0

TOOLS_DESTDIR  ?= $(GOPATH)/bin
STATIK         = $(TOOLS_DESTDIR)/statik

tools: tools-stamp
tools-stamp: statik
	# Create dummy file to satisfy dependency and avoid
	# rebuilding when this Makefile target is hit twice
	# in a row.
	touch $@

# Install the runsim binary with a temporary workaround of entering an outside
# directory as the "go get" command ignores the -mod option and will polute the
# go.{mod, sum} files.
#
# ref: https://github.com/golang/go/issues/30515

# In genera statik allows you to embed a directory of static files into your Go binary
#to be later served from an http.FileSystem. Here it is used as a filesystem for swaggerDocAPI server.
statik: $(STATIK)
$(STATIK):
	@echo "Installing statik..."
	@(cd /tmp && go install github.com/rakyll/statik@v0.1.6)

# Install the runsim binary with a temporary workaround of entering an outside
# directory as the "go get" command ignores the -mod option and will polute the
# go.{mod, sum} files.
#
# ref: https://github.com/golang/go/issues/30515
# runsim is used for running up simulation tests on different threads.

tools-clean:
	rm -f $(STATIK) $(GOLANGCI_LINT) $(RUNSIM)
	rm -f tools-stamp

.PHONY: tools-clean statik
