# go-solr-sample
## Solrのバージョン
```
$ solr version
6.3.0
```
## アプリケーションの起動
- 起動方法
```
make up
```

- Solrコンテナへのログイン
```
make solr
```

## アプリケーションの停止
```
make down
```

- 管理画面のURL
```
http://localhost:8983/solr/#/
```

- サンプル検索アプリケーション: SolritasのURL
```
http://localhost:8983/solr/techproducts/browse
```

## curl
### 1. アプリケーションの起動
```
make up
```
### 2.書籍データ用コアの作成とスキーマ定義ファイルのセットアップ
- ホスト側で実行
```
make setup
```

## 3.`sample-books.json`の登録(インデクシング)
- ホスト側で実行
```
make add
```

## 4.インデックスの更新
- ホスト側で実行
```
make update
```

## 5.インデックスの削除
- ホスト側で実行
```
make delete
```

## 8.コミットせずに登録
- ホスト側で実行
```
make nocommit
```
## 9.ロールバック
- ホスト側で実行
```
make rollback
```
