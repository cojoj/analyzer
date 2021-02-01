package handler

import (
	"net/http"
	"net/url"

	"github.com/cojoj/analyzer/config"
	"github.com/cojoj/analyzer/internal/website"
)

// Index executes index.gohtml template.
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	config.TPL.ExecuteTemplate(w, "index.gohtml", nil)
}

// Analyze handles `/analyze` POST request. It parses passed URL from the form, fetches website's
// HTML code and performs analysis. When everything finishes successfully it executes Report.
func Analyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	websiteUrl, err := url.Parse(r.FormValue("url"))
	if err != nil {
		config.TPL.ExecuteTemplate(w, "index.gohtml", err)
		return
	}

	rc, err := website.Fetch(websiteUrl.String())
	if err != nil {
		config.TPL.ExecuteTemplate(w, "index.gohtml", err)
		return
	}
	defer rc.Close()

	rep, err := website.Analyze(websiteUrl, rc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "report.gohtml", rep)
}
