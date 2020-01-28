package handle

import (
	"FinalDSP/model"
	"FinalDSP/stock"
	"fmt"
	"html/template"
	"net/http"
)

// ategoryHandler handles category handler admin requests
type CategoryHandler struct {
	tmpl        *template.Template
	categoryServ stock.CategoryService
}

func NewCategoryHandler(t *template.Template, cs stock.CategoryService) *CategoryHandler {
	return &CategoryHandler{tmpl: t, categoryServ:cs}
}

func (ach *CategoryHandler) Categories(w http.ResponseWriter, r *http.Request) {
	Caty, errs := ach.categoryServ.Categories()

	if len(errs) > 0  {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(Caty)
	ach.tmpl.ExecuteTemplate(w, "cate.layout", Caty)
}

// AdminCategoriesNew hanlde requests on route categories/new
func  (ach *CategoryHandler) CategoriesNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		ctg := &model.Category{}
		ctg.Name = r.FormValue("catname")

		//fmt.Println("===============",ctg,"==============")
		_, errs := ach.categoryServ.SaveCategory(ctg)

		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return

		}

		http.Redirect(w, r, "/categories",  http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "categ.new.layout", nil)
	}
}


