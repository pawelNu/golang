package model

// TODO napisać modele od nowa

// RequestBody reprezentuje strukturę danych przesyłanych w żądaniu
type RequestBody struct {
	Name  string `json:"name" example:"Jan Kowalski"`              // Imię użytkownika
	Email string `json:"email" example:"jan.kowalski@example.com"` // Adres e-mail użytkownika
}

// ResponseBody reprezentuje strukturę danych zwracanych w odpowiedzi
type ResponseBody struct {
	Message string `json:"message" example:"Data SEND successfully"` // Komunikat do użytkownika
	Status  string `json:"status" example:"success"`                 // Status odpowiedzi
}
