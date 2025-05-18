package usecase

import (
	"context"
	"elasticsearch-search-app/internal/domain"
)

// ProductUseCase は商品に関連するユースケースのインターフェースです。
type ProductUseCase interface {
	SearchProducts(ctx context.Context, query string) ([]*domain.Product, error)
}

// productUseCaseImpl は ProductUseCase の実装です。
type productUseCaseImpl struct {
	productRepo ProductRepository // ProductRepositoryインターフェースに依存
}

// NewProductUseCase は新しい ProductUseCase を作成します。
func NewProductUseCase(repo ProductRepository) ProductUseCase {
	return &productUseCaseImpl{
		productRepo: repo,
	}
}

// SearchProducts は指定されたクエリで商品を検索します。
func (uc *productUseCaseImpl) SearchProducts(ctx context.Context, query string) ([]*domain.Product, error) {
	if query == "" {
		return []*domain.Product{}, nil // またはエラーを返す
	}

	products, err := uc.productRepo.Search(ctx, query)
	if err != nil {
		return nil, err // エラーを上位に伝播
	}
	return products, nil
}
