package uidgen

import "github.com/google/uuid"

type UIDGen interface {
	New() string
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (uidgen uidgen) New() string {
	return uuid.New().String()
}
