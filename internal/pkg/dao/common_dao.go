package dao

import (
	"github.com/pkg/errors"
)

type Common struct{}

type Do interface {
	TableName() string
}

func (Common) Create(r Do) error {
	return GetDB().Create(r).Error
}

func (Common) GetFirst(src Do, dst Do) error {
	err := GetDB().Where(src).First(dst).Error
	return errors.Wrap(err, "")
}