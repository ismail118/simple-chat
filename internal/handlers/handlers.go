package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
