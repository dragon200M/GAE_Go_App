package app

import (
	"net/http"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/appengine"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
	"io/ioutil"
	"fmt"
)

const userKey = "users"

func newUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	pass := req.FormValue("password")

	if len(pass) > 0 {
		passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)

		if err != nil {
			log.Errorf(ctx, "password err: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}

		usrName := req.FormValue("user")

		if userExists(req, usrName) {
			http.Redirect(res, req, "/new/adduser", 302)
			return
		}

		usr := User{
			Email: req.FormValue("email"),
			UserName: usrName,
			Password: string(passHash),

		}

		//key := datastore.NewKey(ctx, userKey, usr.UserName, 0, nil)

		key := usr.key(req)
		key, err = datastore.Put(ctx, key, &usr)

		if err != nil {
			log.Errorf(ctx, "user adding err: %v", err)
			http.Error(res, err.Error(), 500)
			return

		}

		newSession(res, req, usr)
		http.Redirect(res, req, "/", 302)
	}
	http.Redirect(res, req, "/new/adduser", 302)
}

func userNameExists(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	u, _ := ioutil.ReadAll(req.Body)
	ub := string(u)

	b := userExists(req, ub)

	if b {
		fmt.Fprint(res, "true")
	} else {
		fmt.Fprint(res, "false")
	}

}

func userExists(req *http.Request, usrName string) bool {

	ctx := appengine.NewContext(req)

	var usr User

	k := datastore.NewKey(ctx, userKey, usrName, 0, nil)
	err := datastore.Get(ctx, k, &usr)

	if err != nil {
		//no user name in datastore

		return false
	} else {
		//exists
		return true
	}
	return true
}









