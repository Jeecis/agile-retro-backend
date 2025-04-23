package repository

import (
	"errors"

	"github.com/Jeecis/goapi/internal/models"
	"gorm.io/gorm"
)

type ColumnRepository struct {
	db *gorm.DB
}

var ErrColumnNotFound = errors.New("column not found")

// NewColumnRepository creates a new instance of ColumnRepository
func NewColumnRepository(db *gorm.DB) *ColumnRepository {
	return &ColumnRepository{db: db}
}

// Create adds a new Column to the database
func (r *ColumnRepository) Create(column *models.Column) error {
	return r.db.Create(column).Error
}

// GetByID retrieves a Column by its ID
func (r *ColumnRepository) GetByID(id string) (*models.Column, error) {
	var column models.Column
	err := r.db.First(&column, id).Error
	return &column, err
}

// Delete removes a Column by its ID
func (r *ColumnRepository) Delete(id uint) error {
	return r.db.Delete(&models.Column{}, id).Error
}

// GetAll retrieves all Columns from the database
func (r *ColumnRepository) GetAll() ([]models.Column, error) {
	var column []models.Column
	err := r.db.Find(&column).Error
	return column, err
}

// Update modifies an existing Column in the database
func (r *ColumnRepository) Update(column models.Column) error {
	return r.db.Save(column).Error
}

func (r *ColumnRepository) ColumnExists(id string) bool {
	var column models.Column
	err := r.db.Where("column_id = ?", id).First(&column).Error
	return err == nil
}

func (r *ColumnRepository) QueryBoardColumns(boardID string) ([]models.Column, error) {
	var columns []models.Column
	err := r.db.Where("board_id = ?", boardID).Find(&columns).Error
	return columns, err
}

func (r *ColumnRepository) DeleteAllByBoardID(boardID string) error {
	result := r.db.Where("board_id = ?", boardID).Delete(&models.Column{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
