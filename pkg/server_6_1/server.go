package server_6_1

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Genre string   `json:"genre"`
	Songs []string `json:"songs"`
}

var artists = map[string]Artist{
	"30 seconds to Mars": {
		ID:    1,
		Name:  "30 seconds to Mars",
		Genre: "rock",
		Songs: []string{
			`The Kill`,
			`A Beautiful Lie`,
			`Attack`,
			`Live Like A Dream`,
		},
	},
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	band := r.URL.Query().Get("id")
	resp, err := json.Marshal(artists[band])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
