package app


import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

)

const categoryKey = "category"

func putCategory(req *http.Request, usr *User, cat *Category) error {

	ctx := appengine.NewContext(req)

	usrKey := datastore.NewKey(ctx, userKey, usr.UserName, 0, nil)

	key := datastore.NewIncompleteKey(ctx, categoryKey, usrKey)
	_, err := datastore.Put(ctx, key, cat)

	return err
}


func getCategory(req *http.Request, usr *User)([]Category, error){
	ctx := appengine.NewContext(req)

	var cat []Category

	query :=datastore.NewQuery(categoryKey)



		usrK :=datastore.NewKey(ctx, userKey,usr.UserName,0,nil)
		query = query.Ancestor(usrK)

	_, err := query.GetAll(ctx, &cat)


	return cat , err



}
