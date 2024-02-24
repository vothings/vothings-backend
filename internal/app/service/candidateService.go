package service

import (
	"vothings/internal/app/dto"
	"vothings/internal/app/model"

	"github.com/google/uuid"
)

type CandidateRepository interface {
	Insert(user model.Candidates) error
	GetAll() ([]model.Candidates, error)
	Update(id uuid.UUID, candidate model.Candidates) error
	Delete(id uuid.UUID) error
}

type candidateService struct {
	repo CandidateRepository
}

func (c *candidateService) CreateCandidate(user dto.CandidateDtoRequest) error {
	if err := c.repo.Insert(model.Candidates{
		Name: user.Name,
	}); err != nil {
		return err
	}
	return nil
}

func (c *candidateService) GetAllCandidate() ([]model.Candidates, error) {
	candidates, err := c.repo.GetAll()

	if err != nil {
		return []model.Candidates{}, err
	}

	return candidates, nil
}

func (c *candidateService) UpdateCandidate(id uuid.UUID, candidate dto.CandidateDtoRequest) error {
	if err := c.repo.Update(id, model.Candidates{
		Name: candidate.Name,
	}); err != nil {
		return err
	}
	return nil
}

func (c *candidateService) DeleteCandidate(id uuid.UUID) error {
	if err := c.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func NewCandidateService(repo CandidateRepository) candidateService {
	return candidateService{
		repo: repo,
	}
}
