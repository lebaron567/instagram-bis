package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

func createPost(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal the JSON into a Post struct
	var newPost Post
	err = json.Unmarshal(body, &newPost)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate the Post content (optional)
	if newPost.Title == "" || newPost.Content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	// Respond with the created post
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{
		"message": "Post created successfully",
		"title":   newPost.Title,
		"author":  newPost.Author,
	}
	json.NewEncoder(w).Encode(response)

	fmt.Println("Post created:", newPost)
}

func getPostId(w http.ResponseWriter, r *http.Request) {
	// Get the post ID from the URL parameters
	id := chi.URLParam(r, "id")

	// Respond with the post ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Post retrieved successfully",
		"id":      id,
	}
	json.NewEncoder(w).Encode(response)

	fmt.Println("Post retrieved:", id)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	// Get the post ID from the URL parameters
	id := chi.URLParam(r, "id")

	// Respond with the deleted post ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Post deleted successfully",
		"id":      id,
	}
	json.NewEncoder(w).Encode(response)

	fmt.Println("Post deleted:", id)
}

func deletePostUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameters
	id := chi.URLParam(r, "id")

	// Respond with the user ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Posts for user retrieved successfully",
		"user_id": id,
	}
	json.NewEncoder(w).Encode(response)

	fmt.Println("Posts for user retrieved:", id)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	// Respond with all posts
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "All posts retrieved successfully",
	}
	json.NewEncoder(w).Encode(response)

	fmt.Println("All posts retrieved")
}
