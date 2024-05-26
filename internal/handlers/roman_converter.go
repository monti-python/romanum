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
	RomanNumeral string `json:"roman_numeral"`
}

func ConversionHandler(w http.ResponseWriter, req *http.Request) {
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

	results := make([]ConversionResult, 0)
	for i := lower; i <= upper; i++ {
		result, _ := converter.ToRoman(i)
		results = append(results, ConversionResult{
			Number:      i,
			RomanNumeral: result,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}