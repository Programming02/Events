package main

type Service struct {
	repo      Repository
	conflicts chan<- ConflictMessage
}

func NewService(repo Repository, conflicts chan<- ConflictMessage) Service {
	return Service{
		repo:      repo,
		conflicts: conflicts,
	}
}
