<template>
  
    <MainBoard v-if="boardIsReady" :rawBoard="rawBoard"/>

</template>

<script>

import MainBoard from '../main-board/MainBoard';
import { getNewGame, getGame } from '../../api'

export default {
    name: "MainComponent",
    components: { MainBoard },
    data(){
        return {
            boardIsReady: false,
            rawBoard: []
        }
    },
    computed: {
    },
    mounted() {
        getNewGame().then((response) => {
            this.rawBoard = response;
            this.boardIsReady = true;
        });
    },
    methods: {
    }
}

</script>

<style>

</style>