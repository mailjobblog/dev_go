## protobuf 快速上手

接下来，我们创建一个非常简单的示例，`user.proto`

```protobuf
syntax = "proto3";

package api;

// protoc-gen-go 版本大于1.4.0, proto文件需要加上go_package,否则无法生成
option go_package="proto/api";

message User {
  string name = 1;
  bool status = 2;
  repeated string hobby = 3;
}
```

在当前目录下执行代码生成命令

```bash
$ protoc --go_out=. proto/user.proto
```

执行此生成命令是将该目录下的所有的 .proto 文件转换为 Go 代码，我们可以看到该目录下多出了一个 Go 文件 `user.pb.go`。这个文件内部定义了一个结构体 User，以及相关的方法：

```go
type User struct {
    state         protoimpl.MessageState
    sizeCache     protoimpl.SizeCache
    unknownFields protoimpl.UnknownFields
    
    Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
    Status bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
    Hobby  []string `protobuf:"bytes,3,rep,name=hobby,proto3" json:"hobby,omitempty"`
}
```
得到生成的 `user.pb.go` 文件后，可以在项目代码中直接使用了。  
以下是一个例子，即证明被序列化的和反序列化后的实例，包含相同的数据。
```go
func main() {
	test := &Student{
		Name: "geektutu",
		Male:  true,
		Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}
}
```
