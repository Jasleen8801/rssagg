package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jasleen8801/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(r.Body)
	if err != nil {
		respondWithErr(w, 400, "Error parsing JSON")
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: uuid.MustParse(params.FeedID),
	})
	if err != nil {
		respondWithErr(w, 500, "Error creating feed follow")
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}
