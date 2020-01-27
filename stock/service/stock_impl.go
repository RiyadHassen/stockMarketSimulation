package service

import (
	"FinalDSP/model"
	"FinalDSP/stock"
)

type StockService struct {
	stockRepo stock.StockRepository

}


//constructer
func NewStockService(stockRepo stock.StockRepository) *StockService {
	return &StockService{stockRepo:stockRepo}
}

func (stockServ *StockService) Stocks()  ([]model.Stock,[]error){
	stks, errs := stockServ.stockRepo.Stocks()
	if len(errs) > 0 {
		return nil,errs
	}
	return stks,errs
}

func (stockServ *StockService)Stock(Id uint) (*model.Stock,[]error)  {
	stk,errs := stockServ.stockRepo.Stock(Id)
	if len(errs) > 0 {
		return nil,errs
	}
	return stk,errs
}

func (stockServ *StockService) UpdateStock(stock *model.Stock) (*model.Stock,[]error) {
	stk ,errs := stockServ.stockRepo.UpdateStock(stock)
	if len(errs) > 0 {
		return nil,errs
	}
	return stk,errs
}
func (stockServ *StockService) DeleteStock(id uint) (*model.Stock,[]error) {
	stk ,errs := stockServ.stockRepo.DeleteStock(id)
	if len(errs) > 0 {
		return nil,errs
	}
	return stk,errs
}

func (stockServ *StockService) SaveStock(stock *model.Stock)(*model.Stock,[]error) {
	stk ,errs := stockServ.stockRepo.SaveStock(stock)
	if len(errs) > 0 {
		return nil,errs
	}
	return stk,errs
}
