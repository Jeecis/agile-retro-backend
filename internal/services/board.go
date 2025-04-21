package service

import (
	"github.com/Jeecis/goapi/internal/models"
	"github.com/Jeecis/goapi/internal/repository"
)

var firstColumn = "What went well?"
var secondColumn = "What didn't go well?"
var thirdColumn = "What will you do to improve?"

func CreateBoard(boardRepo *repository.BoardRepository, columnRepo *repository.ColumnRepository, name string) (*models.Board, error) {

	// check if ID exists
	id, err := boardIDGenerate()
	if err != nil {
		return nil, err
	}

	for {
		if !boardRepo.BoardExists(id) {
			break
		}
		id, err = boardIDGenerate()
		if err != nil {
			return nil, err
		}
	}

	// check if deletion ID exists
	delID, err := deletionIDGenerate()
	if err != nil {
		return nil, err
	}

	for {
		if !boardRepo.DelIDExists(id) {
			break
		}

		delID, err = deletionIDGenerate()
		if err != nil {
			return nil, err
		}
	}

	board := models.Board{
		ID:         id,
		Name:       name,
		DeletionID: delID,
	}

	if err := boardRepo.Create(&board); err != nil {
		return nil, err
	}

	if err := initColumns(board.ID, columnRepo); err != nil {
		return nil, err
	}

	return &board, nil
}

// initializes 3 columns with predefined names
func initColumns(boardID string, columnRepo *repository.ColumnRepository) error {
	colID1, err := generateUUID()
	if err != nil {
		return err
	}

	colID2, err := generateUUID()
	if err != nil {
		return err
	}

	colID3, err := generateUUID()
	if err != nil {
		return err
	}

	columns := []models.Column{
		{BoardID: boardID, Name: firstColumn, ColumnID: colID1},
		{BoardID: boardID, Name: secondColumn, ColumnID: colID2},
		{BoardID: boardID, Name: thirdColumn, ColumnID: colID3},
	}

	for _, column := range columns {
		if err := columnRepo.Create(&column); err != nil {
			return err
		}
	}

	return nil
}

func GetBoard(
	boardRepo *repository.BoardRepository,
	columnRepo *repository.ColumnRepository,
	recordRepo *repository.RecordRepository,
	id string,
) (*models.BoardFull, error) {
	fullBoard := &models.BoardFull{}

	board, err := boardRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	fullBoard.Board = *board

	columns, err := columnRepo.QueryBoardColumns(id)
	if err != nil {
		return nil, err
	}

	for _, col := range columns {
		records, err := recordRepo.QueryColumnRecords(col.ColumnID)
		if err != nil {
			return nil, err
		}
		fullBoard.Columns = append(fullBoard.Columns, models.BoardColumnWithRecords{
			Column:  col,
			Records: records,
		})
	}

	return fullBoard, nil

}

func CreateRecord(
	boardRepo *repository.BoardRepository,
	columnRepo *repository.ColumnRepository,
	recordRepo *repository.RecordRepository,
	recordData models.CreateRecord,
) (*models.Record, error) {
	// Step 1: validate board/column
	if !boardRepo.BoardExists(recordData.BoardID) {
		return nil, repository.ErrBoardNotFound
	}
	if !columnRepo.ColumnExists(recordData.ColumnID) {
		return nil, repository.ErrColumnNotFound
	}

	// Step 2: get existing records in the column (sorted by position)
	records, err := recordRepo.GetByColumnSorted(recordData.ColumnID)
	if err != nil {
		return nil, err
	}

	// Step 3: calculate the new position (at the bottom by default)
	newPos := 1000.0 // default if column is empty
	if len(records) > 0 {
		last := records[len(records)-1]
		newPos = last.Position + 1.0
	}

	recordID, err := generateUUID()
	if err != nil {
		return nil, err
	}

	// Step 4: create and save the record
	record := &models.Record{
		BoardID:  recordData.BoardID,
		ColumnID: recordData.ColumnID,
		RecordID: recordID,
		Text:     recordData.Text,
		Likes:    0,
		Position: newPos,
	}
	if err := recordRepo.Create(record); err != nil {
		return nil, err
	}

	return record, nil
}

