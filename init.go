package app

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
)

var t *template.Template
func init() {
	t = template.Must(template.New("").ParseGlob("templates/*.gohtml"))

	router := httprouter.New()
	router.GET("/", indexhandler)
	http.Handle("/", router)


}
