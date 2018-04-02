define list_proto_names
	find ./ -name *.proto -execdir basename {} .proto ';'
endef

.PHONY: all
all: protos

.PHONY: protos
protos:
	# Make necessary dir
	$(call list_proto_names)| xargs  mkdir -p;
	# Generate go files
	$(call list_proto_names)| xargs -I % protoc --go_out=./% %.proto;

.PHONY: clean
clean:
	$(call list_proto_names) | xargs rm -rf;

