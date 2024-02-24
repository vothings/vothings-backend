package repository

import (
	"fmt"
	"vothings/internal/app/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type candidateRepository struct {
	db *gorm.DB
}

func (c *candidateRepository) Insert(user model.Candidates) error {
	user.ID = uuid.New()
	result := c.db.Create(&user)

	if result.Error != nil {
		return fmt.Errorf("error when insert candidate: %s", result.Error)
	}
	return nil
}

func (c *candidateRepository) GetAll() ([]model.Candidates, error) {
	var candidateDatas []model.Candidates
	result := c.db.Find(&candidateDatas)

	if result.Error != nil {
		return []model.Candidates{}, fmt.Errorf("error when getAll candidate: %s", result.Error)
	}
	return candidateDatas, nil
}

func (c *candidateRepository) Update(id uuid.UUID, candidate model.Candidates) error {
	result := c.db.Model(&candidate).Where("id = ?", id).Update("name", candidate.Name)

	if result.Error != nil {
		return fmt.Errorf("error when update candidate: %s", result.Error)
	}
	return nil
}

func (c *candidateRepository) Delete(id uuid.UUID) error {
	var candidates model.Candidates
	result := c.db.Where("id = ?", id).Delete(candidates)

	if result.Error != nil {
		return fmt.Errorf("error when delete candidate: %s", result.Error)
	}
	return nil
}

func NewCandidateRepository(db *gorm.DB) *candidateRepository {
	return &candidateRepository{
		db: db,
	}
}
