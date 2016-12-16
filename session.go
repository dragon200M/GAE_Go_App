package app

import (
	"net/http"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine"
)

const CookieName string =  "session"

func newSession(res http.ResponseWriter, req *http.Request, usr User) {
	ctx := appengine.NewContext(req)

	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  CookieName,
		Value: id.String(),
		Path:  "/",
		MaxAge: 60 * 10,

	}
	http.SetCookie(res, cookie)

	json, err := json.Marshal(usr)
	if err != nil {

		http.Error(res, err.Error(), 500)
		return
	}
	sd := memcache.Item{
		Key:   id.String(),
		Value: json,
	}

	memcache.Set(ctx, &sd)
}

func getTemplate(res http.ResponseWriter, req *http.Request, templateName string) {
	memItem, err := getSession(req)

	if err != nil {

		t.ExecuteTemplate(res, templateName, SessionData{})
		return
	}

	var sd SessionData
	json.Unmarshal(memItem.Value, &sd)
	sd.LoggedIn = true
	t.ExecuteTemplate(res, templateName, &sd)
}

func getSession(req *http.Request) (*memcache.Item, error) {
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie(CookieName)
	if err != nil {
		return &memcache.Item{}, err
	}

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		return &memcache.Item{}, err
	}
	return item, nil
}




