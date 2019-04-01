package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"lottery/models"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine:engine,
	}
}

func (d *CodeDao) Get(id int) *models.LtCode {
	data := &models.LtCode{Id:id}
	ok, err := d.engine.Get(&data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *CodeDao) GetAll(page, size int) []models.LtCode {
	start := (page - 1) * size
	codelist := make([]models.LtCode, 0)
	err := d.engine.Desc("Id").
		Limit(size, start).
		Find(&codelist)
	if err != nil {
		log.Println("code_dao error=", err)
	}
	return codelist
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtCode{})
	if err != nil {
		log.Println("code_dao.CountAll error=", err)
		return 0
	}
	return num
}

func (d *CodeDao) CountByGift(giftId int) int64 {
	num, err := d.engine.Where("gift_id=?",giftId).
		Count(&models.LtCode{})
	if err != nil {
		log.Println("code_dao.CountByGift error=", err)
		return 0
	}
	return num
}

func (d *CodeDao) Search(giftId int) []models.LtCode {
	codelist := make([]models.LtCode, 0)
	err := d.engine.Where("gift_id=?", giftId).
		Desc("id").
		Find(&codelist)
	if err != nil {
		log.Println("code_dao.Search error=", err)
	}
	return codelist
}

// 删除优惠券， 只更改其状态，不真正将其从表中删除
func (d *CodeDao) Delete(id int) error {
	data := &models.LtCode{
		Id:id,
		SysStatus:1,
	}
	_, err := d.engine.Id(id).Update(data)
	return err
}

// 更新数据， 指定某些列更新
func (d *CodeDao) Update(data *models.LtCode, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Create(data *models.LtCode) error {
	_, err := d.engine.Insert(data)
	return err
}