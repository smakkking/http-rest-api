package store

import "errors"

// смотрите, каким образом здесь создается иерархия ошибок
var (
	ErrUserNotFound = errors.New("user not found")
)
