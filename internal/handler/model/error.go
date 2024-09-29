package model

import "errors"

var InvalidLevelErr = errors.New("invalid level")
var EmptyModuleName = errors.New("empty module name")

type ErrorResp struct {
	Error string `json:"error"`
}
