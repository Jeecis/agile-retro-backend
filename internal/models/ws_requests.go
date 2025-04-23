package models

type JoinBoard struct {
	BoardID string `json:"id"`
}
type DeleteRecord struct {
	RecordID string `json:"id"`
}

type CreateRecord struct {
	BoardID  string `json:"id"`
	ColumnID string `json:"column_id"`
	Text     string `json:"text"`
}

type DeleteBoard struct {
	DelID string `json:"deletion_id"`
}

type MoveRecord struct {
	RecordID       string  `json:"record_id"`
	SourceColumnID string  `json:"source_column_id"`
	TargetColumnID string  `json:"target_column_id"`
	OldPos         float64 `json:"old_position"`
	NewPos         float64 `json:"new_position"`
}
