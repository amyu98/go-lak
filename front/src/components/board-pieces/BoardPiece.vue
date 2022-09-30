<template>
    <div class="board-piece">
        <span class="piece" v-bind:class="pieceClass" >*</span>
    </div>
</template>

<script>
export default {
    props: {
        pieceColor: {
            type: String,
            required: true
        }
    },
    computed: {
        pieceClass() {
            return {
                'black-piece': this.pieceColor === 'black',
                'white-piece': this.pieceColor === 'white'
            }
        }
    },
    mounted() {
    }
}
</script>

<style>
.board-piece {
    display: inline-block;
    width: 20px;
    height: 20px;
    border: 1px solid black;
    border-radius: 50%;
    text-align: center;
    vertical-align: middle;
    line-height: 20px;
    margin: 2px;
}
.black-piece {
    background-color: black;
    color: white;
}
.white-piece {
    background-color: white;
    color: black;
}
</style>