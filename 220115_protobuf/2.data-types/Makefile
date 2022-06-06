.PHONY:proto
proto: ## 生成proto的go代码
	# 枚举
	protoc --go_out=. proto/user_enum.proto
	# 别名（Alias）
	protoc --go_out=. proto/user_alias.proto
	# 嵌套
	protoc --go_out=. proto/user_rep.proto
	# any
	protoc --go_out=. proto/user_any.proto
	# oneof
	protoc --go_out=. proto/user_oneof.proto
	# map
	protoc --go_out=. proto/user_map.proto

.PHONY:help
.DEFAULT_GOAL:=help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'