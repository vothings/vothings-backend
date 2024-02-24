package model

import "github.com/google/uuid"

type Votes struct {
	ID          uuid.UUID
	CandidateId uuid.UUID
	UserID      uuid.UUID
}
