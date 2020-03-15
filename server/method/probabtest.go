package probabtest

import (
	"math/rand"
	"net/http"
)

//Ptest is probability test
func Ptest(w http.ResponseWriter) {
	flag := rand.Intn(100)

	switch {
	case flag < 30:
		http.Error(w, "Not Found", 404)
		return
	case flag >= 20 && flag < 40:
		http.Error(w, "HTTP Version not supported", 505)
		return
	case flag >= 40 && < 60:
		http.Error(w, "Bad Request", 400)
		return
	default:
		return
	}
}
