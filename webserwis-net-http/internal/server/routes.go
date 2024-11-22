package server

import (
	"net/http"
	_ "webserwis/docs"
	"webserwis/utils/middleware"

	// uiHandler "webserwis/utils/ui"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	// swagger := http.NewServeMux()
	file := http.NewServeMux()
	tickets := http.NewServeMux()
	crud := http.NewServeMux()
	// ui := http.NewServeMux()

	// mux.Handle("GET /{$}", uiHandler.UiHandler())
	mux.HandleFunc("POST /login", middleware.LoginHandler)
	mux.HandleFunc("GET /hello-world", s.HelloWorldHandler)

	// swagger.HandleFunc("GET /swagger", httpSwagger.WrapHandler)
	mux.HandleFunc("/swagger", httpSwagger.WrapHandler)

	file.HandleFunc("GET /display", s.ShowFileHandler)
	file.HandleFunc("GET /download", s.DownloadFileHandler)

	tickets.HandleFunc("GET /{id}", s.GetTicketById)
	tickets.HandleFunc("GET /{id}/history", s.GetTicketHistoryByTicketId)
	tickets.HandleFunc("GET /history", s.SearchInAllTicketsHistory)

	crud.HandleFunc("POST /{$}", s.SendItem)
	crud.HandleFunc("PUT /{id}", s.UpdateItem)
	crud.HandleFunc("DELETE /{id}", s.DeleteItem)
	crudWithAuth := middleware.AuthMiddleware(crud)

	// mux.Handle("/docs/", http.StripPrefix("/docs", swagger))
	// mux.Handle("/swagger/", http.StripPrefix("/swagger", swagger))
	mux.Handle("/file/", http.StripPrefix("/file", file))
	mux.Handle("/tickets/", http.StripPrefix("/tickets", tickets))
	// mux.Handle("/crud/", http.StripPrefix("/crud", crud))
	mux.Handle("/crud/", http.StripPrefix("/crud", crudWithAuth))

	// TODO zrobić builda ze svelte i wsadzić go do gotowego serwera i zobaczyć czy działą

	return middleware.Logging(mux)
}
