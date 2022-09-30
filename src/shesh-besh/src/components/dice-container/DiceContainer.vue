<template>
  <div class="dice-container">
    <div class="dice" v-for="d in dice" :key="d">
      <GameDice :dice="d" />
    </div>
  </div>
</template>

<script>
import GameDice from "../game-dice/GameDice";
export default {
  name: "DiceContainer",
  components: { GameDice },
  props: {
    dice: {
      type: Array,
      required: true,
    },
  },
};
</script>

<style>
  .dice-container {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 10px;
  }
</style>
