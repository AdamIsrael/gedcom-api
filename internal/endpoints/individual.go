package endpoints

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/adamisrael/gedcom-api/internal/utils"

	"github.com/adamisrael/gedcom"
	"github.com/gorilla/mux"
)

// Individual is the structure to use for the JSON output
type Individual struct {
	Xref     string
	Name     string
	Family   string
	Father   string
	Mother   string
	Siblings []string
}

// AllIndividualHandler handles requests for all Individual records
func AllIndividualHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write(json)
	fmt.Fprintln(w, "{}")
}

// IndividualHandler handles requests for Individual records
func IndividualHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Get the GEDCOM id
	gedcomid := params["gedcomid"]

	// Get the Individual ID
	var individualid int
	var err error

	individualid, err = strconv.Atoi(params["id"])
	if utils.CheckErr(err) {
		fmt.Printf("ID \"%s\" is not a valid integer", params["id"])
		return
	}

	// If no individual id is zero, pass this to the AllIndividualHandler
	if individualid == 0 {
		AllIndividualHandler(w, r)
		return
	}

	if len(gedcomid) == 0 {
		io.WriteString(w, `{"error": "Invalid Gedcom ID"}`)
	} else {

		config := utils.GetConfig()

		filename := path.Join(config.Gedcom.Path, gedcomid+".ged")
		if !utils.FileExists(filename) {
			fmt.Fprintln(w, `{status: "Invalid GEDCOM"}`)
			return
		}

		// Parse the GEDCOM file
		g := gedcom.Gedcom(filename)

		individual := Individual{}

		// TODO: Figure out a better way of finding an individual by id
		for _, i := range g.Individual {

			// TODO: Return the matching individual
			// fmt.Printf("Looking for Individual P%d\n", individualid)

			if i.Xref == fmt.Sprintf("P%d", individualid) {
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
				utils.CheckErr(err)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(individualJSON)
				return

			}
		}

		fmt.Fprintf(w, "Hello, individual %q", html.EscapeString(params["gedcomid"]))
	}
}
