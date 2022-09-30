<template>
  <div class="board-wrapper">
    <div class="turn">Turn: {{ this.turn?.toUpperCase() }}</div>
    <DiceContainer :dice="this.dice" />
    {{ this.victory ? this.victory + " Won!" : "" }}
    <div class="center-container">
      <div class="goals">
        <div
          class="white-goal goal"
          :class="{
            enabled: this.goalsEnabled['white'],
            optinal: this.goalsOptinal['white'],
          }"
          @click="goalClicked('white')"
        >
          {{ this.whiteGoals }} White Goal
        </div>
        <div
          class="black-goal goal"
          :class="{
            enabled: this.goalsEnabled['black'],
            optinal: this.goalsOptinal['black'],
          }"
          @click="goalClicked('black')"
        >
          {{ this.blackGoals }} Black Goal
        </div>
      </div>
      <JailContainer @jailClicked="jailClicked" />
      <div class="board">
        <div
          class="row"
          v-for="(row, index) in rows"
          v-bind:class="rowClass(index)"
        >
          <div class="section" v-for="section in row">
            <div class="cell" v-for="cell in section">
              <BoardCell
                :cell="cell"
                @cellClick="cellClick"
                :optinalMove="possibleMoves.includes(cell.Index)"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BoardCell from "../board-cell/BoardCell";
import DiceContainer from "../dice-container/DiceContainer";
import JailContainer from "../jail-container/JailContainer";
import { getPossibleMoves, movePiece } from "../../api";
import { useBoardState } from "../../store/state";
import { gameMixin } from "../../mixins/gameMixin";

export default {
  name: "MainBoard",
  components: { BoardCell, DiceContainer, JailContainer },
  props: {},
  data() {
    return {
      rows: [],
      dice: [],
      possibleMoves: [],
      turn: null,
      blackGoals: 0,
      whiteGoals: 0,
      goalsEnabled: {
        white: false,
        black: false,
      },
      goalsOptinal: {
        white: false,
        black: false,
      },
      victory: null,
    };
  },
  mounted() {
    this.state = useBoardState();
    this.updateFromState();
    this.state.$subscribe(() => {
      this.updateFromState();
    });
  },
  methods: {
    rowClass(index) {
      if (index == 0) {
        return "top-row";
      }
      return "bottom-row";
    },
    rawBoardToRows(rawBoard) {
      const rows = [];
      for (let i = 0; i < 2; i++) {
        const row = [];
        for (let j = 0; j < 2; j++) {
          const section = [];
          for (let k = 0; k < 6; k++) {
            const cell = rawBoard[i * 12 + j * 6 + k];
            section.push(cell);
          }
          row.push(section);
        }
        rows.push(row);
      }
      return rows;
    },
    async cellClick(cell) {
      await gameMixin.methods.cellClicked(cell);
    },
    async jailClicked(color) {
      await gameMixin.methods.jailClicked(color);
    },
    async goalClicked(color) {
      await gameMixin.methods.goalClicked(color);
    },
    async movePiece(cell) {
      this.state.movePiece(cell);
    },
    updateFromState() {
      this.state = useBoardState();
      this.rows = this.rawBoardToRows(this.state.board);
      this.dice = this.state.diceRoll;
      this.turn = this.state.currentPlayer;
      this.possibleMoves = this.state.possibleMoves;
      this.goalsEnabled = this.state.goalsEnabled;
      this.goalsOptinal = this.state.goalsOptinal;
      this.victory = this.state.victory;
      this.blackGoals = this.state.blackGoals;
      this.whiteGoals = this.state.whiteGoals;
    },
  },
};
</script>

<style>
.center-container {
  display: flex;
}

.goals {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 120px;
}

.goal {
  border: 1px solid black;
  padding: 3px;
  height: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.enabled {
  background: #d6ffd6;
  border-radius: 5%;
}

.optinal {
  border: 4px solid brown;
}

.turn {
  border: 1px solid black;
  margin: 12px;
  padding: 12px;
  font-weight: 600;
}

.board-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 10px;
}
.board {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 45vh;
}

.section {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  border: 3px solid black;
  margin: 1px;
}

.row {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 45%;
  margin: 3px;
}

.top-row {
  flex-direction: row;
}

.bottom-row {
  flex-direction: row-reverse;
}

.bottom-row .section {
  flex-direction: row-reverse;
}

.bottom-row .cell {
  justify-content: end;
}
</style>
