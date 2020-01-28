package handle

import (
	"FinalDSP/model"
	"FinalDSP/stock"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type UesrBidHandler struct {
	tmpl        *template.Template
	bidserv      stock.BidService
}

func NewUserBidHandler(t *template.Template, ub stock.BidService)  *UesrBidHandler{
	return &UesrBidHandler{t,ub}
}

func (bidhdl *UesrBidHandler) UserBids(w http.ResponseWriter,r *http.Request){
	bids,errs := bidhdl.bidserv.Bids()
	fmt.Println(errs)

	if len(errs) > 0  {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(bids)
	bidhdl.tmpl.ExecuteTemplate(w, "usr.bid.layout", bids)
}
func (bidhdl *UesrBidHandler) NewBid(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodPost {

		//idRaw := r.URL.Query().Get("id")
		//id, err := strconv.Atoi(idRaw)

		postBid := model.BidDb{}

		postBid.Name = r.FormValue("bidname")

		// not accepted
		postBid.Status  = false

		postBid.StockID = uint(2)

		price := r.FormValue("price")

		postBid.Price, _ = strconv.ParseFloat(price,12)

		_, errs := bidhdl.bidserv.SaveBid(&postBid)

		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return

		}
		fmt.Println("still good")

		http.Redirect(w, r, "/bids", http.StatusSeeOther)

	} else{
		bidhdl.tmpl.ExecuteTemplate(w,"new.bid.layout", nil)
	}
}
