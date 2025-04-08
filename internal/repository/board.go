package repository

import (
	"github.com/Jeecis/goapi/internal/models"
	"gorm.io/gorm"
)

type BoardRepository struct {
	db *gorm.DB
}

// NewBoardRepository creates a new instance of BoardRepository
func NewBoardRepository(db *gorm.DB) *BoardRepository {
	return &BoardRepository{db: db}
}

// Create adds a new board to the database
func (r *BoardRepository) Create(board *models.Board) error {
	return r.db.Create(board).Error
}

// GetByID retrieves a board by its ID
func (r *BoardRepository) GetByID(id string) (*models.Board, error) {
	var board models.Board
	err := r.db.First(&board, id).Error
	return &board, err
}

// Delete removes a board by its ID
func (r *BoardRepository) Delete(id uint) error {
	return r.db.Delete(&models.Board{}, id).Error
}

// GetAll retrieves all boards from the database
func (r *BoardRepository) GetAll() ([]models.Board, error) {
	var boards []models.Board
	err := r.db.Find(&boards).Error
	return boards, err
}

// Update modifies an existing board's parameters except its ID
func (r *BoardRepository) Update(board *models.Board) error {
	return r.db.Save(board).Error
}
