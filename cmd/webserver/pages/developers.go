package pages

import (
	"net/http"

	"github.com/gamedb/gamedb/pkg/helpers"
	"github.com/gamedb/gamedb/pkg/log"
	"github.com/gamedb/gamedb/pkg/sql"
	"github.com/gamedb/gamedb/pkg/tasks"
	"github.com/go-chi/chi"
)

func DevelopersRouter() http.Handler {

	r := chi.NewRouter()
	r.Get("/", developersHandler)
	return r
}

func developersHandler(w http.ResponseWriter, r *http.Request) {

	// Get config
	config, err := tasks.GetTaskConfig(tasks.Developers{})
	if err != nil {
		err = helpers.IgnoreErrors(err, sql.ErrRecordNotFound)
		log.Err(err, r)
	}

	// Get developers
	developers, err := sql.GetAllDevelopers([]string{})
	if err != nil {
		returnErrorTemplate(w, r, errorTemplate{Code: 500, Message: "There was an issue retrieving the developers.", Error: err})
		return
	}

	prices := map[int]string{}
	for _, v := range developers {
		price, err := v.GetMeanPrice(helpers.GetProductCC(r))
		log.Err(err, r)
		prices[v.ID] = price
	}

	// Template
	t := statsDevelopersTemplate{}
	t.fill(w, r, "Developers", "All the software developers that create Steam content.")
	t.Developers = developers
	t.Date = config.Value
	t.Prices = prices

	returnTemplate(w, r, "developers", t)
}

type statsDevelopersTemplate struct {
	GlobalTemplate
	Developers []sql.Developer
	Date       string
	Prices     map[int]string
}
