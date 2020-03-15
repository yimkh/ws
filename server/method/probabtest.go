package probabtest

import (
	"math/rand"
	"net/http"
)

//Ptest is probability test
func Ptest(w http.ResponseWriter) {
	flag := rand.Intn(70)

	switch {
	case flag < 30:
		http.Error(w, "", 404)
		return
	case flag >= 20 && flag < 40:
		http.Error(w, "", 505)
		return
	default:
		return
	}
}
