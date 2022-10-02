import axios from "axios"
import { useBoardState } from "./store/state"

export function initHttp() {
    axios.defaults.baseURL = "http://localhost:8080/api/v1";
    axios.defaults.headers.common["Accept"] = "application/json";
    axios.defaults.headers.common["Content-Type"] = "application/json";
    axios.defaults.transformResponse = [
        function (data) {
            return data ? JSON.parse(data) : data;
        },
    ];
}

export const createNewGame = async () => {
    const data = await axios.get(
        "/new_game"
    );
    return data.data;
};

export const getGame = async () => {
    const data = await axios.get(
        `${gameSlugPrefix()}/get_game`
    );
    return data.data;
};

export const getPossibleMoves = async (selectedCell) => {
    const data = await axios.get(
        `${gameSlugPrefix()}/possible_moves?selectedCell=` + selectedCell.Index
    );
    return data.data;
};

export const selectCell = async (targetCell) => {
    const data = await axios.get(
        `${gameSlugPrefix()}/select_cell?target_cell=` + targetCell.Index
    );
    return data.data;
};

export const movePiece = async (targetCell) => {
    const data = await axios.post(
        `${gameSlugPrefix()}/move_piece?target_cell=` + targetCell.Index,
    );
    return data.data;
};

export const gameSlugPrefix = () => {
    const boardState = useBoardState();
    return `slug=${boardState.slug}`;
}

export const getAiRecommendation = async () => {
    const data = await axios.get(
        `${gameSlugPrefix()}/ai_recommendation`
    );
    return data.data;
}