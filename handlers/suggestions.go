package handlers

import (
	groupie "groupie/data"
	"net/http"
	"strings"
	"encoding/json"
)

// SuggestionHandler handles suggestions for atists based on a query parameter
func SuggestionHandler(w http.ResponseWriter, r *http.Request) {
	var results []groupie.Artist
	word := r.URL.Query().Get("query")

	for _, artist := range Artiste {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(word)) {
			results = append(results, artist)
		}
	}

	// Return the results as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
