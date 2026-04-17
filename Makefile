UNAME_S  := $(shell uname -s)
MACOS_VER := $(shell sw_vers -productVersion 2>/dev/null | cut -d. -f1)

# SDL2 pkg-config path (Homebrew on macOS needs the OS-version subdir for zlib)
ifeq ($(UNAME_S),Darwin)
  PKG_CONFIG_PATH := /opt/homebrew/lib/pkgconfig:/opt/homebrew/Homebrew/Library/Homebrew/os/mac/pkgconfig/$(MACOS_VER)
  export PKG_CONFIG_PATH
endif

GOIMPORTS := $(shell which goimports 2>/dev/null || echo $(GOPATH)/bin/goimports)
ROCQ      := $(shell eval $$(opam env --switch=rocq900 2>/dev/null) && which rocq 2>/dev/null || which rocq 2>/dev/null || echo rocq)

THEORY    := theories/RocqmanGo.v
GENERATED := main.go

.PHONY: all extract build run clean

all: build

# Step 1: compile the Rocq theory → generates main.go
extract: $(GENERATED)

$(GENERATED): $(THEORY) theories/GoSDL2.v theories/dune dune-project
	eval $$(opam env --switch=rocq900) && \
	  OCAMLFIND_CONF=~/.opam/rocq900/lib/findlib.conf \
	  dune build theories/RocqmanGo.vo
	cp _build/default/main.go .
	$(GOIMPORTS) -w $(GENERATED)

# Step 2: compile the Go package → produces ./rocqman-go
build: $(GENERATED)
	go build -o rocqman-go .

run: build
	./rocqman-go

clean:
	dune clean
	rm -f $(GENERATED) rocqman-go
