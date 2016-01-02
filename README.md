# Modern Architecture ver. 2016

## 使用技術

* Server: Go
* Client: ES6 / React
* API Framework: gRPC
* Data Store: TBD
 - DB: PostgreSQL
 - AWS: DynamoDB / DynamoDB Local
 - Google Cloud Platform: Cloud Datastore etc. see https://cloud.google.com/docs/storing-your-data?hl=ja
* Infra: CoreOS / Docker

## 手順

* [ ] 必要なモデルを考える
* [ ] Protocol Buffersを定義
* [ ] gRPCでコード生成
* [ ] データを永続化出来るように
* [ ] REST API作る。[grpc-gateway](https://github.com/gengo/grpc-gateway)使ってもいいかも
 - http://yugui.jp/articles/893
* [ ] APIを元にJSでUIを作る

## 必要なモデル

レポジトリの概念は省略する

* User account:string:uniq name
* Issue user:references title content asignee:references
* Comment user:references issue:references content
* Label name color
* Labeling issue:references label:references user_id:integer
* Event action param type  # Issueから参照やラベルやAssigneeの追加や削除を担う

## Protocol Buffersの定義

### 準備

gRPCでrpc定義をすることを考えると、うまくREST形式とRPC形式を融合しないといけない。
ここでなんとなくのルールを作っておく。通常RPCでは

cf. https://talks.golang.org/2015/gotham-grpc.slide#1

```
service Google {
  rpc Search(Request) returns (Result)
}

message Request {
  string query = 1;
}

message Result {
  string title = 1;
  string url = 2;
  string snippet = 3;
}
```

```
message HelloRequest {
  string greeting = 1;
}

message HelloResponse {
  string reply = 1;
}

service HelloService {
  rpc SayHello(HelloRequest) returns
  (HelloResponse);
}
```

と言った形で、RequestとResultのペアをそれぞれrpcごとに定義していく形になる。

ここでResultはUserやLabelなどのモデルそのもので良い。Requestだけ考えてあげて

```
service UserService {
  rpc List() returns (stream User)  // index
  rpc Get(UserID) returns (User) // show
  rpc GetByName(UserName) returns (User)
}

message UserID {
  int32 id = 1;
}

message UserName {
  string name = 1;
}
```

か

```
service UserService {
  rpc List() returns (stream User)  // index
  rpc Get(UserParam) returns (User) // show
}

message UserParam {
  int32 id = 1;
  string name = 2;
}
```

のどちらかを使うようにする。
ポイントは

- index => list
- show => get
- new, edit, create, update, delete => そのまま

の命名規則を使うこと。

後者のほうが後々よいかも？？
