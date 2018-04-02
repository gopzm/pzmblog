.PHONY: all
all: backend

.PHONY: backend
backend: protos
	go build backend/server.go

.PHONY: protos
protos:
	cd ./proto && make protos

.PHONY: clean_protos
clean_protos:
	cd ./proto && make clean

.PHONY: clean
clean: clean_protos
	rm -f server
