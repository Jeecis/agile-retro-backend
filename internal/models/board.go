package models

type Record struct {
	BoardID  string  `json:"board_id"`
	ColumnID string  `json:"column_id"`
	RecordID string  `json:"record_id"`
	Position float64 `json:"position"`
	Text     string  `json:"text"`
	Likes    int     `json:"likes"`
}

type Column struct {
	BoardID  string `json:"board_id"`
	ColumnID string `json:"column_id"`
	Name     string `json:"name"`
}

type Board struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DeletionID string `json:"deletion_id"`
}

type BoardFull struct {
	Board   Board                    `json:"board"`
	Columns []BoardColumnWithRecords `json:"columns"`
}

type BoardColumnWithRecords struct {
	Column  Column   `json:"column"`
	Records []Record `json:"records"`
}
