package db

import (
	"github.com/hackbutteers/groupstore/errors"
)

// Common errors.
var (
	ErrNotFound         = errors.ErrNotFound
	ErrReadOnly         = errors.New("db: read-only mode")
	ErrSnapshotReleased = errors.New("db: snapshot released")
	ErrIterReleased     = errors.New("db: iterator released")
	ErrClosed           = errors.New("db: closed")
)
