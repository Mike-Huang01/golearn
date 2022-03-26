package task1_error

import (
	"github.com/stretchr/testify/assert"
	"gotime/internal/pkg/dao"
	"testing"
	"time"
)

func TestError(t *testing.T) {
	dao.GetDB()
}

func TestCreateItem(t *testing.T) {
	item := Item{
		Title:  "学习go",
		Author: "mike",
		Date: time.Now(),
	}

	c := dao.Common{}
	err := c.Create(&item)
	assert.NoError(t, err)
}

func TestGetItemByAuthor(t *testing.T) {
	GetFirstItemByAuthor("mikey")
}
