#!/bin/bash
curl localhost:8888/jsonrpc -X POST \
    --data '{"method":"HelloService.Length","params":["test str 666"],"id":0}'