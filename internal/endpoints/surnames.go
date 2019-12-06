package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strings"

	. "github.com/adamisrael/gedcom-api/internal/utils"

	"github.com/adamisrael/gedcom"
	"github.com/gorilla/mux"
)

// SurnameHandler returns a list of surnames and how frequently they occur
func SurnameHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	gedcomid := params["gedcomid"]

	// Specify that this is JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(gedcomid) == 0 {
		io.WriteString(w, `{"error": "Invalid Gedcom ID"}`)
	} else {

		config := GetConfig()

		filename := path.Join(config.Gedcom.Path, gedcomid+".ged")
		if !FileExists(filename) {
			io.WriteString(w, `{"error": "Gedcom not found."}`)
			return
		}

		// Parse the GEDCOM file
		g := gedcom.Gedcom(filename)

		surnames := make(map[string]int)

		for _, i := range g.Individual {
			for _, n := range i.Name {
				surname := strings.TrimSpace(n.Surname)
				if len(surname) > 0 {
					surnames[surname]++
				}
			}
		}

		surnamesJSON, err := json.Marshal(surnames)
		CheckErr(err)

		w.Write(surnamesJSON)
		// io.WriteString(w, `{"alive": true}`)

	}

}
