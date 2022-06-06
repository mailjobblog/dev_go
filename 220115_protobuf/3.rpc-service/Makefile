.PHONY:proto
proto: ## 生成proto的go代码
	# 由proto生成go代码
	protoc --go_out=. proto/*.proto
	# 由proto生成go的grpc代码
	protoc --go-grpc_out=. proto/*.proto

.PHONY:help
.DEFAULT_GOAL:=help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'