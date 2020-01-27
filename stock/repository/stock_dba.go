package repository

import (
	"FinalDSP/model"
	"github.com/jinzhu/gorm"

)

type StockDb struct {
	conn *gorm.DB
}

func NewStockDB(db *gorm.DB) *StockDb {
	return &StockDb{db}
}

func (stockDb *StockDb) Stocks() ([]model.Stock,[]error) {
	stocks := []model.Stock{}
	errs := stockDb.conn.Find(&stocks).GetErrors()
	if len(errs) > 0 {
		return nil,errs
	}
	return stocks,errs
}

func (stockDb *StockDb) Stock(id uint) (*model.Stock,[]error)  {
	stock := model.Stock{}
	errs := stockDb.conn.First(&stock,id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &stock,errs
}

func (stockDb *StockDb) UpdateStock(stock *model.Stock)(*model.Stock,[]error)   {
	stk := stock
	errs := stockDb.conn.Save(stk).GetErrors()
	if len(errs) > 0 {
		return stock,errs
	}
	return stk,errs
}

func (stokDb *StockDb) DeleteStock(id uint)(*model.Stock,[]error)  {
	stk,errs := stokDb.Stock(id)
	if len(errs) > 0 {
		return nil,errs
	}

	errs = stokDb.conn.Delete(stk,id).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return stk,errs
}

func (stockDb *StockDb) SaveStock(stock *model.Stock)(*model.Stock,[]error)  {
	stk := stock
	errs := stockDb.conn.Create(stk).GetErrors()
	if len(errs) > 0  {
		return nil,errs
	}
	return stk,errs
}


//type StockRepository interface {
//	Stocks() ([]model.Stock,[]error)
//	Stock(Id uint) (*model.Stock,[]error)
//	UpdateStock(stock *model.Stock) (*model.Stock,[]error)
//	DeleteStock(id uint) (*model.Stock,[]error)
//	SaveStock(stock *model.Stock)(*model.Stock,[]error)
//}

