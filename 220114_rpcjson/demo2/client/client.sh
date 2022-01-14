#!/bin/bash
curl localhost:8888/jsonrpc -X POST \
    --data '{"method":"HelloService.Length","params":["hello"],"id":0}'