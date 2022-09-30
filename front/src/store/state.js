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
      logs: [],
      goalsEnabled: {},
      blackGoals: 0,
      whiteGoals: 0,
      victory: null,
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
      this.whiteGoals = boardState.WhiteGoals;
      this.blackGoals = boardState.BlackGoals;
      this.victory = boardState.Victory;
      this.goalsEnabled = this.getGoalsEnabled();
      this.goalsOptinal = this.getGoalsOptional();
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
    getGoalsEnabled() {
      const res = {};

      ["white", "black"].forEach((color) => {
        const home =
          color == "white" ? [0, 1, 2, 3, 4, 5] : [18, 19, 20, 21, 22, 23];

        const playableOutsideHome = this.board.filter((cell) => {
          if (home.includes(cell.Index)) {
            return false;
          }
          const friendlyPieces =
            color == "black" ? cell.BlackPieces : cell.WhitePieces;
          return friendlyPieces > 0;
        });
        const friendlyJail = color == "black" ? this.blackJail : this.whiteJail;
        const goalEnabled =
          playableOutsideHome.length == 0 && friendlyJail == 0;
        res[color] = goalEnabled;
      });

      return res;
    },
    getGoalsOptional() {
      return {
        white: this.possibleMoves.includes(30),
        black: this.possibleMoves.includes(-30),
      }
    },
  },
});
