# go-solr-sample
## Solrのバージョン
```
$ solr version
6.3.0
```
## コンテナの起動
```
make up
```

## コンテナへのログイン
- Solrコンテナへのログイン
```
make solr
```
- Goコンテナへのログイン
```
make go
```

## コンテナの停止
```
make down
```

## 管理画面のURL

- 管理画面のURL
```
http://localhost:8983/solr/#/
```
- 検索方法
- [solrbook][Query][q]に次のように入力して、[Execute Query]を押下
```
title:"Apache Solr 入門"
```

- サンプル検索アプリケーション: SolritasのURL
```
http://localhost:8983/solr/techproducts/browse
```

## Goアプリケーション

### 1.書籍データ用コアの作成・スキーマ定義ファイルのセットアップ
`make up`を実行後に、実行
```
make setup-go
```

### 2.アプリケーションの起動
- `prompt ui`を使用しており、実行したいメソッドを選択できる
```
make run
```


## curl

### 1.書籍データ用コアの作成・スキーマ定義ファイルのセットアップ、インデクシング
- ホスト側で実行
```
make setup
```

## 2.`sample-books.json`の登録(インデクシング)
- ホスト側で実行
```
make add
```

## 3.インデックスの更新
- ホスト側で実行
```
make update
```

## 4.インデックスの削除
- ホスト側で実行
```
make delete
```

## 5.コミットせずに登録
- ホスト側で実行
```
make nocommit
```
## 6.ロールバック
- ホスト側で実行
```
make rollback
```

## 7. コアの削除
- ホスト側で実行
```
make rmcore
```
