## 方針
### Rest API
- DBはsqlite3でliteStreamを使いレプリケーションする
- Cloud RunでREST API+LiteStreamをコンテナ上に稼働する
### Front
- Cloudflare Workerにホスト
  - Cloudflare PagesがNode18のビルドにまだ対応していないため
- 基本スマホ用にしか作らない
  - PC対応はしない
## 課題
