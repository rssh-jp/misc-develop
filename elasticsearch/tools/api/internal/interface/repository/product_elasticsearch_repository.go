package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"elasticsearch-search-app/internal/domain"
	"elasticsearch-search-app/internal/usecase" // Usecase層のインターフェースを利用

	"github.com/elastic/go-elasticsearch/v8"
	// "github.com/elastic/go-elasticsearch/v8/esapi" // esapiは通常client.Searchなどでラップされている
)

// elasticsearchProductRepository は ProductRepository の Elasticsearch 実装です。
type elasticsearchProductRepository struct {
	client    *elasticsearch.Client
	indexName string
}

// NewElasticsearchProductRepository は新しい elasticsearchProductRepository を作成します。
func NewElasticsearchProductRepository(client *elasticsearch.Client, indexName string) usecase.ProductRepository {
	return &elasticsearchProductRepository{
		client:    client,
		indexName: indexName,
	}
}

// Search は Elasticsearch を使って商品を検索します。
func (r *elasticsearchProductRepository) Search(ctx context.Context, query string) ([]*domain.Product, error) {
	var products []*domain.Product

	var buf bytes.Buffer
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
            "bool": map[string]interface{}{
                "must": []interface{}{
                    map[string]interface{}{
                        "multi_match": map[string]interface{}{
				            "query":  query,
				            "fields": []string{"title", "text"}, // 検索対象フィールド
                        },
                    },
                },
            },
		},
        "size": 10, // 最大100件の結果を取得
        "sort": []interface{}{
            map[string]interface{}{
                "_score": map[string]interface{}{
                    "order": "desc", // スコアの降順
                },
            },
        },
	}
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil, fmt.Errorf("failed to encode query: %w", err)
	}

    log.Println("Elasticsearch search query:", searchQuery)

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(r.indexName),
		r.client.Search.WithBody(&buf),
		r.client.Search.WithTrackTotalHits(true),
		r.client.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response from Elasticsearch: %s", err)
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the Elasticsearch error response body: %s", err)
			return nil, fmt.Errorf("failed to parse Elasticsearch error response body: %w", err)
		}
		errMsg := fmt.Sprintf("Elasticsearch error: [%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
		log.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	var rmap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&rmap); err != nil {
		log.Printf("Error parsing the Elasticsearch response body: %s", err)
		return nil, fmt.Errorf("failed to parse Elasticsearch response body: %w", err)
	}

	hits, ok := rmap["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return products, nil // ヒットなし
	}

	for _, hit := range hits {
		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			log.Printf("Warning: _source is not a map[string]interface{} for hit: %v", hit)
			continue
		}

		product := domain.Product{
			ID: hit.(map[string]interface{})["_id"].(string),
		}
		if title, ok := source["title"].(string); ok {
			product.Title = title
		}
		if text, ok := source["text"].(string); ok {
			product.Text = text
		}
		if url, ok := source["url"].(string); ok {
			product.Url = url
		}
		products = append(products, &product)
	}

	return products, nil
}
