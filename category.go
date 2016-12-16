package app


import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

)

func putCategory(req *http.Request, usr *User, cat *Category) error {

	ctx := appengine.NewContext(req)

	usrKey := datastore.NewKey(ctx, "users", usr.UserName, 0, nil)

	key := datastore.NewIncompleteKey(ctx, "category", usrKey)
	_, err := datastore.Put(ctx, key, cat)

	return err
}
