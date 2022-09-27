<template>
  
    <MainBoard v-if="boardIsReady"/>

</template>

<script>

import MainBoard from '../main-board/MainBoard';
import { getNewGame, getGame } from '../../api'
import { useBoardState } from "../../store/state"

export default {
    name: "MainComponent",
    components: { MainBoard },
    data(){
        return {
            boardIsReady: false,
            rawBoard: [],
            state: null,
        }
    },
    computed: {
    },
    mounted() {
        this.state = useBoardState();
        getNewGame().then((response) => {
            this.state.setBoard(response);
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