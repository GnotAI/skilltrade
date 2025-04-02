package auth

import (
	"io"
	"net/http"

	"github.com/GnotAI/skilltrade/internal/users"
	"github.com/bytedance/sonic"
)

// AuthHandler struct with services
type AuthHandler struct {
	Service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// SignUpHandler handles new user signups
func (h *AuthHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
  body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
	}

	// Parse request body
	if err := sonic.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create a user
	user := &users.User{
		Email:    req.Email,
		Password: req.Password, 
		FullName: req.FullName,
	}

	// Call the service to create the user and generate the JWT token
	err = h.Service.SignUp(user)
	if err != nil {
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the JWT token to the user
	resp, _ := sonic.Marshal(map[string]string{"message": "User created successfully"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// SignInHandler handles user sign-ins (login)
func (h *AuthHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
  body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse request body
	if err := sonic.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service to validate the credentials and generate JWT token
	token, err := h.Service.SignIn(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	// Return the JWT token to the user
	resp, err := sonic.Marshal(map[string]string{
		"token": token,
	}) 
  w.Header().Set("Content-Type", "application/json")
  w.Write(resp)
}

// RefreshHandler handles refreshing of the JWT token
func (h *AuthHandler) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	// Get the token from the Authorization header
  authHeader, ok := r.Context().Value("AuthorizationToken").(string)
	if !ok {
		http.Error(w, "Authorization token not found", http.StatusUnauthorized)
		return
	}

	// Call the service to refresh the JWT token
	newToken, err := h.Service.RefreshToken(authHeader)
	if err != nil {
		http.Error(w, "Failed to refresh token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Authorization", "Bearer "+newToken)
	// Return the refreshed token to the user
	resp, _ := sonic.Marshal(map[string]string{
		"token": newToken,
	})
	w.Header().Set("Content-Type", "application/json")
  w.Write(resp)
}
