package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func convertHandler(w http.ResponseWriter, req *http.Request) {
	// Get the range URL params
	lowerStr := req.URL.Query().Get("start")
	upperStr := req.URL.Query().Get("end")
	if lowerStr == "" || upperStr == "" {
		http.Error(w, "Params 'start' and 'end' needed", http.StatusBadRequest)
        return
	}
    // Cast bounds to integer
	lower, err1 := strconv.Atoi(lowerStr)
	upper, err2 := strconv.Atoi(upperStr)
	if err1 != nil || err2 != nil {
		http.Error(w, "Range bounds must be integers", http.StatusBadRequest)
		return
	}
	// Sanity check
    if lower < 1 || upper > 3999 || lower > upper {
	    http.Error(w, "Invalid range. Ensure 1 <= start <= end <= 3999", http.StatusBadRequest)
		return
	}
	fmt.Printf("Bounds are '%v' and '%v'", lower, upper)
}

func main() {
	http.HandleFunc("/convert", convertHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
