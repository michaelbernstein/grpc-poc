PROTO_DIR = ../panopticon
GO_DIR = ../
ELIXIR_CLIENT_DIR = elixir_client

PROTO_FILES = $(PROTO_DIR)/panopticon.proto
GO_OUT_DIR = $(GO_DIR)/panopticon
ELIXIR_OUT_DIR = $(ELIXIR_CLIENT_DIR)/lib

PROTOC = protoc
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc
PROTOC_GEN_ELIXIR = protoc-gen-elixir

GO_PKG_DIR = $(GO_DIR)
GO_PKG = panopticon

MIX = mix

.PHONY: all clean proto go_server elixir_client go_client

all: proto go_server elixir_client go_client

clean:
	rm -f $(GO_OUT_DIR)/*.pb.go
	rm -f $(ELIXIR_OUT_DIR)/*.ex
	rm pan_server
	rm pan_client

proto: $(PROTO_FILES)
	$(PROTOC) --go_out=$(GO_OUT_DIR) --go-grpc_out=$(GO_OUT_DIR) -I$(PROTO_DIR) $(PROTO_FILES)
	$(PROTOC) --elixir_out=plugins=grpc:$(ELIXIR_OUT_DIR) -I$(PROTO_DIR) $(PROTO_FILES)

go_server:
	go build -o pan_server server/main.go 

go_client:
	go build -o pan_client client/main.go


elixir_client:
	cd $(ELIXIR_CLIENT_DIR) && $(MIX) deps.get
