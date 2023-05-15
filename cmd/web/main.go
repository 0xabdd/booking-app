package main

import (
	"fmt"
	"github.com/0xabdd/booking-app/pkg/config"
	"github.com/0xabdd/booking-app/pkg/handlers"
	"github.com/0xabdd/booking-app/pkg/render"
	"github.com/alexedwards/scs/v2"
	"net/http"
	"time"
)

const PORT string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT))

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}
