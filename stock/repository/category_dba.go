package repository

import (
	"FinalDSP/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CategoryDb struct {
	conn *gorm.DB
}

func NewCategoryDb(db *gorm.DB) *CategoryDb {
	return &CategoryDb{db}
}

func (caty *CategoryDb) Categories() ([]model.Category,[]error) {
	categories := []model.Category{}
	errs := caty.conn.Find(&categories).GetErrors()
	fmt.Println("from databasee",categories)
	if len(errs) > 0 {
		return nil,errs
	}
	return categories,errs
}

func (caty *CategoryDb) Category(Id uint)(*model.Category,[]error) {
	category := model.Category{}
	errs := caty.conn.First(&category,Id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &category,errs
}

func (caty *CategoryDb)  UpdateCategory(category *model.Category)(*model.Category,[]error) {
	cat := category
	errs := caty.conn.Save(cat).GetErrors()
	if len(errs) > 0 {
		return category,errs
	}
	return cat,errs
}

func (caty *CategoryDb)  DeleteCategory(Id uint)(*model.Category,[]error)  {
	cat,errs := caty.Category(Id)
	if len(errs) > 0 {
		return nil,errs
	}

	errs = caty.conn.Delete(cat,Id).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return cat,errs
}

func (caty *CategoryDb) SaveCategory(category *model.Category)(*model.Category,[]error) {
	cat := category
	errs := caty.conn.Create(cat).GetErrors()
	if len(errs) > 0  {
		return nil,errs
	}
	return cat,errs
}

func (caty *CategoryDb) StockCategory(category *model.Category)([]model.Stock,[]error){
	stocks:= []model.Stock{}
	cats,errs := caty.Category(category.ID)
	if len(errs) > 0{
		return nil,errs
	}
	errs = caty.conn.Model(cats).Related(stocks,"Stocks").GetErrors()
	if len(errs) > 0  {
		return nil,errs
	}
	return stocks,errs
}