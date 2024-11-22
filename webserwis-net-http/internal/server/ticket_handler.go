package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

// GetTicketById godoc
//
//	@Summary		Get ticket by id
//	@Description	Returns ticket
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Ticket Id"
//	@Success		200	{object}	map[string]any
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/tickets/{id} [get]
func (s *Server) GetTicketById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	resp := map[string]any{
		"zgl_id":        id,
		"zgl_id_string": "string: " + string(id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GetTicketHistoryByTicketId godoc
//
//	@Summary		Get ticket history by ticket id
//	@Description	Returns ticket history
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Ticket Id"
//	@Success		200	{object}	map[string]any
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/tickets/{id}/history [get]
func (s *Server) GetTicketHistoryByTicketId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	history := []string{"wpis1", "wpis2", "wpis3"}

	resp := map[string]any{
		"zgl_id":  id,
		"history": history,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// SearchInAllTicketsHistory godoc
//
//	@Summary		Search keywords in all tickets history
//	@Description	Returns ticket history containing searched keywords
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			keywords	query		[]string	true	"List of keywords" collectionFormat(csv)
//	@Success		200			{object}	map[string]any
//	@Failure		500			{string}	string	"Internal Server Error"
//	@Router			/tickets/history [get]
func (s *Server) SearchInAllTicketsHistory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	keywordsString := query.Get("keywords")
	keywords := strings.Split(keywordsString, ",")

	resp := map[string]any{
		"tickets_history": "tickets_history",
		"keywords":        keywords,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// TODO przeszukiwanie historii zgłoszeń
// TODO wyświetlanie listy zgł do zrobienia dla danego usera
// TODO wyświetlanie wpisanych nakładów dla danego usera
// TODO wyświetlanie zrobionych zgł przez danego usera
// TODO dodawanie, modyfikowanie, usuwanie nakładów
// TODO dodawanie, modyfikowanie, usuwanie załączników
// TODO pomyśleć nad tym, ale to będzie możliwe odczyt, wyświetlanie treści załącznika
// TODO dodawanie nowych zgł, podzgł
// TODO zmiana stanu i wpisywanie historii zgł
// TODO edycja historii zgł (tylko niektórych przypadków) - to mogę przenieść z kodu już używanej przeze mnie aplikacji
