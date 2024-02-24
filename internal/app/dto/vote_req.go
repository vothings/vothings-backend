package dto

type VoteDtoRequest struct {
	CandidateId string `json:"candidateId"`
	UserId      string `json:"userId"`
}
