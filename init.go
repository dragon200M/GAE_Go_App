package app

import (
	"net/http"
	"html/template"

	"github.com/julienschmidt/httprouter"
)

var t *template.Template
func init() {
	t = template.Must(template.New("").ParseGlob("templates/*.html"))
	router := httprouter.New()
	http.Handle("/", router)


	router.GET("/",indexHandle)
	router.GET("/new/adduser",addUserForm)
	router.POST("/user/create", newUser)



}
