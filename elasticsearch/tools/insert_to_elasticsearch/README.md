# データ挿入まで
## 準備（データ作成含む）
```
wget https://dumps.wikimedia.org/jawiki/latest/jawiki-latest-pages-articles.xml.bz2
git@github.com:zaemyung/wikiextractor.git
python wikiextractor/WikiExtractor.py -o output -b 80M jawiki-latest-pages-articles.xml.bz2 --json
ls ./output/AA/* -d | xargs -L 1 -P 10 bash -c './reformat_to_ndjson.py $0'
```

## データ挿入
ここからはローカル環境のElasticsearchを立ち上げた状態で行う

### mapping作成
```
PUT /wiki-20250413_1
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0,
    "analysis": {
      "analyzer": {
        "ngram": {
          "filter": [
            "cjk_width",
            "lowercase"
          ],
          "char_filter": [
            "html_strip"
          ],
          "type": "custom",
          "tokenizer": "ngram"
        }
      },
      "tokenizer": {
        "ngram": {
          "token_chars": [
            "letter",
            "digit"
          ],
          "min_gram": "1",
          "type": "ngram",
          "max_gram": "2"
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "text": {
        "analyzer": "ngram",
        "type": "text"
      },
      "url": {
        "type": "keyword"
      },
      "id": {
        "type": "keyword"
      },
      "title": {
        "analyzer": "ngram",
        "type": "text"
      }
    }
  }
}
```

### エイリアス追加
```
POST /_aliases
{
  "actions": [
    {
      "add": {
        "index": "wiki-20250413_1",
        "alias": "wiki"
      }
    }
  ]
}
```

### データ挿入
```
ls ./output/AA/*_new.ndjson -d | xargs -L 1 -P 3 bash -c 'echo $0 ; cat $0 | curl -s -X POST -H '\''Content-Type: application/x-ndjson'\'' '\''http://127.0.0.1:9200/wiki/_bulk?pretty'\'' --data-binary @-;'
```

