package comments

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// middleware wraps an http.Handler with additional
// functionality.
type middleware func(http.Handler) http.Handler

// Comment represents a single comment.
type Comment struct {
	Key    *datastore.Key `json:"id" datastore:"-"`
	Body   string         `json:"body"`
	Author string         `json:"author"`
	URL    string         `json:"-"`
}

// Handler is the API.
type Handler struct {
	router *mux.Router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// NewHandler creates a new Handler.
func NewHandler() *Handler {
	router := mux.NewRouter()
	h := &Handler{
		router: router,
	}
	router.Handle("/comments", withContext(h.handleCommentsRead)).Methods("GET")
	router.Handle("/comments", withContext(h.handleCommentsCreate)).Methods("POST")
	return h
}

func (h *Handler) handleCommentsRead(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	urlStr := r.URL.Query().Get("url")
	if len(urlStr) == 0 {
		http.Error(w, "must provide url parameter", http.StatusBadRequest)
		return
	}
	u, err := url.Parse(urlStr)
	if len(urlStr) == 0 {
		http.Error(w, "bad url: "+err.Error(), http.StatusBadRequest)
		return
	}
	simpleURL := SimplifyURL(u)

	var cs []*Comment
	keys, err := datastore.NewQuery("Comment").Filter("URL =", simpleURL).GetAll(ctx, &cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, comment := range cs {
		comment.Key = keys[i]
	}

	err = json.NewEncoder(w).Encode(cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) handleCommentsCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	urlStr := r.URL.Query().Get("url")
	if len(urlStr) == 0 {
		http.Error(w, "must provide url parameter", http.StatusBadRequest)
		return
	}
	u, err := url.Parse(urlStr)
	if len(urlStr) == 0 {
		http.Error(w, "bad url: "+err.Error(), http.StatusBadRequest)
		return
	}
	simpleURL := SimplifyURL(u)

	var c Comment
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.URL = simpleURL

	key := datastore.NewIncompleteKey(ctx, "Comment", nil)
	key, err = datastore.Put(ctx, key, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Key = key

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// HandlerFuncContext is a HandlerFunc with context.
type HandlerFuncContext func(context.Context, http.ResponseWriter, *http.Request)

func withContext(h HandlerFuncContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		h(ctx, w, r)
	})
}
