## 方針
- DBはsqlite3でlitestreamを使いレプリケーションする
- Cloud RunでDBコンテナ、APIコンテナを稼働
- フロントは適当なところにホスト
## 課題
- [ ] alpineベースだとgo-sqlite3がCGOに依存しているらしく動かない