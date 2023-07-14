package main

import (
	"net/http"
	"fmt"

	"github.com/Jasleen8801/rssagg/internal/auth"
	"github.com/Jasleen8801/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User) 

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithErr(w, 403, fmt.Sprintf("Error getting API key: %s", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil{
			respondWithErr(w, 400, fmt.Sprintf("Error getting user: %s", err))
			return
		}

		handler(w, r, user)
	}
}