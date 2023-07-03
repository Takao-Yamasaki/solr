## Solrの起動方法
- Solrの起動方法
```
make up
```

- Solrコンテナへのログイン
```
make exec
```

- 管理画面のURL
```
http://localhost:8983/solr/#/
```

- サンプルコアとデータの作成
```
bin/post -c techproducts example/exampledocs/*
```
