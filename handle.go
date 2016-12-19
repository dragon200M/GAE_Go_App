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
	"strconv"
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
	var catdata CategoryData
	var usr User
	memitem , err := getSession(req)


	json.Unmarshal(memitem.Value, &usr)

	if err == nil {
		cat,_ := getCategory(req,&usr)

		catdata.User = 	usr
		catdata.Categories = cat

		t.ExecuteTemplate(res,"newCategory.html", catdata)

	}

	if err != nil {
		log.Infof(ctx, "You must be logged in")
		//http.Error(res, "You must be logged in", http.StatusForbidden)
		http.Redirect(res, req, "/new/login", 302)
		return
	}


}

func newCategory(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)

	if err != nil {
		log.Infof(ctx, "You must be logged in")


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

		err = category.putCategory(req,&usr)


		if err != nil {
			log.Errorf(ctx, "category adding err: %v", err)
			http.Error(res, err.Error(), 500)
			return

		}



	}

	http.Redirect(res, req, "/new/category", 302)

}


func deleteCategory(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)

	if err != nil {
		log.Infof(ctx, "You must be logged in")
		http.Error(res, "You must be logged in", http.StatusForbidden)

		return
	}



	var usr User
	json.Unmarshal(memItem.Value, &usr)


	nameValue := req.FormValue("delname")

	if len(nameValue) > 0{

		nameValue = strings.Title(nameValue)

		err2 := delCategory(req,nameValue,&usr)

		if err2 != nil {
			log.Errorf(ctx, "category remove err: %v", err)
			http.Error(res, err.Error(), 500)
			http.Redirect(res, req, "/new/category", 302)
			return

		}
	}

	http.Redirect(res, req, "/new/category", 302)
}


func newExpenses(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)

	if err != nil {
		log.Infof(ctx, "You must be logged in")


		return
	}



	var usr User
	json.Unmarshal(memItem.Value, &usr)



	amLen := req.FormValue("amount")
	catLen := req.FormValue("category")

	if len(amLen) > 0 && len(catLen) > 0{





		amount ,err := strconv.ParseFloat(req.FormValue("amount"),64)

		if err != nil{
			var expData ExpensesData
			expData.Categories, _ = getCategory(req,&usr)
			expData.BadNumberFormat = "Zły format kwoty"
			expData.Expenses, _= getExpenses(req,&usr)
			t.ExecuteTemplate(res,"newExpense.html",expData)
			return
		}


		exp := Expenses{
			Category: req.FormValue("category"),
			Amount: amount,
			Description: req.FormValue("desc"),
			Month: time.Now().Month(),
			Date: time.Now(),
		}

		err = exp.putExpenses(req,exp.Category,&usr)


		if err != nil {
			log.Errorf(ctx, "expense adding err: %v", err)
			http.Error(res, err.Error(), 500)
			return

		}



	}

	http.Redirect(res, req, "/new/expense", 302)



}



func addExpenseForm(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	var expData ExpensesData
	var usr User


	memitem , err := getSession(req)



	json.Unmarshal(memitem.Value, &usr)


	expData.User = usr
	expData.Categories,_= getCategory(req,&usr)
	expData.Expenses ,_ = getExpenses(req,&usr)
	if err == nil {


		t.ExecuteTemplate(res,"newExpense.html", expData)

	}

	if err != nil {
		log.Infof(ctx, "You must be logged in")
		//http.Error(res, "You must be logged in", http.StatusForbidden)
		http.Redirect(res, req, "/new/login", 302)
		return
	}

}




func summaryPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	var usr User
	var sumData SummaryData
	memitem , err := getSession(req)


	if err != nil {
		log.Infof(ctx, "You must be logged in")
		//http.Error(res, "You must be logged in", http.StatusForbidden)
		http.Redirect(res, req, "/new/login", 302)
		return
	}

	json.Unmarshal(memitem.Value, &usr)

	s :=summary(req,&usr)
	sumData.Summary = s
	sumData.User = usr
	sumData.Len = len(s)



		t.ExecuteTemplate(res,"summary.html",sumData)





}

