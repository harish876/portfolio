package models

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Languages struct {
	Edges []LanguageEdge `json:"edges"`
}

type LanguageEdge struct {
	Node Language `json:"node"`
}

type PinnedItemNode struct {
	Name        string    `json:"name"`
	ID          string    `json:"id"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	HomepageURL string    `json:"homepageUrl"`
	UpdatedAt   string    `json:"updatedAt"`
	Languages   Languages `json:"languages"`
}

type Edge struct {
	Node PinnedItemNode `json:"node"`
}

type PinnedItems struct {
	TotalCount int    `json:"totalCount"`
	Edges      []Edge `json:"edges"`
}

type User struct {
	PinnedItems PinnedItems `json:"pinnedItems"`
}

type ResponseData struct {
	User User `json:"user"`
}
