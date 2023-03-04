package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"nome" validate:"len=10"`
	RG   string `json:"nome" validate:"len=9"`
}

func Validation(student *Aluno) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
