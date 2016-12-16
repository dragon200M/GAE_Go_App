package app

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"google.golang.org/appengine"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
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