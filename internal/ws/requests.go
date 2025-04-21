package ws

type joinBoard struct {
	BoardID string `json:"id"`
}

type CreateRecord struct {
	BoardID  string `json:"id"`
	ColumnID string `json:"column_id"`
	Text     string `json:"text"`
}
