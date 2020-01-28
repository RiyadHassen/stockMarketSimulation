package repository

import (
	"FinalDSP/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type BidDb struct {
conn *gorm.DB
}

func NewBidDb(db *gorm.DB) *BidDb {
	return &BidDb{db}
}

func (bid *BidDb) Bids() ([]model.Bid,[]error) {

	bids := []model.Bid{}
	errs := bid.conn.Find(&bid).GetErrors()
	fmt.Println("from databasee",bids)
	if len(errs) > 0 {
		return nil,errs
	}
	return bids,errs
}


func (bid *BidDb) Bid(Id uint)(*model.Bid,[]error) {

	oneBid := model.Bid{}
	errs := bid.conn.First(&oneBid,Id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &oneBid,errs
}

func  (bid *BidDb) UpdateBid(ubid *model.Bid)(*model.Bid,[]error){
	onebid := ubid
	errs := bid.conn.Save(onebid).GetErrors()
	if len(errs) > 0 {
		return ubid,errs
	}
	return onebid,errs
}

func (bid *BidDb)  DeleteBid(Id uint)(*model.Bid,[]error)  {
	oneBid,errs := bid.Bid(Id)

	if len(errs) > 0 {
		return nil,errs
	}

	errs = bid.conn.Delete(oneBid,Id).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return oneBid,errs
}

func (bid *BidDb) SaveBid(oneBid *model.Bid)(*model.Bid,[]error) {
	cat := oneBid
	errs := bid.conn.Create(cat).GetErrors()
	if len(errs) > 0  {
		return nil,errs
	}
	return cat,errs
}
