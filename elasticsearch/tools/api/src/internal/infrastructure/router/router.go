package router

import (
	"log"
	"net/http"
	"os"

	infraES "elasticsearch-search-app/internal/infrastructure/elasticsearch" // エイリアスで区別
	"elasticsearch-search-app/internal/interface/controller"
	"elasticsearch-search-app/internal/interface/repository"
	"elasticsearch-search-app/internal/usecase"
)

// NewRouter はHTTPルーターをセットアップして返します。
func NewRouter() http.Handler {
	// 1. Elasticsearch Client の初期化
	esClient, err := infraES.NewClient()
	if err != nil {
		log.Fatalf("Failed to initialize Elasticsearch client: %v", err)
	}

	// 2. Repository の初期化
	esIndexName := os.Getenv("ELASTICSEARCH_INDEX_NAME")
	if esIndexName == "" {
		esIndexName = "wiki" // デフォルトのインデックス名
		log.Printf("ELASTICSEARCH_INDEX_NAME not set, using default: %s", esIndexName)
	}
	productRepo := repository.NewElasticsearchProductRepository(esClient, esIndexName)

	// 3. Usecase の初期化
	productUseCase := usecase.NewProductUseCase(productRepo)

	// 4. Controller の初期化
	productController := controller.NewProductController(productUseCase)

	// 5. ルーティングの設定 (net/http の例)
	mux := http.NewServeMux()
	mux.HandleFunc("/search/products", productController.SearchProductsHandler) // GET /search/products?q=your_query

	return mux
}