func UpdateRecord(
	boardRepo *repository.BoardRepository,
	columnRepo *repository.ColumnRepository,
	recordRepo *repository.RecordRepository,
	recordData models.Record,
) (*models.Record, error) {
	// Step 1: check if record exists
	existingRecord, err := recordRepo.GetByID(recordData.RecordID)
	if err != nil {
		return nil, err
	}
	if existingRecord == nil {
		return nil, repository.ErrRecordNotFound
	}

	// Step 2: validate board/column only if updated
	if recordData.BoardID != "" && recordData.BoardID != existingRecord.BoardID {
		if !boardRepo.BoardExists(recordData.BoardID) {
			return nil, repository.ErrBoardNotFound
		}
		existingRecord.BoardID = recordData.BoardID
	}

	if recordData.ColumnID != "" && recordData.ColumnID != existingRecord.ColumnID {
		if !columnRepo.ColumnExists(recordData.ColumnID) {
			return nil, repository.ErrColumnNotFound
		}
		existingRecord.ColumnID = recordData.ColumnID
	}

	// Step 3: selectively update only non-zero fields
	if recordData.Text != "" {
		existingRecord.Text = recordData.Text
	}
	if recordData.Likes != 0 {
		existingRecord.Likes = recordData.Likes
	}

	// Step 4: update in DB
	if err := recordRepo.Update(existingRecord); err != nil {
		return nil, err
	}

	return existingRecord, nil
}

func DeleteRecord(
	boardRepo *repository.BoardRepository,
	columnRepo *repository.ColumnRepository,
	recordRepo *repository.RecordRepository,
	recordID string,
) error {
	// Step 1: check if record exists
	existingRecord, err := recordRepo.GetByID(recordID)
	if err != nil {
		return err
	}
	if existingRecord == nil {
		return repository.ErrRecordNotFound
	}

	// Step 2: delete the record
	if err := recordRepo.DeleteByRecordID(existingRecord.RecordID); err != nil {
		return err
	}

	return nil
}

// Add this to the board.go service file
func MoveRecord(
	recordRepo *repository.RecordRepository,
	moveData models.MoveRecord,
) (*models.Record, error) {
	// Get the record to move
	record, err := recordRepo.GetByID(moveData.RecordID)
	if err != nil {
		return nil, err
	}

	// Update column ID if changed
	record.ColumnID = moveData.TargetColumnID

	// Update position
	record.Position = moveData.NewPos

	// Save the updated record
	if err := recordRepo.Update(record); err != nil {
		return nil, err
	}

	// Fetch all records in the target column, sorted by position
	records, err := recordRepo.GetByColumnSorted(moveData.TargetColumnID)
	if err != nil {
		return nil, err
	}

	// Reorder positions to ensure uniqueness
	for i, rec := range records {
		if rec.RecordID == record.RecordID {
			continue
		}

		if rec.Position == record.Position && moveData.OldPos > moveData.NewPos {
			rec.Position = moveData.NewPos + 1 // Adjust position to avoid conflict
			if err := recordRepo.Update(record); err != nil {
				return nil, err
			}
			i = int(rec.Position)
		} else if rec.Position == record.Position && moveData.OldPos < moveData.NewPos {
			rec.Position = moveData.NewPos - 1 // Adjust position to avoid conflict
			if err := recordRepo.Update(record); err != nil {
				return nil, err
			}
			i = int(rec.Position) + 1
		} else {
			rec.Position = float64(i) + 1 // Reorder positions
		}
		if err := recordRepo.Update(&rec); err != nil {
			return nil, err
		}
	}

	return record, nil
}
