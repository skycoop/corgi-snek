package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/skycoop/corgi-snek/strategy"
	"github.com/skycoop/corgi-snek/types"
)

type ApiController struct {
	strategy          strategy.Strategy
	battlesnakeConfig *types.BattlesnakeInfoResponse
}

func New(s strategy.Strategy, c *types.BattlesnakeInfoResponse) *ApiController {
	return &ApiController{strategy: s, battlesnakeConfig: c}
}

func (c *ApiController) NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", c.HandleGetBattlesnake)
	router.POST("/start", c.HandleStartGame)
	router.POST("/move", c.HandleMove)
	router.POST("/end", c.HandleEndGame)

	return router
}
