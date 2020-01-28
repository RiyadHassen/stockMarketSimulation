package service

import (
	"FinalDSP/model"
	"FinalDSP/stock"
	"fmt"
)

type BidService struct {
	bidRepo stock.BidRepo

}


//UpdateBid(bid *model.Bid)(*model.Bid,[]error)
//DeleteBid(Id uint)(*model.Bid,[]error)
//SaveBid(bid *model.Bid)(*model.Bid,[]error)

//constructer
func NewBidService(repo stock.BidRepo ) *BidService {
	return &BidService{bidRepo:repo}
}

func (bidServ *BidService)  Bids() ([]model.BidDb,[]error){
	bides , errs := bidServ.bidRepo.Bids()
	fmt.Println(bides)
	if len(errs) > 0 {
		return nil,errs
	}
	return bides,errs
}

func  (bidServ *BidService)  Bid(Id uint)(*model.BidDb,[]error) {
	oneBid, errs := bidServ.bidRepo.Bid(Id)
	if len(errs) > 0 {
		return nil, errs
	}
	return oneBid,errs
}

func (bidServ *BidService)   UpdateBid(bid *model.BidDb)(*model.BidDb,[]error) {
	oneBid ,errs := bidServ.bidRepo.UpdateBid(bid)
	if len(errs) > 0 {
		return bid,errs
	}
	return oneBid,errs
}

func (bidServ *BidService)  DeleteBid(Id uint)(*model.BidDb,[]error) {
	cat,errs := bidServ.bidRepo.DeleteBid(Id)
	if len(errs) > 0{
		return nil,errs
	}
	return cat,errs
}

func (bidServ *BidService) SaveBid(bid *model.BidDb)(*model.BidDb,[]error){
	onebid , errs := bidServ.bidRepo.SaveBid(bid)
	if len(errs) > 0  {
		return nil,errs
	}
	return onebid,errs
}

