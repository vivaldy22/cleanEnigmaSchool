package tools

import "net/http"

func ReadQueryParam(p string, r *http.Request) string {
	return r.URL.Query().Get(p)
}
