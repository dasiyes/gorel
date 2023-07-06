package tools

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./src/html"),
	jet.InDevelopmentMode(),
)

// RenderPage is a helper function used to render the HTML templates, passing in any dynamic data
func RenderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	v, err := views.GetTemplate(tmpl)
	if err != nil {
		return err
	}

	err = v.Execute(w, data, nil)
	if err != nil {
		return err
	}

	return nil
}
