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
