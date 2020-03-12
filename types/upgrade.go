package upgrade

import (
	"net/http"
	"strings"

	"golang.org/x/net/http/httpguts"
)

// UpgradeType is upgrade type
func UpgradeType(h http.Header) string {
	if !httpguts.HeaderValuesContainsToken(h["Connection"], "Upgrade") {
		return ""
	}
	return strings.ToLower(h.Get("Upgrade"))
}
