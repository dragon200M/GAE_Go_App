package app

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"google.golang.org/appengine"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
	"fmt"
)

func indexHandle(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var sd SessionData

	memItem, err := getSession(req)

	if err == nil {

		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
	}
	t.ExecuteTemplate(res, "index", &sd)
}

func addUserForm(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	getTemplate(res, req, "createForm")
}


func loginUserForm(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	getTemplate(res, req, "loginForm")
}

func newUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	passHash, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.MinCost)

	if err != nil {
		log.Errorf(ctx, "password err: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}

	usr := User{
		Email: req.FormValue("email"),
		UserName: req.FormValue("user"),
		Password: string(passHash),

	}

	key := datastore.NewKey(ctx, "user", usr.UserName, 0, nil)

	key, err = datastore.Put(ctx, key, &usr)

	if err != nil {
		log.Errorf(ctx, "user adding err: %v", err)
		http.Error(res, err.Error(), 500)
		return

	}




	newSession(res, req, usr)
	http.Redirect(res, req, "/", 302)


}




func loginUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	key :=datastore.NewKey(ctx,"user",req.FormValue("user"),0,nil)

	var usr User

	err := datastore.Get(ctx,key,&usr)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(usr.Password),[]byte(req.FormValue("password"))) != nil{

		var session SessionData
		session.LoginFail = true

		t.ExecuteTemplate(res, "loginForm", session)
	}else{

		usr.UserName =req.FormValue("user")
		newSession(res,req,usr)

		http.Redirect(res, req, "/", 302)

	}



}