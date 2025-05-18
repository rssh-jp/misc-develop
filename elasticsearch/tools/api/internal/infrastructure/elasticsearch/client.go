package elasticsearch

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

// NewClient は新しい Elasticsearch クライアントを作成して返します。
func NewClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_URLS"), // 例: "http://localhost:9200"
		},
		// 必要に応じて認証情報などを設定
		// Username:  os.Getenv("ELASTICSEARCH_USERNAME"),
		// Password:  os.Getenv("ELASTICSEARCH_PASSWORD"),
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
		return nil, err
	}

	// 接続テスト (オプション)
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting Elasticsearch info: %s", err)
		return nil, err
	}
	defer res.Body.Close()
	log.Printf("Elasticsearch client initialized, version: %s", elasticsearch.Version)

	return es, nil
}
