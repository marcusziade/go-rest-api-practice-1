package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcusziade/go-rest-api-practice-1/comment"
)

// Stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response object
type Response struct {
	Message string
	Error   string
}

// Returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// Sets up all the routes for our application
func (handler *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	handler.Router = mux.NewRouter()

	handler.Router.HandleFunc("/api/comments", handler.GetAllComments).Methods("GET")
	handler.Router.HandleFunc("/api/comment", handler.PostComment).Methods("POST")
	handler.Router.HandleFunc("/api/comment/{id}", handler.GetComment).Methods("GET")
	handler.Router.HandleFunc("/api/comment/{id}", handler.UpdateComment).Methods("PUT")
	handler.Router.HandleFunc("/api/comment/{id}", handler.DeleteComment).Methods("DELETE")

	handler.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, r *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusOK)
		if error := json.NewEncoder(writer).Encode(Response{Message: "I am Alive!"}); error != nil {
			panic(error)
		}
	})
}

func sendErrorResponse(writer http.ResponseWriter, message string, error error) {
	writer.WriteHeader(http.StatusInternalServerError)
	response := Response{Message: message, Error: error.Error()}
	if error := json.NewEncoder(writer).Encode(response); error != nil {
		panic(error)
	}
}
