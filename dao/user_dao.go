//抽奖系统中用户数据的数据库操作

package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"lottery/models"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine:engine,
	}
}

func (d *UserDao) Get(uid int) *models.LtUser {
	data := &models.LtUser{
		Id:         uid,
	}
	ok, err := d.engine.Get(data)
	if ok && err == nil  {
		return data
	} else {
		data.Id = 0
		log.Println("user_dao.get error=", err)
		return nil
	}
}

func (d *UserDao) GetAll(page, size int) []models.LtUser {
	start := (page - 1) * size
	datalist := make([]models.LtUser, 0)
	err := d.engine.Asc("Id").
		Limit(size, start).Find(&datalist)

	if err != nil {
		log.Println("user_dao.GetAll error=", err)
		return nil
	}

	return datalist
}

func (d *UserDao) CountAll() int64 {
	num, err := d.engine.Count()
	if err != nil {
		log.Println("user_dao.CountAll error=", err)
		return -1
	}

	return num
}

func (d *UserDao) Update(data models.LtUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDao) Create(data *models.LtUser) error {
	_, err := d.engine.Insert(data)
	return err
}