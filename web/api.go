package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type status struct {
	ID          int `json:"ID"`
	Title       string `json:"Title"`
	Good bool `json:"good"`
}

type allChecks []status

func apiHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API home was accessed")
}

func getCheckStatus(w http.ResponseWriter, r *http.Request) {
	checkID := mux.Vars(r)["check"]

	raw, err = ioutil.ReadFile("/opt/gingertechengine/js/current.json")
	if err != nil {
		helpers.Log.Println("Previous check file could not be read, error return:")
		helpers.Log.Println(err.Error())
		os.Exit(1)
	}

	var checks = json.Unmarshal(raw, &allChecks)

	for _, singleCheck := range allChecks {
		if singleCheck.ID == checkID {
			json.NewEncoder(w).Encode(singleCheck)
		}
	}
}

func getAllChecks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("/opt/gogios/js/current.json")
}
