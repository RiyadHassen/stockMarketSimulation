package service

import (
	"FinalDSP/model"
	"FinalDSP/stock"
)




type CategoryService struct {
	catRepo stock.CategoryRepository

}


//constructer
func NewCategoryService(cateRepo stock.CategoryRepository) *CategoryService {
	return &CategoryService{catRepo: cateRepo}
}

func (caty *CategoryService) Categories() ([]model.Category,[]error) {
	categories , errs := caty.catRepo.Categories()
	if len(errs) > 0 {
		return nil,errs
	}
	return categories,errs
}

func (caty *CategoryService) Category(Id uint)(*model.Category,[]error) {
	categorc, errs := caty.catRepo.Category(Id)
	if len(errs) > 0 {
		return nil, errs
	}
	return categorc,errs
}

func (caty *CategoryService)  UpdateCategory(category *model.Category)(*model.Category,[]error) {
	cat ,errs := caty.catRepo.UpdateCategory(category)
	if len(errs) > 0 {
		return category,errs
	}
	return cat,errs
}

func (caty *CategoryService)  DeleteCategory(Id uint)(*model.Category,[]error)  {
	cat,errs := caty.catRepo.DeleteCategory(Id)
	if len(errs) > 0{
		return nil,errs
	}
	return cat,errs
}

func (caty *CategoryService) SaveCategory(category *model.Category)(*model.Category,[]error) {
	cat , errs := caty.catRepo.SaveCategory(category)
	if len(errs) > 0  {
		return nil,errs
	}
	return cat,errs
}

func (caty *CategoryService) StockCategory(category *model.Category)([]model.Stock,[]error){
	stocks,errs := caty.catRepo.StockCategory(category)
	if len(errs) > 0  {
		return nil,errs
	}
	return stocks,errs
}
//type CategoryService interface {
//	Categories() ([]model.Category,[]error)
//	Category(Id uint)(*model.Category,[]error)
//	UpdateCategory(category *model.Category)(*model.Category,[]error)
//	DeleteCategory(Id uint)(*model.Category,[]error)
//	SaveCategory(category *model.Category)(*model.Category,[]error)
//	StockCategory(category *model.Category)([]model.Stock,[]error)
//}