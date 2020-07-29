# Filecoin p2p oracle

This oracle allow you to get data from different lotus nodes on filecoin blockchain network

## Requirements
Require go version >=1.14 , so make sure your go version is okay.
WARNING! Building happen only when this project locates outside of GOPATH environment.

## Getting started

1. Clone this repository
2. Build binary files by `make build`
3. Run created binaries

## Example .env file 

1. SERVER_PORT=":8000"
2. LOG_LEVEL="debug"
3. NODE_MULTIADDRESS="/ip4/127.0.0.1/tcp/0"