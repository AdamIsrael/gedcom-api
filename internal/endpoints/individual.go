package endpoints

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"path"
	"strconv"

	. "github.com/adamisrael/gedcom-api/internal/utils"

	"github.com/adamisrael/gedcom"
	"github.com/gorilla/mux"
)

type Individual struct {
	Xref     string
	Name     string
	Family   string
	Father   string
	Mother   string
	Siblings []string
}

// IndividualHandler handles requests for Individual records
func IndividualHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	gedcomid := params["gedcomid"]

	if len(gedcomid) == 0 {
		io.WriteString(w, `{"error": "Invalid Gedcom ID"}`)
	} else {

		config := GetConfig()

		filename := path.Join(config.Gedcom.Path, gedcomid+".ged")
		if !FileExists(filename) {
			fmt.Fprintln(w, `{status: "Invalid GEDCOM"}`)
			return
		}

		// Parse the GEDCOM file
		g := gedcom.Gedcom(filename)

		individual := Individual{}

		// TODO: Figure out a better way of finding an individual by id
		for _, i := range g.Individual {

			individualID, err := strconv.Atoi(params["id"])
			if CheckErr(err) {
				fmt.Printf("ID \"%s\" is not a valid integer", params["id"])
				return
			}

			// TODO: Return the matching individual
			fmt.Printf("Looking for Individual P%d\n", individualID)

			if i.Xref == fmt.Sprintf("P%d", individualID) {
				individual.Xref = i.Xref
				individual.Name = i.Name[0].Name
				individual.Family = i.Parents[0].Family.Xref
				individual.Father = i.Parents[0].Family.Husband.Xref
				individual.Mother = i.Parents[0].Family.Wife.Xref

				for _, n := range i.Parents[0].Family.Child {
					if i.Xref != n.Xref {
						individual.Siblings = append(individual.Siblings, n.Xref)
					}
				}

				individualJSON, err := json.Marshal(individual)
				CheckErr(err)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(individualJSON)
				return

			}
		}

		fmt.Fprintf(w, "Hello, individual %q", html.EscapeString(params["gedcomid"]))
	}
}
