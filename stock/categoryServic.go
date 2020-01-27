package stock

import "FinalDSP/model"

type CategoryService interface {
	Categories() ([]model.Category,[]error)
	Category(Id uint)(*model.Category,[]error)
	UpdateCategory(category *model.Category)(*model.Category,[]error)
	DeleteCategory(Id uint)(*model.Category,[]error)
	SaveCategory(category *model.Category)(*model.Category,[]error)
	StockCategory(category *model.Category)([]model.Stock,[]error)
}