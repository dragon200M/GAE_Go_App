package app

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"


)

const expensesKey = "expenses"

func (exp *Expenses) putExpenses(req *http.Request, cat string, usr *User) error {

	ctx := appengine.NewContext(req)



	_, err := datastore.Put(ctx, exp.key(req,cat,usr), exp)

	return err
}


