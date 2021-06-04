package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

func Recovery(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("panic occurred:", r)
				msg, _ := json.Marshal(map[string]string{"Message": "Internal Error"})
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(msg)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
