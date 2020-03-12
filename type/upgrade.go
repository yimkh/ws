package upgrade

import (
	"net/http"
	"strings"

	"github.com/yimkh/fake/net/http/httpguts"
)

// UpgradeType is upgrade type
func UpgradeType(h http.Header) string {
	if !httpguts.HeaderValuesContainsToken(h["Connection"], "Upgrade") {
		return ""
	}
	return strings.ToLower(h.Get("Upgrade"))
}
