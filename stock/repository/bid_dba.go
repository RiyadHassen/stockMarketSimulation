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

func (bid *BidDb) Bids() ([]model.BidDb,[]error) {

	bids := []model.BidDb{}
	errs := bid.conn.Find(&bids).GetErrors()
	fmt.Println("from databasee",bids)
	if len(errs) > 0 {
		return nil,errs
	}
	return bids,errs
}


func (bid *BidDb) Bid(Id uint)(*model.BidDb,[]error) {

	oneBid := model.BidDb{}
	errs := bid.conn.First(&oneBid,Id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &oneBid,errs
}

func  (bid *BidDb) UpdateBid(ubid *model.BidDb)(*model.BidDb,[]error){
	onebid := ubid
	errs := bid.conn.Save(onebid).GetErrors()
	if len(errs) > 0 {
		return ubid,errs
	}
	return onebid,errs
}

func (bid *BidDb)  DeleteBid(Id uint)(*model.BidDb,[]error)  {
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

func (bid *BidDb) SaveBid(oneBid *model.BidDb)(*model.BidDb,[]error) {
	cat := oneBid
	errs := bid.conn.Create(cat).GetErrors()
	if len(errs) > 0  {
		return nil,errs
	}
	return cat,errs
}
