package model

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/internal/web/orm"
	"time"
)

type Category struct {
	Id          int
	ShortName   string
	Title       string
	Description string
	Sort        int
	UpdateAt    int64
	CreateAt    int64
	IsDel       int
}

type CategoryList []Category

type ConCategoryTags struct {
	Id          int
	ShortName   string
	Title       string
	Description string
	Sort        int
	Tags        []ConTag
}

func (c *Category) Insert() (rs int64, err error) {
	c.UpdateAt = time.Now().UnixNano()
	rs, err = orm.SocialBotOrm.Insert(c)
	if err != nil {
		return 0, errors.Wrap(err, "Insert failed")
	}
	return rs, nil
}

func (c *Category) UpdateById(id int) (rs int64, err error) {
	rs, err = orm.SocialBotOrm.Where("id = ? ", id).Update(c)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("updateById(%d) failed", id))
	}
	return rs, nil
}

func (c *Category) UpdateColsById(id int, cols ...string) (rs int64, err error) {
	c.UpdateAt = time.Now().UnixNano()
	cols = append(cols, "create_at")
	rs, err = orm.SocialBotOrm.Cols(cols...).Where("id = ? ", id).Update(c)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("UpdateColsById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (c *Category) GetOneById(id int) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Where("is_del=?", 0).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetOneById(%d) failed", id))
	}
	return rs, nil
}

func (c *Category) GetColsOneById(id int, cols ...string) (rs bool, err error) {
	rs, err = orm.SocialBotOrm.Where("id=?", id).Where("is_del=?", 0).Cols(cols...).Get(c)
	if err != nil {
		return rs, errors.Wrap(err, fmt.Sprintf("GetColsOneById(%d) cols(%s) failed", id, cols))
	}
	return rs, nil
}

func (cl *CategoryList) GetList() (err error) {
	err = orm.SocialBotOrm.Where("is_del=?", 0).Find(cl)
	if err != nil {
		return errors.Wrap(err, "get list failed")
	}
	return nil
}
