package model

import (
	"errors"
)

var (
	InvalidCredentialsErr = errors.New("invalid username or password")
)
