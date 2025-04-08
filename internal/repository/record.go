package repository

import (
	"github.com/Jeecis/goapi/internal/models"
	"gorm.io/gorm"
)

type RecordRepository struct {
	db *gorm.DB
}

// NewRecordRepository creates a new instance of RecordRepository
func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

// Create adds a new record to the database
func (r *RecordRepository) Create(record *models.Record) error {
	return r.db.Create(record).Error
}

// GetByID retrieves a record by its ID
func (r *RecordRepository) GetByID(id string) (*models.Record, error) {
	var record models.Record
	err := r.db.First(&record, id).Error
	return &record, err
}

// Delete removes a record by its ID
func (r *RecordRepository) Delete(id uint) error {
	return r.db.Delete(&models.Record{}, id).Error
}

// GetAll returns all records from the database
func (r *RecordRepository) GetAll() ([]models.Record, error) {
	var record []models.Record
	err := r.db.Find(&record).Error
	return record, err
}

// Update modifies an existing record in the database
func (r *RecordRepository) Update(record *models.Record) error {
	return r.db.Save(record).Error
}
