package handle

import (
	"FinalDSP/stock"
	"html/template"
	"net/http"
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

func (handler *StockHandler) Index(writer http.ResponseWriter,r *http.Request)  {
	if r.URL.Path !="/"{
		http.NotFound(writer,r)
		return
	}

}
