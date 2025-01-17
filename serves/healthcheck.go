package serves

import (
	"encoding/json"
	"net/http"
)

func (a Serve) HealthCheck(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
