package route

import (
	"bitrixapi/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func Handlers() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Use(CommonMiddleWare)

	router.HandleFunc("/", controller.Test).Methods("GET")

	router.HandleFunc("/checkPost", controller.InsertCust).Methods("POST")


	return router
}


func CommonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header, x-access-token")
		next.ServeHTTP(res, req)
	})
}