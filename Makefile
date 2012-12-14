#PROTO = $(wildcard $(PROTODIR)/optimization/messages/*.proto)

#PROTOFILES = $(foreach i,$(PROTO),$(notdir $(i)))
#GOFILES = $(PROTOFILES:.proto=.pb.go)

#BUILTFILES = $(foreach i,$(GOFILES),messages/$(i:.pb.go=.pb)/$(i))

# All the messages
PROTO_MESSAGES = task command discovery monitor

# The temporary directory to fetch the proto files in
FETCH_DIR = .fetch-proto

# The fetched .proto files for each message
PROTO_FILES = $(foreach i,$(PROTO_MESSAGES),$(FETCH_DIR)/optimization/messages/$(i).proto)

# The go version of the proto files in messages/
PROTO_GO_FILES = $(foreach i,$(PROTO_MESSAGES),messages/$(i).pb/$(i).pb.go)

# The url where to fetch the proto files from
FETCH_BASE = https://ponyo.epfl.ch/cgit/index.cgi/optimization/liboptimization.git/plain/optimization/messages

all:
	@echo "Use the update-proto target to update the go protobuf files."

$(PROTO_FILES):
	@mkdir -p $(dir $@) || exit 1;				\
	echo "Fetching $@";					\
	curl -s -o "$@" "$(FETCH_BASE)/$(notdir $@)";

update-proto: $(PROTO_GO_FILES)
	@rm -rf $(FETCH_DIR);							\
	for i in $(PROTO_MESSAGES); do						\
		echo "[FIX] $$i.pb.go";						\
		go run fixproto/fixproto.go '' messages/$$i.pb/$$i.pb.go;	\
	done

clean-proto:
	@rm -rf messages/* $(FETCH_DIR)

$(PROTO_GO_FILES): $(PROTO_FILES)
	@echo "[GEN] $(notdir $@)"; \
	pname=$(basename $(basename $(notdir $@)));				\
	protoc --go_out=$(FETCH_DIR) -I$(FETCH_DIR) $(FETCH_DIR)/optimization/messages/$$pname.proto; \
	mkdir -p messages/$$pname.pb;			\
	mv $(FETCH_DIR)/optimization/messages/$$pname.pb.go messages/$$pname.pb/

.PHONY : fetch-proto clean-proto update-proto all
