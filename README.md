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
* [ ] protocol bufferを定義
* [ ] gRPCでコード生成
* [ ] データを永続化出来るように
* [ ] REST API作る。[grpc-gateway](https://github.com/gengo/grpc-gateway)使ってもいいかも
 - http://yugui.jp/articles/893
* [ ] APIを元にJSでUIを作る
