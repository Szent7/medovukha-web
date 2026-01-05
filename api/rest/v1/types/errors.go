package types

import "errors"

var (
	ErrContainerNotFound    = errors.New("container not found err")
	ErrContainerIsMedovukha = errors.New("container is medovukha err")
	ErrEmptyTags            = errors.New("tags list is empty err")
)

type ErrorCode string

const (
	ErrNotFound   ErrorCode = "NOT_FOUND"
	ErrValidation ErrorCode = "VALIDATION_ERR"
	ErrCore       ErrorCode = "MEDOVUKHA_CORE_ERR"
	ErrWeb        ErrorCode = "MEDOVUKHA_WEB_ERR"
)
