package storage

import "errors"

// TODO: implement PostgreSQL storage

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url already exists")
)
