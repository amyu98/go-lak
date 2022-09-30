import { defineStore } from "pinia";

export const useBoardState = defineStore("boardState", {
  state: () => {
    return {
      slug: null,
      board: null,
      currentPlayer: null,
      whiteJail: null,
      blackJail: null,
      diceRoll: null,
      usedMoves: [],
      selectedCell: null,
      possibleMoves: [],
      logs: []
    };
  },
  actions: {
    setBoardState(boardState) {
      this.slug = boardState.Slug;
      this.board = boardState.Board;
      this.currentPlayer = boardState.CurrentPlayer;
      this.whiteJail = boardState.WhiteJail;
      this.blackJail = boardState.BlackJail;
      this.diceRoll = boardState.DiceRoll;
      this.usedMoves = boardState.UsedMoves;
      this.selectedCell = boardState.SelectedCell;
      this.possibleMoves = boardState.PossibleMoves;
      this.logs = boardState.Logs;
    },
    setGameSlug(slug) {
      this.slug = slug;
    },
    actionState() {
      if (this.selectedCell != -99) {
        return "moving";
      }
      return "selecting";
    },
    setDice(dice) {
      this.dice = dice;
    },
  },
});
