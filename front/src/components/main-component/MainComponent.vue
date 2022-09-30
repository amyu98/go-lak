<template>
  <div class="wrapper">
    <div>
      <header>
        <div class="title">Shesh Besh</div>
        <button @click="newGame" class="new-game">New Game</button>
      </header>
      <!-- <GameLogs v-if="boardIsReady" /> -->
    </div>
    <MainBoard v-if="boardIsReady" />
  </div>
</template>

<script>
import MainBoard from "../main-board/MainBoard";
import GameLogs from "../game-logs/GameLogs";
import { getNewGame, getGame } from "../../api";
import { useBoardState } from "../../store/state";
import { gameMixin } from "../../mixins/gameMixin";

export default {
  name: "MainComponent",
  components: { MainBoard, GameLogs },
  mixins: [gameMixin],
  data() {
    return {
      boardIsReady: false,
      rawBoard: [],
      state: null,
    };
  },
  async mounted() {
    await gameMixin.methods.startGame();
    this.boardIsReady = true;
  },
  methods: {
    async newGame() {
      await gameMixin.methods.createNewGame();
    },
  },
};
</script>

<style>
.wrapper {
  display: flex;
  flex-direction: row;
}
.title{
  font-size: 26px;
}

.new-game{
  margin-bottom: 30px;
}
</style>
