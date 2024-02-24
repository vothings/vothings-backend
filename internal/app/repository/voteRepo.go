package repository

import (
	"fmt"
	"vothings/internal/app/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type voteRepository struct {
	db *gorm.DB
}

func (v *voteRepository) Insert(vote model.Votes) error {
	vote.ID = uuid.New()
	result := v.db.Create(&vote)

	if result.Error != nil {
		return fmt.Errorf("error when insert vote: %s", result.Error)
	}
	return nil
}

func NewVoteRepository(db *gorm.DB) *voteRepository {
	return &voteRepository{
		db: db,
	}
}
