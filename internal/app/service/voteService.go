package service

import (
	"fmt"
	"vothings/internal/app/dto"
	"vothings/internal/app/model"

	"github.com/google/uuid"
)

type VoteRepository interface {
	Insert(vote model.Votes) error
}

type voteService struct {
	repo VoteRepository
}

func (v *voteService) CreateVote(voteData dto.VoteDtoRequest) error {

	candidateId, err := uuid.Parse(voteData.CandidateId)

	if err != nil {
		return fmt.Errorf("error when parse uuid: %s", err)
	}

	userId, err := uuid.Parse(voteData.UserId)
	if err != nil {
		return fmt.Errorf("error when parse uuid: %s", err)
	}

	data := model.Votes{
		CandidateId: candidateId,
		UserID:      userId,
	}

	if err := v.repo.Insert(data); err != nil {
		return err
	}
	return nil
}

func NewVoteService(repo VoteRepository) *voteService {
	return &voteService{
		repo: repo,
	}
}
