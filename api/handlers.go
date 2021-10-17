package api

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/skycoop/corgi-snek/types"
)

func (c *ApiController) HandleGetBattlesnake(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	err := jsoniter.NewEncoder(w).Encode(c.battlesnakeConfig)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode BattlesnakeInfoResponse")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (c *ApiController) HandleStartGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	state := &types.GameState{}
	err := jsoniter.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Warn().Err(err).Msg("Received invalid StartGame request")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Info().Str("game_id", state.Game.ID).Msg("Game started")
}

func (c *ApiController) HandleMove(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	state := &types.GameState{}
	err := jsoniter.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Warn().Err(err).Msg("Received invalid StartGame request")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	logger := log.With().Str("game_id", state.Game.ID).Logger()

	move, err := c.strategy.GetMove(state)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get move for strategy")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = jsoniter.NewEncoder(w).Encode(move)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to encode BattlesnakeMoveResponse")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (c *ApiController) HandleEndGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	state := &types.GameState{}
	err := jsoniter.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Warn().Err(err).Msg("Received invalid StartGame request")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Info().Str("game_id", state.Game.ID).Msg("Game ended")
}
