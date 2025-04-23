package repository

import (
	"errors"

	"github.com/Jeecis/goapi/internal/models"
	"gorm.io/gorm"
)

type RecordRepository struct {
	db *gorm.DB
}

var ErrRecordNotFound = errors.New("record not found")

// NewRecordRepository creates a new instance of RecordRepository
func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

// Create adds a new record to the database
func (r *RecordRepository) Create(record *models.Record) error {
	return r.db.Create(record).Error
}

func (r *RecordRepository) GetByID(id string) (*models.Record, error) {
	var record models.Record
	err := r.db.Where("record_id = ?", id).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// DeleteByRecordID deletes the first record with the given record_id
func (r *RecordRepository) DeleteByRecordID(recordID string) error {
	result := r.db.Where("record_id = ?", recordID).Delete(&models.Record{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

// GetAll returns all records from the database
func (r *RecordRepository) GetAll() ([]models.Record, error) {
	var record []models.Record
	err := r.db.Find(&record).Error
	return record, err
}

// Update modifies an existing record in the database
func (r *RecordRepository) Update(record *models.Record) error {
	return r.db.Model(record).Where("record_id = ?", record.RecordID).Updates(record).Error
}

func (r *RecordRepository) QueryColumnRecords(boardID string) ([]models.Record, error) {
	var records []models.Record
	err := r.db.Where("column_id = ?", boardID).Find(&records).Error
	return records, err
}

func (r *RecordRepository) GetByColumnSorted(ColumnID string) ([]models.Record, error) {
	var records []models.Record
	err := r.db.Where("column_id = ?", ColumnID).Order("position").Find(&records).Error
	return records, err
}

func (r *RecordRepository) DeleteAllByBoardID(boardID string) error {
	result := r.db.Where("board_id = ?", boardID).Delete(&models.Record{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
