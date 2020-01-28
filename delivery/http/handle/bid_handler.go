package handle

import (
	"FinalDSP/model"
	"FinalDSP/stock"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type BidHandler struct {
	tmpl *template.Template
	bidSer stock.BidService
	//csrfSignKey []byte
}


//
func NewBidHandler(t *template.Template , bidServ stock.BidService) *BidHandler {
	return &BidHandler{bidSer:bidServ}
}

func  (ach *BidHandler) listBids(w http.ResponseWriter, r *http.Request){
	bids,errs := ach.bidSer.Bids()

	if len(errs) > 0  {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(bids)
	ach.tmpl.ExecuteTemplate(w, "stocks.layout", bids)
}


func  (ach *StockHandler) bidNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		stk := &model.Stock{}
		stk.Name = r.FormValue("stkname")
		stk.Desc = r.FormValue("stkdesc")
		maxValue := r.FormValue("max_value")
		stk.MaxValue, _ = strconv.ParseFloat(maxValue,5)
		minvalue := r.FormValue("min_value")
		stk.MinValue, _ = strconv.ParseFloat(minvalue,5)

		stk.StartTime = time.Now()

		stk.Turnoff = false



		fmt.Println("===============",stk,"==============")
		_, errs := ach.stockServ.SaveStock(stk)

		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return

		}

		http.Redirect(w, r, "/stocks",  http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "stock.new.layout", nil)
	}
}


