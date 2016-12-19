package app


import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"time"

)

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
}





//User
type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}


func (usr *User) key(req *http.Request) *datastore.Key{

	ctx := appengine.NewContext(req)

	key := datastore.NewKey(ctx, userKey,usr.UserName,0,nil)


	return key

}
//end User





//Category
type Category struct {
	Name        string
	Description string

}


func (cat *Category) key(req *http.Request, usr *User) *datastore.Key{

	ctx := appengine.NewContext(req)

	key := datastore.NewKey(ctx, categoryKey,cat.Name,0,usr.key(req))


	return key

}

//end Category



type CategoryData struct {
	User
	Categories []Category
}



//Expenses

type Expenses struct {
	Category string
	Amount float64
	Description string
	Month time.Month
	Date time.Time

}


func (exp *Expenses) key(req *http.Request, cat string, usr *User) *datastore.Key{
	ctx := appengine.NewContext(req)

	//catKey := cat.key(req,usr)
	catKey,_ := getCategoryKeyByName(req,usr,cat)

	key:= datastore.NewIncompleteKey(ctx,expensesKey,catKey)

	return key
}
//end Expenses


type ExpensesData struct{
	User
	Categories 	[]Category
	Expenses 	[]Expenses
	BadNumberFormat string
}


type Summary struct{
	Month string
	CatSum []CatSum
	MonthSum string
}

type CatSum struct {
	Name string
	Sum string
}

type SummaryData struct {
	User
	Summary []Summary
	Len int
}