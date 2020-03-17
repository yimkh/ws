package probabtest

import (
	"math/rand"
	"net/http"

	file "github.com/yimkh/ws/pkg/file"
	pkgmodel "github.com/yimkh/ws/pkg/model"
)

//Ptest is probability test
func Ptest(w http.ResponseWriter) {
	flag := rand.Intn(100)

	m, _ := file.GetFileContent(pkgmodel.FilePath)

	switch {
	case flag < m["404"]:
		http.Error(w, "Not Found", 404)
		return
	case flag >= m["404"] && flag < m["404"]+m["505"]:
		http.Error(w, "HTTP Version not supported", 505)
		return
	case flag >= m["404"]+m["505"] && flag < m["404"]+m["505"]+m["400"]:
		http.Error(w, "Bad Request", 400)
		return
	default:
		return
	}
}
