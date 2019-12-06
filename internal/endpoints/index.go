package endpoints

import (
	"fmt"
	"net/http"
)

type Endpoint struct {
	Endpoint string
}

type IndexPageData struct {
	PageTitle string
	Endpoints []Endpoint
}

// IndexHandler displays a welcome page and a list of available endpoints
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: Reimplement this self-documentation endpoint
	// router := router.getRouter()
	// tmpl := template.Must(template.ParseFiles("templates/routes.html"))

	// data := IndexPageData{
	// 	PageTitle: "Welcome to the gedcom-api microservice!",
	// 	Endpoints: []Endpoint{},
	// }

	// router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	t, err := route.GetPathTemplate()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// TODO: Format the return value better
	// 	// fmt.Fprintf(w, "%q\n", html.EscapeString(t))
	// 	endpoint := Endpoint{
	// 		Endpoint: t,
	// 	}

	// 	data.Endpoints = append(data.Endpoints, endpoint)
	// 	return nil
	// })
	// tmpl.Execute(w, data)

	fmt.Fprintln(w, "Hello")

}
