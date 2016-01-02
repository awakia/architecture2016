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
