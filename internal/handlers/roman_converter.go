package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"romanum/pkg/converter"
)

type ConversionResult struct {
	Number       int    `json:"number"`
	RomanNumeral string `json:"roman"`
}

func ConversionHandler(w http.ResponseWriter, req *http.Request) {
	// Get the range URL params
	lowerStr := req.URL.Query().Get("start")
	upperStr := req.URL.Query().Get("end")
	if lowerStr == "" || upperStr == "" {
		errMsg := "Params 'start' and 'end' needed"
		http.Error(w, errMsg, http.StatusBadRequest)
        return
	}
    // Cast bounds to integer
	lower, err1 := strconv.Atoi(lowerStr)
	upper, err2 := strconv.Atoi(upperStr)
	if err1 != nil || err2 != nil {
		errMsg := "Range bounds must be integers"
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	// Sanity check
    if lower < 1 || upper > 3999 || lower > upper {
		errMsg := "Invalid range. Ensure 1 <= start <= end <= 3999"
	    http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	fmt.Printf("Requested range: %d-%d\n", lower, upper)

	results := make([]ConversionResult, 0)
	for i := lower; i <= upper; i++ {
		result, err := converter.ToRoman(i)
        // Defensive catch in case of implementation errors
		if err != nil {
			errMsg := fmt.Sprintf("Unable to convert number: %d", i)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
		results = append(results, ConversionResult{
			Number:      i,
			RomanNumeral: result,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}