package strategy

import "github.com/skycoop/corgi-snek/types"

type CluelessStrategy struct{}

func (s *CluelessStrategy) GetMove(gs *types.GameState) (types.BattlesnakeMoveResponse, error) {
	return types.BattlesnakeMoveResponse{
		Move:  types.MoveUp,
		Shout: "I'm a clueless snek who can only move up",
	}, nil
}
