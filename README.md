# Apache Solr入門
## Solrのバージョン
```
$ solr version
6.3.0
```
## Solrの起動方法
- Solrの起動方法
```
make up
```

- Solrコンテナへのログイン
```
make exec
```

## Solrの停止方法
```
make down
```

- 管理画面のURL
```
http://localhost:8983/solr/#/
```

## サンプルコアの作成

- サンプルコア: techproductsの作成(コンテナにログインした後)
```
bin/solr create_core -c techproducts
```

- サンプルデータの投入
```
bin/post -c techproducts /opt/solr/example/exampledocs/*.xml
```

- サンプル検索アプリケーション: SolritasのURL
```
http://localhost:8983/solr/techproducts/browse
```

## 書籍データハンズオン
### 1. Solrの起動
```
make up
```
### 2.書籍データ用コアの作成
```
make exec
bin/solr create_core -c solrbook -d basic_configs
```
### 3.スキーマ定義
セットアップ用の`schema.xml`と`solrconfig.xml`を`solrbook/conf/配下にコピーする
```
cd server/solr
cp solrbook_setup/schema.xml solrbook/conf/
cp solrbook_setup/solrconfig.xml solrbook/conf/
```

## 4.solrコレクションをリロード
- ホスト側で実行
```
curl "http://localhost:8983/solr/admin/cores?action=RELOAD&core=solrbook"
```

## 5.`sample-books.json`の登録(インデクシング)
- ホスト側で実行
```
curl "http://localhost:8983/solr/solrbook/update?commit=true" --data-binary @sample-books.json -H 'Content-Type: application/json; charset=utf8' -T "sample-books.json" -X POST
```
