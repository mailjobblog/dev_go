# Go语言使用protobuf

## 前言

protobuf 即 Protocol Buffers，是一种轻便高效的结构化数据存储格式，与语言、平台无关，可扩展可序列化。  
protobuf 性能和效率大幅度优于 JSON、XML 等其他的结构化数据格式。  
protobuf 是以二进制方式存储的，占用空间小，但也带来了可读性差的缺点。protobuf 在通信协议和数据存储等领域应用广泛。  
  
Protobuf 在 `.proto` 定义需要处理的结构化数据，可以通过 `protoc` 工具，将 `.proto` 文件转换为 C、C++、Golang、Java、Python 等多种语言的代码，兼容性好，易于使用。

## 参考文献

> protobuf3 官方文档：https://link.jianshu.com/?t=https://developers.google.com/protocol-buffers/docs/proto3  
> Protocol Buffer 编码：https://developers.google.com/protocol-buffers/docs/encoding?hl=zh-cn#packed  
> proto service grpc 生成插件：https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md  
> 本文代码下载：https://github.com/mailjobblog/dev_go/tree/master/220115_protobuf  

## 安装

**安装 protoc**   
从 [Protobuf Releases](https://github.com/protocolbuffers/protobuf/releases) 下载最先版本的发布包安装。

```bash
brew intall protoc
```

**安装 protoc-gen-go**
我们需要在 Golang 中使用 protobuf，还需要安装 protoc-gen-go，这个工具用来将 .proto 文件转换为 Golang 代码。
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
Tips：  
这儿有个小小的坑，`github.com/golang/protobuf/protoc-gen-go` 和 `google.golang.org/protobuf/cmd/protoc-gen-go`是不同的。  
区别在于前者是旧版本，后者是google接管后的新版本，他们之间的API是不同的，也就是说用于生成的命令，以及生成的文件都是不一样的。  

**检查是否安装成功**

```bash
$ protoc --version
libprotoc 3.19.3

$ protoc-gen-go --version
protoc-gen-go v1.27.1
```


## protobuf生成代码

### 快速上手

接下来，我们创建一个非常简单的示例，`student.proto`
```protobuf
syntax = "proto3";
package main;

// this is a comment
message Student {
  string name = 1;
  bool male = 2;
  repeated int32 scores = 3;
}
```
在当前目录下执行代码生成命令
```bash
$ protoc --go_out=. *.proto

$ ls
student.pb.go  student.proto
```
执行此生成命令是将该目录下的所有的 .proto 文件转换为 Go 代码，我们可以看到该目录下多出了一个 Go 文件 `student.pb.go`。这个文件内部定义了一个结构体 Student，以及相关的方法：
```go
type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Male   bool    `protobuf:"varint,2,opt,name=male,proto3" json:"male,omitempty"`
	Scores []int32 `protobuf:"varint,3,rep,packed,name=scores,proto3" json:"scores,omitempty"`
}
```
得到生成的 `student.pb.go` 文件后，可以在项目代码中直接使用了。  
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

### 枚举(Enumerations)
枚举类型适用于提供一组预定义的值，选择其中一个。例如我们将性别(gender)定义为枚举类型。
```protobuf
message StudentEnum {
  string name = 1;
  enum Gender {
    FEMALE = 0;
    MALE = 1;
  }
  Gender gender = 2;
  repeated int32 scores = 3;
}
```
生成的Go代码主要信息如下：
```go
type StudentEnum_Gender int32

const (
	StudentEnum_FEMALE StudentEnum_Gender = 0
	StudentEnum_MALE   StudentEnum_Gender = 1
)

type StudentEnum struct {
    state         protoimpl.MessageState
    sizeCache     protoimpl.SizeCache
    unknownFields protoimpl.UnknownFields
    
    Name   string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
    Gender StudentEnum_Gender `protobuf:"varint,2,opt,name=gender,proto3,enum=main.StudentEnum_Gender" json:"gender,omitempty"`
    Scores []int32            `protobuf:"varint,3,rep,packed,name=scores,proto3" json:"scores,omitempty"`
}
```
枚举类型的第一个选项的标识符`必须是0`，这也是枚举类型的默认值。  

### 别名（Alias）
允许为不同的枚举值赋予相同的标识符，称之为别名，需要打开allow_alias选项。
```protobuf
message StudentAlias {
  enum Status {
    option allow_alias = true;
    UNKOWN = 0;
    STARTED = 1;
    RUNNING = 1;
  }
}
```
生成的Go语言主要代码如下：
```go
type StudentAlias_Status int32

const (
	StudentAlias_UNKOWN  StudentAlias_Status = 0
	StudentAlias_STARTED StudentAlias_Status = 1
	StudentAlias_RUNNING StudentAlias_Status = 1
)
```

### 使用其他消息类型
Result是另一个消息类型，在 SearchReponse 作为一个消息字段类型使用。
```protobuf
message StudentResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```
嵌套写也是支持的：
```protobuf
message Student2Response {
  message Result2 {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result2 results2 = 1;
}
```
如果定义在其他文件中，可以导入其他消息类型来使用：
```protobuf
import "myproject/other_protos.proto";
```

### 任意类型(Any)

在使用 GRPC 时，常规的操作是将 message 定义好后进行数据传输，但总会遇到某些数据结构进行组合的操作，采用默认的定义 message 方式，造成代码量的激增。  
为了解决这个问题 protobuf 提供类型 any 解决 GRPC 中泛型的处理方式
```protobuf
message StudentAny {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```
生成的Go语言主要代码如下：
```go
type StudentAny struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Details []*anypb.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
}
```
### oneof
如果你的消息中有很多可选字段， 并且同时至多一个字段会被设置， 你可以加强这个行为，使用oneof特性节省内存。  
Oneof字段就像可选字段， 除了它们会共享内存， 至多一个字段会被设置。 设置其中一个字段会清除其它字段。
```protobuf
message StudentOneOf {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
```
**oneof的特性**
- 设置oneof会自动清楚其它oneof字段的值. 所以设置多次后，只有最后一次设置的字段有值
- 如果解析器遇到同一个oneof中有多个成员，只有最后一个会被解析成消息
- oneof不支持repeated

### map
```protobuf
message StudentMap {
  map<string, int32> points = 1;
}
```

### 定义服务(Services)
如果消息类型是用来远程通信的(Remote Procedure Call, RPC)，可以在 .proto 文件中定义 RPC 服务接口。  
例如我们定义了一个名为 SearchService 的 RPC 服务，提供了 Search 接口，入参是 SearchRequest 类型，返回类型是 SearchResponse
```protobuf
service SearchService {
  rpc Search (SearchRequest) returns (SearchResponse);
}
```
官方仓库也提供了一个 [插件列表](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md)，帮助开发基于 Protocol Buffer 的 RPC 服务。  

**生成go代码和grpc代码：**
```bash
# 由proto生成go代码
protoc --go_out=. *.proto

# 由proto生成go的grpc代码
protoc --go-grpc_out=. *.proto
```

### protoc 命令参数
```bash
protoc --proto_path=IMPORT_PATH --<lang>_out=DST_DIR path/to/file.proto
```
**--proto_path=IMPORT_PATH：** 可以在 .proto 文件中 import 其他的 .proto 文件，proto_path 即用来指定其他 .proto 文件的查找目录。如果没有引入其他的 .proto 文件，该参数可以省略。
**--<lang>_out=DST_DIR：** 指定生成代码的目标文件夹，例如 `–go_out=.` 即生成 GO 代码在当前文件夹，另外支持 cpp/java/python/ruby/objc/csharp/php 等语言

### 推荐风格

#### 文件(Files)
- 文件名使用小写下划线的命名风格，例如 lower_snake_case.proto
- 每行不超过 80 字符
- 使用 2 个空格缩进

#### 包(Packages)
- 包名应该和目录结构对应，例如文件在my/package/目录下，包名应为 my.package

#### 消息和字段(Messages & Fields)
- 消息名使用首字母大写驼峰风格(CamelCase)，例如message StudentRequest { ... }
- 字段名使用小写下划线的风格，例如 string status_code = 1
- 枚举类型，枚举名使用首字母大写驼峰风格，例如 enum FooBar，枚举值使用全大写下划线隔开的风格(CAPITALS_WITH_UNDERSCORES )，例如 FOO_DEFAULT=1

#### 服务(Services)
- RPC 服务名和方法名，均使用首字母大写驼峰风格，例如service FooService{ rpc GetSomething() }


## protobuf文件规范

### syntax
protobuf 有2个版本，默认版本是 proto2，如果需要 proto3，则需要在非空非注释第一行使用 `syntax = "proto3"` 标明版本。

### package
package，即包名声明符是可选的，用来防止不同的消息类型有命名冲突。

### option go_package
```protobuf
option go_package="./proto/pb;pb";
```
这部分的内容是关于最后生成的go文件是处在哪个目录哪个包中，`./proto/pb` 代表在当前目录生成，`pb` 代表了生成的go文件的包名是 pb。

### message
消息类型 使用 message 关键字定义，Student 是类型名，name, male, scores 是该类型的 3 个字段，类型分别为 string, bool 和 []int32。字段可以是标量类型，也可以是合成类型。
相当于Go语言中的 struct 结构体。一个 .proto 文件中可以写多个消息类型，即对应多个结构体(struct)。  

### 修饰符
每个字段的修饰符默认是 `singular`，一般省略不写，`repeated` 表示字段可重复，即用来表示 Go 语言中的切片类型。

### 标识符
每个字符 = 后面的数字称为标识符，每个字段都需要提供一个唯一的标识符。标识符用来在消息的二进制格式中识别各个字段，一旦使用就不能够再改变，标识符的取值范围为 [1, 2^29 - 1] 。

### 文件注释
.proto 文件可以写注释，单行注释 //，多行注释 /* ... */

### 标量类型(Scalar)

| .proto Type | Notes                                                        | C++ Type | Java Type  | Python Type[2] | Go Type | Ruby Type                      | C# Type    | PHP Type       |
| :---------- | :----------------------------------------------------------- | :------- | :--------- | :------------- | :------ | :----------------------------- | :--------- | :------------- |
| double      |                                                              | double   | double     | float          | float64 | Float                          | double     | float          |
| float       |                                                              | float    | float      | float          | float32 | Float                          | float      | float          |
| int32       | 使用变长编码，对于负值的效率很低，如果你的域有可能有负值，请使用sint64替代 | int32    | int        | int            | int32   | Fixnum 或者 Bignum（根据需要） | int        | integer        |
| uint32      | 使用变长编码                                                 | uint32   | int        | int/long       | uint32  | Fixnum 或者 Bignum（根据需要） | uint       | integer        |
| uint64      | 使用变长编码                                                 | uint64   | long       | int/long       | uint64  | Bignum                         | ulong      | integer/string |
| sint32      | 使用变长编码，这些编码在负值时比int32高效的多                | int32    | int        | int            | int32   | Fixnum 或者 Bignum（根据需要） | int        | integer        |
| sint64      | 使用变长编码，有符号的整型值。编码时比通常的int64高效。      | int64    | long       | int/long       | int64   | Bignum                         | long       | integer/string |
| fixed32     | 总是4个字节，如果数值总是比总是比228大的话，这个类型会比uint32高效。 | uint32   | int        | int            | uint32  | Fixnum 或者 Bignum（根据需要） | uint       | integer        |
| fixed64     | 总是8个字节，如果数值总是比总是比256大的话，这个类型会比uint64高效。 | uint64   | long       | int/long       | uint64  | Bignum                         | ulong      | integer/string |
| sfixed32    | 总是4个字节                                                  | int32    | int        | int            | int32   | Fixnum 或者 Bignum（根据需要） | int        | integer        |
| sfixed64    | 总是8个字节                                                  | int64    | long       | int/long       | int64   | Bignum                         | long       | integer/string |
| bool        |                                                              | bool     | boolean    | bool           | bool    | TrueClass/FalseClass           | bool       | boolean        |
| string      | 一个字符串必须是UTF-8编码或者7-bit ASCII编码的文本。         | string   | String     | str/unicode    | string  | String (UTF-8)                 | string     | string         |
| bytes       | 可能包含任意顺序的字节数据。                                 | string   | ByteString | str            | []byte  | String (ASCII-8BIT)            | ByteString | string         |

**标量类型如果没有被赋值，则不会被序列化，解析时，会赋予默认值**
- strings：空字符串
- bytes：空序列
- bools：false
- 数值类型：0


## 常见问题

### go_package报错

```text
Please specify either:
• a "go_package" option in the .proto source file, or
• a "M" argument on the command line.
```
在go的1.14版本以后，proto文件中不添加go_package 会报错。  
解决方法： option go_package = "./"  
或者填写自己的包路径也行如option go_package = "http://github.com/package/name"  

### 安装protoc-gen-go报错
```text
can't load package: package google.golang.org/protobuf/cmd/protoc-gen-go: cannot find package "google.golang.org/protobuf/cmd/protoc-gen-go" in any of:
        C:\Go\src\google.golang.org\protobuf\cmd\protoc-gen-go (from $GOROOT)
        C:\Users\peikai\go\src\google.golang.org\protobuf\cmd\protoc-gen-go (from $GOPATH)
```
解决方法：  
先 `go get google.golang.org/protobuf/cmd/protoc-gen-go` 然后再 `install` 安装。

