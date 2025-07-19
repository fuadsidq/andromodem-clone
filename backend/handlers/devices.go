package handlers

import (
	"encoding/json"
	"net/http"
)

type Device struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Signal  int    `json:"signal"`
	Battery int    `json:"battery"`
}

func DevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices := []Device{
		{ID: "modem1", Name: "Modem Android A", Signal: 80, Battery: 90},
		{ID: "modem2", Name: "Modem Android B", Signal: 60, Battery: 40},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}
