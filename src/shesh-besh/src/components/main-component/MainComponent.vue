<template>
  <div class="wrapper">
    <GameLogs v-if="boardIsReady" />
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
};
</script>

<style>
.wrapper {
  display: flex;
  flex-direction: row;
}
</style>
