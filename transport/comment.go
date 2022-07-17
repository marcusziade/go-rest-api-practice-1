package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marcusziade/go-rest-api-practice-1/comment"
)

// Retrieve a comment by ID
func (handler *Handler) GetComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	vars := mux.Vars(request)
	id := vars["id"]

	i, error := strconv.ParseUint(id, 10, 64)
	if error != nil {
		sendErrorResponse(writer, "Unable to parse UINT from ID", error)
	}
	comment, error := handler.Service.GetComment(uint(i))
	if error != nil {
		sendErrorResponse(writer, "Error Retrieving Comment By ID", error)
	}

	if error := json.NewEncoder(writer).Encode(comment); error != nil {
		panic(error)
	}
}

// Retrieves all comments from the comment service
func (handler *Handler) GetAllComments(writer http.ResponseWriter, request *http.Request) {
	comments, error := handler.Service.GetAllComments()
	if error != nil {
		sendErrorResponse(writer, "Failed to retrieve all comments", error)
	}
	if error := json.NewEncoder(writer).Encode(comments); error != nil {
		panic(error)
	}
}

// Adds a new comment
func (handler *Handler) PostComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if error := json.NewDecoder(request.Body).Decode(&comment); error != nil {
		sendErrorResponse(writer, "Failed to decode JSON Body", error)
	}

	comment, error := handler.Service.PostComment(comment)
	if error != nil {
		sendErrorResponse(writer, "Failed to post new comment", error)
	}
	if error := json.NewEncoder(writer).Encode(comment); error != nil {
		panic(error)
	}
}

// Updates a comment by ID
func (handler *Handler) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	vars := mux.Vars(request)
	id := vars["id"]
	commentID, error := strconv.ParseUint(id, 10, 64)
	if error != nil {
		sendErrorResponse(writer, "Failed to parse uint from ID", error)
	}

	var comment comment.Comment
	if error := json.NewDecoder(request.Body).Decode(&comment); error != nil {
		sendErrorResponse(writer, "Failed to decode JSON Body", error)
	}

	comment, error = handler.Service.UpdateComment(uint(commentID), comment)
	if error != nil {
		sendErrorResponse(writer, "Failed to update comment", error)
	}
	if error := json.NewEncoder(writer).Encode(comment); error != nil {
		panic(error)
	}
}

// Deletes a comment by ID
func (handler *Handler) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	vars := mux.Vars(request)
	id := vars["id"]
	commentID, error := strconv.ParseUint(id, 10, 64)
	if error != nil {
		sendErrorResponse(writer, "Failed to parse uint from ID", error)
	}

	error = handler.Service.DeleteComment(uint(commentID))
	if error != nil {
		sendErrorResponse(writer, "Failed to delete comment by comment ID", error)
	}

	if error := json.NewEncoder(writer).Encode(Response{Message: "Successfully Deleted"}); error != nil {
		panic(error)
	}
}
