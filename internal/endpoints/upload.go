package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/adamisrael/gedcom-api/internal/utils"
	"github.com/google/uuid"
)

// GetGedcomFilename returns the path to use to store an uploaded GEDCOM file
func getGedcomFilename() (string, string) {
	id, err := uuid.NewUUID()
	utils.CheckErr(err)

	config := utils.GetConfig()
	filename := path.Join(config.Gedcom.Path, id.String()+".ged")
	if utils.FileExists(filename) {
		return getGedcomFilename()
	}
	return id.String(), filename
}

// UploadHandler handles uploads of GEDCOM files
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: Move hardcoded values to a config
	// TODO: Create a yaml file with every upload to capture metadata: timestamp, ??

	config := utils.GetConfig()

	// Parse the multipart form, specifying a max upload of 20 MB
	r.ParseMultipartForm(config.Gedcom.MaxFileSize << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("gedcom")
	if utils.CheckErr(err) {
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	uuid, filename := getGedcomFilename()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	utils.CheckErr(err)
	// write this byte array to our temporary file
	err = ioutil.WriteFile(filename, fileBytes, 0644)
	utils.CheckErr(err)
	// return that we have successfully uploaded our file!

	output := map[string]string{
		"uuid":   uuid,
		"status": "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonString, err := json.Marshal(output)
	fmt.Fprintf(w, string(jsonString))

}
