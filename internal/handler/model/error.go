package model

import "errors"

var InvalidLevelErr = errors.New("invalid level")
var EmptyModuleNameErr = errors.New("empty module name")

type ErrorResp struct {
	Error string `json:"error"`
}
