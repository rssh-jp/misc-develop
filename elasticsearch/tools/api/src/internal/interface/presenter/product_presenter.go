package presenter

import "elasticsearch-search-app/internal/domain"

// ProductResponse はAPIで返す商品情報の構造です。
type ProductResponse struct {
	ID          string  `json:"id"`
	DisplayName string  `json:"display_name"`
	Price       float64 `json:"price"`
}

// ProductsSearchResponse は商品検索結果のAPIレスポンス構造です。
type ProductsSearchResponse struct {
	Products []ProductResponse `json:"products"`
	Count    int               `json:"count"`
}

// NewProductSearchResponse はドメインモデルのProductスライスからAPIレスポンスを生成します。
func NewProductSearchResponse(products []*domain.Product) ProductsSearchResponse {
	var productResponses []ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, ProductResponse{
			ID:          p.ID,
			DisplayName: p.Name, // 例: ドメインモデルと異なるフィールド名でマッピング
			Price:       p.Price,
		})
	}
	return ProductsSearchResponse{
		Products: productResponses,
		Count:    len(productResponses),
	}
}
