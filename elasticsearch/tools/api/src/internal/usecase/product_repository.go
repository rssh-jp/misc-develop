package usecase

import (
	"context"
	"elasticsearch-search-app/internal/domain"
)

// ProductRepository は商品データへのアクセスを提供するインターフェースです。
type ProductRepository interface {
	Search(ctx context.Context, query string) ([]*domain.Product, error)
	// 必要に応じて他のメソッド (FindByID, Store など) を定義
}
