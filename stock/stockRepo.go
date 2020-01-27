package stock

import "FinalDSP/model"

type StockRepository interface {
	Stocks() ([]model.Stock,[]error)
	Stock(Id uint) (*model.Stock,[]error)
	UpdateStock(stock *model.Stock) (*model.Stock,[]error)
	DeleteStock(id uint) (*model.Stock,[]error)
	SaveStock(stock *model.Stock)(*model.Stock,[]error)
}


