package app


import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

)

const categoryKey = "category"

func (cat *Category) putCategory(req *http.Request, usr *User) error {

	ctx := appengine.NewContext(req)

	//usrKey := datastore.NewKey(ctx, userKey, usr.UserName, 0, nil)

	//key := datastore.NewKey(ctx, categoryKey,cat.Name,0,usrKey)
	key:= cat.key(req,usr)

	_, err := datastore.Put(ctx, key, cat)

	return err
}


func getCategory(req *http.Request, usr *User)([]Category, error){
	ctx := appengine.NewContext(req)

	var cat []Category

	query :=datastore.NewQuery(categoryKey).Order("Name")

	usrK :=usr.key(req)

	query = query.Ancestor(usrK)

	_, err := query.GetAll(ctx, &cat)

	if err != nil{
		return nil, err
	}


	return cat , err



}

func getCategoryKeyByName(req *http.Request, usr *User, catName string)(*datastore.Key,error){

	ctx := appengine.NewContext(req)

	var cat []Category
	usrK :=usr.key(req)
	query :=datastore.NewQuery(categoryKey).Filter("Name =",catName).Ancestor(usrK)

	key, err := query.GetAll(ctx, &cat)

	//key := query.KeysOnly()

	if err != nil{
		return nil,err
	}



	return key[0],err


}

func delCategory(req *http.Request, catName string,usr *User) error{
	ctx := appengine.NewContext(req)

	var cat []Category

	query :=datastore.NewQuery(categoryKey).Filter("Name =", catName)

	usrK :=usr.key(req)
	query = query.Ancestor(usrK)

	key, err := query.GetAll(ctx, &cat)

	if err != nil{
		log.Errorf(ctx,"%v",err)
		return err
	}

	if len(key)> 0{
		err = datastore.Delete(ctx, key[0])
		return err
	}


	return nil

}
