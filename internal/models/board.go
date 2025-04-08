package models

type Record struct {
	BoardID  string `json:"id"`
	ColumnID string `json:"column_id"`
	Text     string `json:"text"`
	Name     string `json:"name"`
	Likes    string `json:"likes"`
}

type Column struct {
	BoardID  string `json:"id"`
	ColumnID string `json:"column_id"`
	Name     string `json:"name"`
}

type Board struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DeletionID string `json:"deletion_id"`
}
