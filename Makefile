PROTODIR ?= __noproto

ifeq ($(PROTODIR),__noproto)

PKGCONFIG = $(shell which pkg-config)

ifeq ($(PKGCONFIG),)
$(error Please install pkg-config (or specify PROTODIR manually))
endif

PROTODIR = $(shell $(PKGCONFIG) --variable=protodir liboptimization-2.0 2>/dev/null)

ifeq ($(PROTODIR),)
$(error Please install the liboptimization-2.0 development files or set PROTODIR)
endif

NOFETCH = 1

endif

PROTO = $(wildcard $(PROTODIR)/optimization/messages/*.proto)
BASE = $(GOPATH)/src/optimization/messages

PROTOFILES = $(foreach i,$(PROTO),$(notdir $(i)))
GOFILES = $(PROTOFILES:.proto=.pb.go)

BUILTFILES = $(foreach i,$(GOFILES),$(BASE)/$(i:.pb.go=.pb)/$(i))

all:
	@echo "Use 'make install' to generate and install the proto files. The generated files will be installed in $(BASE)."; \
	echo "Found proto files: $(sort $(PROTOFILES))"

FETCH_BASE = https://ponyo.epfl.ch/cgit/index.cgi/optimization/liboptimization.git/plain/

fetch-proto:
	@if ! test -z "$(NOFETCH)"; then					\
		echo "Can't fetch in the system protodir, use make fetch-proto PROTODIR=";	\
		exit 1;								\
	fi;									\
	mkdir -p "$(PROTODIR)/optimization/messages" || exit 1;			\
	for i in task command discovery monitor; do 				\
		echo "Fetching $$i.proto to $(PROTODIR)/optimization/messages";	\
		curl -s -o "$(PROTODIR)/optimization/messages/$$i.proto"	\
		        "$(FETCH_BASE)/optimization/messages/$$i.proto";	\
	done

install: $(BUILTFILES)

uninstall:
		rm -f $(BUILTSOURCES)

$(BUILTFILES): $(BASE)/%.pb.go:
	@echo "[GEN] $(notdir $@)"; \
	pname=$(basename $(basename $(notdir $@)));				\
	protoc --go_out=$(GOPATH)/src -I$(PROTODIR) $(PROTODIR)/optimization/messages/$$pname.proto; \
	mkdir -p $(GOPATH)/src/optimization/messages/$$pname.pb;		\
	mv $(GOPATH)/src/optimization/messages/$$pname.pb.go $(GOPATH)/src/optimization/messages/$$pname.pb/

.PHONY : all install uninstall fetch-proto
