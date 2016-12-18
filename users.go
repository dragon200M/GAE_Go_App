package app


import (
	"net/http"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/appengine"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
)


const userKey  ="users"

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

	//key := datastore.NewKey(ctx, userKey, usr.UserName, 0, nil)

	key:=usr.key(req)
	key, err = datastore.Put(ctx, key, &usr)

	if err != nil {
		log.Errorf(ctx, "user adding err: %v", err)
		http.Error(res, err.Error(), 500)
		return

	}

	newSession(res, req, usr)
	http.Redirect(res, req, "/", 302)

}








