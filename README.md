# RPC Demos

The project shows the RPC demos in Golang including net/rpc, gRPC, etc.

## Contents

- netrpc1: A simplified RPC demo using net/rpc.
- netrpc2: An RPC demo improvement based on netrpc1, the features are
    - launch a goroutine to handle an RPC request;
    - encapsulate a client for HelloService to ensure the arguments are correct.
- netrpc3: An RPC demo using JsonRPC to call methods across languages.
