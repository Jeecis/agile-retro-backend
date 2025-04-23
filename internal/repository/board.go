package repository

import (
	"errors"
	"log"

	"github.com/Jeecis/goapi/internal/models"
	"gorm.io/gorm"
)

type BoardRepository struct {
	db *gorm.DB
}

var ErrBoardNotFound = errors.New("board not found")

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
	err := r.db.Where("id = ?", id).First(&board).Error
	return &board, err
}

// Delete removes a board by its ID
func (r *BoardRepository) Delete(id uint) error {
	return r.db.Delete(&models.Board{}, id).Error
}

// Update modifies an existing board's parameters except its ID
func (r *BoardRepository) Update(board *models.Board) error {
	return r.db.Save(board).Error
}

// BoardExists checks if a board with the given ID exists
func (r *BoardRepository) BoardExists(id string) bool {
	var board models.Board
	err := r.db.Where("id = ?", id).First(&board).Error
	return err == nil
}

// DelIDExists checks if a board with the given deletion ID exists
func (r *BoardRepository) DelIDExists(delID string) bool {
	var board models.Board
	err := r.db.Where("deletion_id = ?", delID).First(&board).Error
	return err == nil
}

// Retrieves a board by its deletion ID
func (r *BoardRepository) GetBoardByDelID(delID string) (*models.Board, error) {
	var board models.Board
	log.Print("delID: ", delID)
	err := r.db.Where("deletion_id = ?", delID).First(&board).Error
	return &board, err
}

func (r *BoardRepository) DeleteByDelID(delID string) error {
	result := r.db.Where("deletion_id = ?", delID).Delete(&models.Board{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrBoardNotFound
	}
	return nil
}
