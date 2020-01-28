package stock

import "FinalDSP/model"

type BidRepo interface {
	Bids() ([]model.Bid,[]error)
	Bid(Id uint)(*model.Bid,[]error)
	UpdateBid(bid *model.Bid)(*model.Bid,[]error)
	DeleteBid(Id uint)(*model.Bid,[]error)
	SaveBid(bid *model.Bid)(*model.Bid,[]error)
}