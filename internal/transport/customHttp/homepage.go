package customHttp

import "net/http"

func (handler *HandlerHttp) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handler.notFound(w)
		return
	}
	data := map[string]interface{}{
		"Title": "Home",
	}
	handler.Render(w, r, "home.html", http.StatusOK, data)
}
