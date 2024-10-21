package errors

import "errors"

type RepoError struct {
	Message string `json:"message"`
}

func (r *RepoError) Error() string {
	return r.Message
}

func IsRepoError(err error) bool {
	var repoError *RepoError
	ok := errors.As(err, &repoError)
	return ok
}

var (
	ErrNotFound     = &RepoError{Message: "resource not found"}
	ErrNotConnected = &RepoError{Message: "not connected to database"}
)
