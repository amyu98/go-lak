<template>
  <div class="game-dice" v-bind:class="{ playable: numberIsPlayable }">
    {{ dice }}
  </div>
</template>

<script>
import { useBoardState } from "../../store/state";
export default {
  name: "GameDice",
  data() {
    return {
      state: null
    };
  },
  mounted() {
    this.state = useBoardState();
  },
  computed: {
    numberIsPlayable() {
      if (!this.state?.usedMoves){ return true }
      if (this.state.diceRoll[0] == this.state.diceRoll[1]) {
        return this.state.usedMoves.length < 3 || this.index == 1
      }
      return !this.state.usedMoves.includes(this.dice);
    },
  },
  props: {
    dice: {
      type: Number,
      required: true,
    },
    index: {
      type: Number,
      required: true,
    },
  },
};
</script>

<style>
.game-dice {
  display: inline-block;
  width: 60px;
  height: 60px;
  margin: 10px;
  line-height: 60px;
  text-align: center;
  border: 1px solid black;
  border-radius: 10%;
  color: rgb(228, 228, 228);
  font-weight: 600;
}

.playable {
  color: black;
}
</style>
