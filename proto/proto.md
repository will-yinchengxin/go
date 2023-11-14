# [protobuf](https://protobuf.dev/programming-guides/proto3/)

## 1.定义消息类型
### 1.1 指定字段类型


#### 1.1.1 [常量消息类型](https://protobuf.dev/programming-guides/proto3/#fieldreserved)
```protobuf
// 添加注释: 可以使用 `//` 或者 `/* ... */`
// 文件的第一行指定您正在使用proto3语法：如果您不这样做，协议缓冲区编译器将假设您正在使用 proto2。这必须是文件的第一个非空、非注释行。
syntax = "proto3";

// 所有字段都是标量类型
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 results_per_page = 3;
}

// 还可以为字段指定复合类型
message SearchAllRequest {
  SearchRequest search = 1;
  int64 time = 2;
}

// 还可以为字段指定枚举类型
enum Corpus {
  CORPUS_UNSPECIFIED = 0;
  CORPUS_UNIVERSAL = 1;
  CORPUS_WEB = 2;
  CORPUS_IMAGES = 3;
  CORPUS_LOCAL = 4;
  CORPUS_NEWS = 5;
  CORPUS_PRODUCTS = 6;
  CORPUS_VIDEO = 7;
}

message FindResult {
  string query = 1;
  int32 page_number = 2;
  int32 results_per_page = 3;
  Corpus corpus = 4;
}
````
#### 1.1.2 特殊的消息类型

- **optional：optional字段处于两种可能状态之一**：
- - 该字段已设置，并且包含显式设置或从线路解析的值 (调用者传递的值)。
- - 该字段未设置，将返回默认值 (即该类型对应的零值)。
- - -   对于字符串，默认值为空字符串。
- - -   对于字节，默认值为空字节。
- - -   对于布尔值，默认值为 false。
- - -   对于数字类型，默认值为零。
- - -   对于枚举，默认值是第一个定义的枚举值，该值必须为 0。
- - -   对于消息字段，未设置该字段。它的确切值取决于语言。有关详细信息，请参阅 生成的代码指南。
- - -   重复字段的默认值为空（通常是相应语言的空列表）。

```protobuf
syntax = "proto3";

message Person {
  string name = 1;
  int32 age = 2;
  optional string email = 3;
}
````
```go
package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"example.com/person"
)

func main() {
	p := &person.Person{
		Name: "John",
		Age:  30,
	}

	// 设置可选字段
	p.Email = "john@example.com"

	// 序列化消息
	data, _ := proto.Marshal(p)

	// 反序列化消息
	p2 := &person.Person{}
	_ = proto.Unmarshal(data, p2)

	fmt.Println(p2.Name)  // 输出: John
	fmt.Println(p2.Age)   // 输出: 30
	fmt.Println(p2.Email) // 输出: john@example.com
}
````

- **repeated：此字段类型可以在格式正确的消息中重复零次或多次。将保留重复值的顺序。**
```protobuf
syntax = "proto3";

message Person {
  string name = 1;
  repeated string phone_numbers = 2;
}
````
```go
package main

import (
  "fmt"
  "github.com/golang/protobuf/proto"
  "example.com/person"
)

func main() {
  p := &person.Person{
    Name: "John",
    PhoneNumbers: []string{"123456789", "987654321"},
  }

  // 序列化消息
  data, _ := proto.Marshal(p)

  // 反序列化消息
  p2 := &person.Person{}
  _ = proto.Unmarshal(data, p2)

  fmt.Println(p2.Name)           // 输出: John
  fmt.Println(p2.PhoneNumbers)   // 输出: [123456789 987654321]
  fmt.Println(p2.PhoneNumbers[0]) // 输出: 123456789
}
````
举一个复杂一点的例子
```protobuf
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
````

- **map：这是成对的键/值字段类型。**
```protobuf
syntax = "proto3";

message Person {
  string name = 1;
  map<string, int32> scores = 2;
}
````
```go
package main

import (
  "fmt"
  "github.com/golang/protobuf/proto"
  "example.com/person"
)

func main() {
  p := &person.Person{
    Name: "John",
    Scores: map[string]int32{
      "Math": 90,
      "English": 85,
    },
  }

  // 序列化消息
  data, _ := proto.Marshal(p)

  // 反序列化消息
  p2 := &person.Person{}
  _ = proto.Unmarshal(data, p2)

  fmt.Println(p2.Name)         // 输出: John
  fmt.Println(p2.Scores)       // 输出: map[Math:90 English:85]
  fmt.Println(p2.Scores["Math"])   // 输出: 90
}
````

- **oneof 类型: 用于表示一组互斥的字段, 在一个消息中，只能同时存在oneof中的一个字段，而不能同时存在多个字段。**
```protobuf
syntax = "proto3";

message Shape {
  oneof shape_type {
    Circle circle = 1;
    Rectangle rectangle = 2;
    Triangle triangle = 3;
  }
}

message Circle {
  float radius = 1;
}

message Rectangle {
  float width = 1;
  float height = 2;
}

message Triangle {
  float side1 = 1;
  float side2 = 2;
  float side3 = 3;
}
// 我们定义了一个Shape消息，其中包含了一个oneof字段shape_type。shape_type字段可以是Circle、Rectangle或Triangle中的一个。
// 这意味着在Shape消息中，只能同时存在一个shape_type字段。
// 
// 根据具体的需求，我们可以选择在Shape消息中使用Circle、Rectangle或Triangle中的一个字段。
// 例如，如果我们想表示一个圆形，我们可以设置circle字段，并提供半径值。如果我们想表示一个矩形，我们可以设置rectangle字段，并提供宽度和高度值。
````
  

