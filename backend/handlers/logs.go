package handlers

import (
	"encoding/json"
	"net/http"
)

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	logs := []string{
		"[10:01] Modem A: Signal OK",
		"[10:03] Modem B: Battery low",
		"[10:05] Modem A: Restart triggered",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
