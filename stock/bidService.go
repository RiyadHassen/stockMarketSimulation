package stock

import "FinalDSP/model"

type BidService interface {
	Bids() ([]model.BidDb,[]error)
	Bid(Id uint)(*model.BidDb,[]error)
	UpdateBid(bid *model.BidDb)(*model.BidDb,[]error)
	DeleteBid(Id uint)(*model.BidDb,[]error)
	SaveBid(bid *model.BidDb)(*model.BidDb,[]error)
}