### 1.2 分配字段编号
我们必须要为消息定义中的每个字段指定一个介于 1 和 536,870,911 之间的数字
- 给定的编号在该消息的所有字段中必须是唯一的
- 字段号 19,000 ~19,999 是为协议缓冲区实现保留的 (不能使用)
- 不能使用任何先前保留的字段编号或任何已分配给分机的字段编号 (字段编号决不能重复使用, 想用必须使用新的数字编号)

**一旦您的消息类型被使用，该数字就无法更改**，因为它标识消息传输格式中的字段 。“更改”字段编号相当于删除该字段并创建一个具有相同类型但新编号的新字段。
所以,初始化定义好字段的编号, 生成 proto 文件后, 后续即不能更改编号

**使用字段编号 1 到 15 作为最常设置的字段**, 这是个优化手段, 较低的字段数值在有线格式中占用较少的空间。例如，1 到 15 范围内的字段编号需要一个字节进行编码。16 到 2047 范围内的字段编号占用两个字节。

在 protobuf 中，删除字段可能会导致严重的问题, 当你不再需要一个非必需字段，并且在客户端代码中已经删除了所有对该字段的引用时，你可以从消息定义中删除该字段。
然而，你必须保留被删除字段的字段号（field number）。如果你不保留字段号，将来开发人员可能会重用该字段号，导致冲突。

我们有一个Person消息，包含了name和age两个字段：
```protobuf
syntax = "proto3";

message Person {
  string name = 1;
  int32 age = 2;
}
````
现在，我们决定不再需要age字段，并且在客户端代码中已经删除了所有对age字段的引用。为了遵循删除字段的最佳实践，我们需要进行以下操作：
- 保留被删除字段的字段号。在这个例子中，age字段的字段号是2，我们需要在消息定义中保留该字段号，以防将来有其他字段使用相同的字段号。
  ```protobuf
    syntax = "proto3";

    message Person {
    string name = 1;
    // Reserved field number 2
    reserved 2;
    }
  ````
- 保留被删除字段的字段名。这样可以确保对该消息进行JSON和TextFormat编码的解析仍然有效。
  ```protobuf
    syntax = "proto3";

    message Person {
        string name = 1;
        // Reserved field number 2
        // 这里的 9 to 11 与 9,10,11 是一致的
        reserved 2,  9 to 11;;
        // Reserved field name "age"
        reserved "age";
    }
  ````

### 1.3 包名定义 (package)
package您可以向文件添加可选说明符.proto，以防止协议消息类型之间发生名称冲突。类似于其他编程语言中的命名空间概念，用于组织和管理消息定义，以避免命名冲突。

```protobuf
syntax = "proto3";

package foo.bar;

message Open {
  string name = 1;
}

// 我们定义了 package 为 foo.bar, 这意味着 Open 消息定义只属于 "foo.bar" 包 
// 通过 package 语句，我们可以将消息定义组织在不同的包中，以便更好地管理和组织代码
````
```protobuf
syntax = "proto3";

package foo.bar;

message Foo {
  foo.bar.Open open = 1;
}
````

### 1.4 导入定义 (import)
目录结构
````
 tree
.
├── car.proto
└── user.proto
````
```protobuf
syntax = "proto3";

package user;
// 定义输出文件路径
option go_package = "/gen";

service UserService {
  rpc GetById(GetByIdReq) returns (GetByIdResp);
}

message GetByIdReq {
  uint64 id =1;
}

message GetByIdResp {
  User user = 1;
}

message User {
  uint64 id = 1;
  uint32 status = 2;
}
````
```protobuf
syntax = "proto3";

package car;
import "user.proto";
option go_package = "/gen";

message GetByIdReq {
  repeated user.User users = 1;
}

// 这里定义不同的包名可能对结果有不一样的影响
//package user;
//import "user.proto";
//option go_package = "/gen";
//
//message GetByIdReq {
//  repeated User users = 1;
//}
````

### 1.5 服务定义 (service)
```protobuf
service SearchService {
  rpc Search(SearchRequest) returns (SearchResponse);
}
````

## 2.Protocol Buffer 基础知识：Go
proto 文件
```protobuf
syntax = "proto3";

package user;
// go_package选项定义包的导入路径，其中将包含为此文件生成的所有代码。Go 包名称将是导入路径的最后一个路径组成部分。例如，我们的示例将使用包名称 “gen”。
option go_package = "/gen";

service UserService {
  rpc GetById(GetByIdReq) returns (GetByIdResp);
}

message GetByIdReq {
  uint64 id =1;
}

message GetByIdResp {
  User user = 1;
}

message User {
  uint64 id = 1;
  uint32 status = 2;
}
````
安装 [protoc](https://github.com/protocolbuffers/protobuf) 工具
```shell
# 根据系统选择合适版本 uname -a

$ wget wget https://github.com/protocolbuffers/protobuf/releases/download/v24.0-rc1/protoc-24.0-rc-1-linux-x86_64.zip

$ unzip protoc-3.19.4-linux-aarch_64.zip
 
$ mv ./bin/protoc /usr/local/bin/

# 验证
[root@99 ~]# protoc
````

安装 protoc-gen-go 工具 (确保已经安装 golang 环境, 并将 $GOPATH 加入至环境变量)
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ protoc --go_out=.   user.proto

# 生成的文件 user.pb.go
````
安装 protoc-gen-go-grpc 工具 (确保已经安装 golang 环境, 并将 $GOPATH 加入至环境变量)
```shell
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ protoc --go-grpc_out=.   user.proto

# 生成的文件 user_grpc.pb.go
````
合并以上两条命令
````shell
protoc --go_out=. --go-grpc_out=.   user.proto
````

