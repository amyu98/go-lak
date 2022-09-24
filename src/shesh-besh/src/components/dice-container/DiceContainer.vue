<template>
  <div class="dice-container">
    <div class="dice" v-for="dice in dices" :key="dice">
      <GameDice :dice="dice" />
    </div>
  </div>
</template>

<script>
import GameDice from "../game-dice/GameDice";
export default {
  name: "DiceContainer",
  components: { GameDice },
  props: {
    dices: {
      type: Array,
      required: true,
    },
  },
};
</script>

<style></style>
