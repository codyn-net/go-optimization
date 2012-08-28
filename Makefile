PKGCONFIG = $(shell which pkg-config)

ifeq ($(PKGCONFIG),)
$(error Please install pkg-config...)
endif

PROTODIR = $(shell $(PKGCONFIG) --variable=protodir liboptimization-2.0 2>/dev/null)

ifeq ($(PROTODIR),)
$(error Could not find liboptimization-2.0 development files...)
endif

PROTO = $(wildcard $(PROTODIR)/optimization/messages/*.proto)
BASE = $(GOPATH)/src/optimization/messages

PROTOFILES = $(foreach i,$(PROTO),$(notdir $(i)))
GOFILES = $(PROTOFILES:.proto=.pb.go)

BUILTFILES = $(foreach i,$(GOFILES),$(BASE)/$(i:.pb.go=.pb)/$(i))

all:
	@echo "Use make install to generate and install the proto files. The files will be installed in $(BASE)."

install: $(BUILTFILES)

uninstall:
		rm -f $(BUILTSOURCES)

$(BUILTFILES): $(BASE)/%.pb.go:
	@echo "[GEN] $(notdir $@)"; \
	protoc --go_out=$(TMPDIR) -I$(PROTODIR) $(PROTODIR)/optimization/messages/$(basename $(basename $(notdir $@))).proto && \
	mkdir -p $(dir $@); \
	mv $(TMPDIR)/optimization/messages/$(notdir $@) $@

.PHONY : all install uninstall