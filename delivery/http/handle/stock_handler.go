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

type StockHandler struct {
	tmpl *template.Template
	stockServ stock.StockService
	//csrfSignKey []byte
}


//
func NewStockHandler(t *template.Template , stockServ stock.StockService) *StockHandler {
	return &StockHandler{tmpl:t,stockServ:stockServ}
}

func  (ach *StockHandler) Stocks(w http.ResponseWriter, r *http.Request){
	stks,errs := ach.stockServ.Stocks()

	if len(errs) > 0  {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(stks)
	ach.tmpl.ExecuteTemplate(w, "stocks.layout", stks)
}


func  (ach *StockHandler) StockNew(w http.ResponseWriter, r *http.Request) {

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
func (handler *StockHandler) Index(writer http.ResponseWriter,r *http.Request)  {
	if r.URL.Path !="/"{
		http.NotFound(writer,r)
		return
	}
	stks,errs := handler.stockServ.Stocks()
	if len(errs) > 0 {
		http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}
	handler.tmpl.ExecuteTemplate(writer,"index.layout",stks)
}

