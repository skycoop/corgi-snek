package main

import (
	"net/http"

	"github.com/skycoop/corgi-snek/api"
	"github.com/skycoop/corgi-snek/strategy"
	"github.com/skycoop/corgi-snek/types"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	strat := &strategy.CluelessStrategy{}
	controller := api.New(strat, &types.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "skycoop",
		Color:      "#123F92",
		Tail:       "rbc-necktie",
		Head:       "silly",
		Version:    "testtesttest",
	})
	router := controller.NewRouter()

	log.Info().Msg("Starting snek server")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while running server")
	}
}
