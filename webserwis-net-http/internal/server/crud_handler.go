package server

import (
	"encoding/json"
	"net/http"
	"webserwis/internal/model"
	log "webserwis/utils/log"
)

// SendItem godoc
//
//	@Summary		Submit user data
//	@Description	Receives user data via POST request and returns a success message
//	@Tags			crud
//	@Accept			json
//	@Produce		json
//	@Param			requestBody	body		model.RequestBody	true	"User data payload"
//	@Success		200			{object}	model.ResponseBody
//	@Failure		400			{string}	string	"Failed to parse JSON"
//	@Failure		405			{string}	string	"Invalid request method"
//	@Router			/crud/ [post]
func (s *Server) SendItem(w http.ResponseWriter, r *http.Request) {

	// Dekodowanie JSON-a z żądania
	var requestBody model.RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Zbudowanie odpowiedzi
	response := model.ResponseBody{
		Message: "Data SEND successfully",
		Status:  "success",
	}

	// Zwrócenie odpowiedzi w formacie JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Failed to write response:", err)
	}
}

// UpdateItem godoc
//
//	@Summary		Update user data
//	@Description	Receives user data via PUT request and returns a success message
//	@Tags			crud
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Item Id"
//	@Param			requestBody	body		model.RequestBody	true	"User data payload"
//	@Success		200			{object}	model.ResponseBody
//	@Failure		400			{string}	string	"Failed to parse JSON"
//	@Failure		405			{string}	string	"Invalid request method"
//	@Router			/crud/{id} [put]
func (s *Server) UpdateItem(w http.ResponseWriter, r *http.Request) {

	// Dekodowanie JSON-a z żądania
	var requestBody model.RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	// Zbudowanie odpowiedzi
	response := model.ResponseBody{
		Message: "Data UPDATED successfully for id: " + id,
		Status:  "success",
	}

	// Zwrócenie odpowiedzi w formacie JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Failed to write response:", err)
	}
}

// DeleteItem godoc
//
//	@Summary		Submit user data
//	@Description	Receives user data via POST request and returns a success message2
//	@Tags			crud
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Item Id"
//	@Success		200			{object}	model.ResponseBody
//	@Failure		400			{string}	string	"Failed to parse JSON"
//	@Failure		405			{string}	string	"Invalid request method"
//	@Router			/crud/{id} [delete]
func (s *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	// Zbudowanie odpowiedzi
	response := model.ResponseBody{
		Message: "Data DELETED successfully for id: " + id,
		Status:  "success",
	}

	// Zwrócenie odpowiedzi w formacie JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Failed to write response:", err)
	}
}
