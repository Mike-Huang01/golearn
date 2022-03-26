package task1_error

import (
	"fmt"
	"gotime/internal/pkg/dao"
)

func GetFirstItemByAuthor(author string) (*Item) {
	c := dao.Common{}
	item := Item{Author: author}
	ret := Item{}
	err := c.GetFirst(item, &ret)
	if err != nil {
		// 记录错误及trace
		fmt.Printf("fc=GetFirstItemByAuthor | author:%s %+v\n", author, err)
		return nil
	}

	return &ret
}
