package main

import (
	"FinalDSP/delivery/http/handle"
	"FinalDSP/model"
	"FinalDSP/stock/repository"
	_ "FinalDSP/stock/repository"
	"FinalDSP/stock/service"
	_"FinalDSP/stock/service"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"html/template"
	"net/http"
)
var tmpl = template.Must(template.ParseGlob("../FinalDSP/ui/web/templates/*"))
func index(writer http.ResponseWriter, request *http.Request)  {
	tmpl.ExecuteTemplate(writer,"index.layout",nil)
}

func createTables(dbconn *gorm.DB) []error {

	dbconn.SingularTable(true)
	dbconn.DropTableIfExists(&model.Stock{},&model.Category{},model.BidDb{},model.User{},model.Role{},model.Session{})
	errs := dbconn.CreateTable(&model.Stock{},&model.Category{},model.BidDb{},model.User{},model.Role{},model.Session{}).GetErrors()
	if errs != nil{
		return errs
	}
	fmt.Println("=======Database Created===========")
	return nil

}

func main()  {

	tmpl := template.Must(template.ParseGlob("../FinalDSP/ui/web/templates/*"))

	dbconn,err := gorm.Open("postgres","user=postgres dbname=stock password=root  sslmode=disable")
	dbconn.SingularTable(true)

	if err != nil{
		fmt.Println(err)
		panic(err)

	}

	//createTables(dbconn)


	defer dbconn.Close()

	stockRepo := repository.NewStockDB(dbconn)

	stockServ := service.NewStockService(stockRepo)


	categoryRepo := repository.NewCategoryDb(dbconn)

	categoryService := service.NewCategoryService(categoryRepo)

	bidRepo := repository.NewBidDb(dbconn)

	bidServ := service.NewBidService(bidRepo)

	bh := handle.NewUserBidHandler(tmpl,bidServ)


	ch := handle.NewCategoryHandler(tmpl,categoryService)

	sh := handle.NewStockHandler(tmpl,stockServ)

	fs := http.FileServer(http.Dir("../FinalDSP/ui/web/assets"))
	http.Handle("/assets/",http.StripPrefix("/assets/",fs))
	http.HandleFunc("/",sh.Index)
	http.HandleFunc("/bids", bh.UserBids )
	http.HandleFunc("/bids/new", bh.NewBid )
	http.HandleFunc("/categories",ch.Categories)
	http.HandleFunc("/categories/new",ch.CategoriesNew)
	http.HandleFunc("/stocks/new",sh.StockNew)
	http.HandleFunc("/stocks",sh.Stocks)
	http.ListenAndServe(":8080",nil)
}