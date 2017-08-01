package user

import "net/http"

func Router(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("hello wego"))
}