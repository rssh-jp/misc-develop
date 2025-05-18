package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"elasticsearch-search-app/internal/usecase"
	// "elasticsearch-search-app/internal/interface/presenter" // 必要に応じてPresenterを利用
)

// ProductController は商品関連のHTTPリクエストを処理します。
type ProductController struct {
	productUseCase usecase.ProductUseCase
}

// NewProductController は新しい ProductController を作成します。
func NewProductController(uc usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: uc,
	}
}

// SearchProductsHandler は商品検索リクエストを処理するHTTPハンドラです。
func (c *ProductController) SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	query := queryValues.Get("q")

	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	products, err := c.productUseCase.SearchProducts(r.Context(), query)
	if err != nil {
		log.Printf("Error searching products: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
