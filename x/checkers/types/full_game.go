package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/malakaja/cosmo-checkers/checkers/x/checkers/rules"
)

func (s StoredGame) GetBlackAddress() (black sdk.Address, err error) {
	black, errBlack := sdk.AccAddressFromBech32(s.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), s.Black)
}

func (s StoredGame) GetRedAddress() (red sdk.Address, err error) {
	red, errBlack := sdk.AccAddressFromBech32(s.Red)
	return red, sdkerrors.Wrapf(errBlack, ErrInvalidRed.Error(), s.Red)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	return err
}
