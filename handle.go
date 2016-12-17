package app

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"google.golang.org/appengine"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"

	"google.golang.org/appengine/memcache"
	"time"
	"strings"
)

func indexHandle(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var sd SessionData

	memItem, err := getSession(req)

	if err == nil {

		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
	}
	t.ExecuteTemplate(res, "index.html", &sd)
}

func addUserForm(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	//getTemplate(res, req, "newUserForm.html")
	t.ExecuteTemplate(res, "newUserForm.html", "")
}

func loginUserForm(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	getTemplate(res, req, "loginForm.html")
}



func loginUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	key := datastore.NewKey(ctx, userKey, req.FormValue("user"), 0, nil)

	var usr User

	err := datastore.Get(ctx, key, &usr)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(req.FormValue("password"))) != nil {

		var session SessionData
		session.LoginFail = true

		t.ExecuteTemplate(res, "loginForm.html", session)
	} else {

		usr.UserName = req.FormValue("user")
		newSession(res, req, usr)

		http.Redirect(res, req, "/", 302)

	}

}

func logoutUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie(CookieName)

	if err != nil {
		http.Redirect(res, req, "/", 302)
		return
	}

	session := memcache.Item{
		Key:        cookie.Value,
		Value:      []byte(""),
		Expiration: time.Duration(1 * time.Microsecond),
	}
	memcache.Set(ctx, &session)

	cookie.MaxAge = -1
	http.SetCookie(res, cookie)

	http.Redirect(res, req, "/", 302)

}

func addCategoryForm(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)


	memitem , err := getSession(req)

	var usr User
	json.Unmarshal(memitem.Value, &usr)

	if err == nil {
		cat,_ := getCategory(req,&usr)


		t.ExecuteTemplate(res,"newCategory.html", cat)

	}

	if err != nil {
		log.Infof(ctx, "You must be logged in")
		http.Error(res, "You must be logged in", http.StatusForbidden)
		return
	}


}

func newCategory(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)

	if err != nil {
		log.Infof(ctx, "You must be logged in")
		http.Error(res, "You must be logged in", http.StatusForbidden)
		return
	}

	var usr User
	json.Unmarshal(memItem.Value, &usr)


	nameValue := req.FormValue("name")

	if len(nameValue) > 0{

		nameValue = strings.Title(nameValue)


		category := Category{
			Name: nameValue,
			Description: req.FormValue("description"),

		}

		err = putCategory(req, &usr, &category)

		if err != nil {
			log.Errorf(ctx, "category adding err: %v", err)
			http.Error(res, err.Error(), 500)
			return

		}
	}

	http.Redirect(res, req, "/new/category", 302)

}