package comments

import "net/http"

func init() {
	h := NewHandler()
	http.Handle("/", h)
}
