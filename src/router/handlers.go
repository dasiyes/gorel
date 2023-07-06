package router

import (
	"net/http"

	"github.com/dasiyes/gorel/src/tools"
)

func getHome() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := tools.RenderPage(w, "home.jet", nil)
		if err != nil {
			http.Error(w, "Error rendering home page", http.StatusInternalServerError)
		}
	})
}
