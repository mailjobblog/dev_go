#!/bin/bash

# 由proto生成go代码
protoc --go_out=. *.proto

# 由proto生成go的grpc代码
protoc --go-grpc_out=. *.proto