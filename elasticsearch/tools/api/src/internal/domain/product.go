package domain

// Product は検索対象のデータ構造を表します。
type Product struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
	// 他のフィールド...
}
