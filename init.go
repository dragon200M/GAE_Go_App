package app

import (
	"net/http"
	"html/template"

	"github.com/julienschmidt/httprouter"
)

var t *template.Template

func init() {
	t = template.Must(t.ParseGlob("templates/*.html"))
	router := httprouter.New()
	http.Handle("/", router)

	router.GET("/", indexHandle)
	router.GET("/new/adduser", addUserForm)
	router.GET("/new/login", loginUserForm)

	router.POST("/user/create", newUser)
	router.POST("/user/login", loginUser)
	router.GET("/user/logout",logoutUser)

	router.GET("/new/category", addCategoryForm)
	router.GET("/new/expense", addExpenseForm)

	router.POST("/expense/create", newExpenses)
	router.POST("/category/create", newCategory)
	router.POST("/category/delete",deleteCategory)

	router.GET("/summary", summaryPage)
	router.POST("/user/check",userNameExists)

	router.GET("/serve/time",serveTime)


}
