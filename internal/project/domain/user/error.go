package user

import "errors"

var (
	ErrNotFound   = errors.New("user.not_found")
	ErrNotCreated = errors.New("user.not_created")
	ErrNotDeleted = errors.New("user.not_deleted")
)
