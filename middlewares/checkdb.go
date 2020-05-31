package middlewares

import (
	"net/http"
	"github.com/AdrianMendez1199/red-social-go/db"
)

/* middlewares check if db is up */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		if !db.PingToDB(){
			http.Error(w, "Ocurrio un error al connectar a la bd", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}