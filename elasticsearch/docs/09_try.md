```
// インデックスを取得
// vはverbose。詳細な出力を有効
// hで表示する項目のヘッダーを指定
// sでソートするヘッダーを指定
GET /_cat/indices?v&h=index,docs.count,store.size&s=index

// インデックスの作成（Create Index）
// 新しいインデックスを作成する操作。マッピングを指定してデータ構造を定義します。
PUT /users-test
{
  "mappings": {
    "properties": {
      "user_id": {
        "type": "integer"
      },
      "name": {
        "type": "text"
      },
      "email": {
        "type": "keyword"
      },
      "age": {
        "type": "integer"
      }
    }
  }
}

POST /_aliases
{
  "actions": [
    {
      "add": {
        "index": "users-test",
        "alias": "users"
      }
    }
  ]
}

// ドキュメントのバルク操作（Bulk）
// 複数のドキュメントを一度にインデックス、更新、削除する操作。効率的にデータ操作を行うために使用します。
POST /users/_bulk
{ "index": { "_id": 1 } }
{ "user_id": 1, "name": "John Doe", "email": "john.doe@example.com", "age": 31 }
{ "index": { "_id": 2 } }
{ "user_id": 2, "name": "Jane Smith", "email": "jane.smith@example.com", "age": 28 }
{ "index": { "_id": 3 } }
{ "user_id": 3, "name": "Alice Johnson", "email": "alice.johnson@example.com", "age": 45 }
{ "index": { "_id": 4 } }
{ "user_id": 4, "name": "Bob Brown", "email": "bob.brown@example.com", "age": 22 }
{ "index": { "_id": 5 } }
{ "user_id": 5, "name": "Charlie Davis", "email": "charlie.davis@example.com", "age": 39 }
{ "index": { "_id": 6 } }
{ "user_id": 6, "name": "Eve Green", "email": "eve.green@example.com", "age": 51 }
{ "index": { "_id": 7 } }
{ "user_id": 7, "name": "Frank Harris", "email": "frank.harris@example.com", "age": 33 }
{ "index": { "_id": 8 } }
{ "user_id": 8, "name": "Grace Lee", "email": "grace.lee@example.com", "age": 29 }
{ "index": { "_id": 9 } }
{ "user_id": 9, "name": "Henry Martin", "email": "henry.martin@example.com", "age": 48 }
{ "index": { "_id": 10 } }
{ "user_id": 10, "name": "Ivy Nelson", "email": "ivy.nelson@example.com", "age": 25 }

// データの検索（Search）
// インデックス内のドキュメントを検索する基本的な操作。クエリを使って特定の条件に合ったデータを見つけることができます。
GET /users/_search

// データの更新（Update）
// 既存のドキュメントを更新する操作。部分的な更新が可能です。
POST /users/_update/1
{
  "doc": {
    "name": "John Doe Updated",
    "email": "john.doe.updated@example.com"
  }
}

// データの削除（Delete）
// 指定したドキュメントを削除する操作。ドキュメントのIDを指定して削除します。
DELETE /users/_doc/1

// インデックスの削除（Delete Index）
// 指定したインデックスを削除する操作。インデックス全体が削除されます。
DELETE /users

// データの集計（Aggregation）
// 特定のフィールドに基づいてデータを集計する操作。統計情報やヒストグラムを生成するために使用されます。
GET /users/_search
{
  "aggs": {
    "user_counts": {
      "terms": {
        "field": "age"
      }
    }
  }
}

// インデックスの情報を取得（Get Index Information）
// 特定のインデックスの詳細情報を取得する操作。
GET /users/_settings

// 全インデックスの概要を取得（Get All Indices Overview）
// 全インデックスの概要情報を取得する操作。インデックスの名前、状態、ドキュメント数などが表示されます。
GET /_cat/indices?v

// インストールされているプラグインの情報を取得（Get Installed Plugins Information）
// 現在インストールされているプラグインの一覧と詳細情報を取得する操作。
GET /_cat/plugins?v
```
