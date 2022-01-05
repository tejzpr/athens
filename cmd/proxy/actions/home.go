package actions

import (
	"net/http"
)

func proxyHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`"Go Modules Proxy"`))
}
