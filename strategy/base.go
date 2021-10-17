package strategy

import "github.com/skycoop/corgi-snek/types"

type Strategy interface {
	GetMove(*types.GameState) (types.BattlesnakeMoveResponse, error)
}
