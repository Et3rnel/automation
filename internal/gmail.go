// Package auth provides a one-shot command to obtain Gmail OAuth2 tokens.
// This should be run once to get the access and refresh tokens.
// After obtaining the tokens, they should be stored securely and reused in the main application.
package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauth2Config = &oauth2.Config{
	ClientID:     "todo",
	ClientSecret: "todo",
	RedirectURL:  "http://localhost:8080/oauth2callback",
	Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
	Endpoint:     google.Endpoint,
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		return
	}

	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Code exchange failed: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Access Token: %s\nRefresh Token: %s", token.AccessToken, token.RefreshToken)
}

func Run() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/oauth2callback", handleOAuth2Callback)

	log.Println("Server is running at http://localhost:8080/")
	log.Println("Please open http://localhost:8080/ in your browser to start the OAuth flow")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
