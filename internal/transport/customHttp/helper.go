package customHttp

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (handler *HandlerHttp) serverError(w http.ResponseWriter, err error) {
	ErrorMessage := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	handler.ErrorLog.Print(ErrorMessage)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (handler *HandlerHttp) clientError(w http.ResponseWriter, statusCode int) {
	handler.ErrorLog.Print(http.StatusText(statusCode))
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (handler *HandlerHttp) notFound(w http.ResponseWriter) {
	handler.clientError(w, http.StatusNotFound)
}

// Render need TemplateData struct(DTO) to execute

func (handler *HandlerHttp) Render(page string, w http.ResponseWriter, r http.Request, status int, data map[string]interface{}) {
	ts, ok := handler.TemplateCache.templates[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		handler.serverError(w, err)
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		handler.serverError(w, err)
	}
}
