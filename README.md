# RPC Demos

The project shows the RPC demos in Golang including net/rpc, gRPC, etc.

## Contents

- netrpc1: A simplified RPC demo using net/rpc.
- netrpc2: An RPC demo improvement based on netrpc1, the features are
    - launch a goroutine to handle an RPC request;
    - encapsulate a client for HelloService to ensure the arguments are correct.
- netrpc3: An RPC demo using JsonRPC to call methods across languages.
- netrpcpb: Generate code automatically by customized plug-ins of protobuf.
- kvdb: A kv-pair storage based on RPC.
- rvsrpc: A reversed RPC.
- rpcctx: A RPC with context.
- grpcdm: A gRPC demo.

## Getting Started

### netrpcpb

pbgenerator is a customized protoc-gen-go, which contains the customized plug-ins defined in `pbplug/generator.go`. So the first thing is building the pbgenerator by the following commands.

```shell
cd netrpcpb/pbgenerator \
  && go build \
  && mv ./pbgenerator $PATH/protoc-gen-go-netrpc
```

Then, generate the code automatically.

```shell
cd ${PROJECT_ROOT} \
  && protoc --go-netrpc_out=plugins=netrpc:. netrpcpb/hello.proto
```

where `${PROJECT_ROOT}` is the path of the root of your project.
