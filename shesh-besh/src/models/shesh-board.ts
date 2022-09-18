// import { Cell } from "./cell";

// export interface SheshBoard{

//   cells: Cell[];
//   jail: {
//     blackPieces: number;
//     whitePieces: number;
//   };
//   turn: "black" | "white";
//   winner: "black" | "white" | null;
//   dice: {
//     first: number | null;
//     second: number | null;
//   };

//   constructor() {
//     this.cells = [];
//     this.jail = {
//       blackPieces: 0,
//       whitePieces: 0,
//     };
//     this.turn = "black";
//     this.winner = null;
//     this.dice = {
//       first: null,
//       second: null,
//     };

//     this.cells = this.createCells();
//     this.populateCells("init");
//   }

//   createCells(): Cell[] {
//     const cells: Cell[] = [];
//     for (let i = 0; i < 24; i++) {
//       cells.push({
//         id: i,
//         displayNumber: i + 1,
//         blackPieces: 0,
//         whitePieces: 0,
//       });
//     }
//     return cells;
//   }

//   initState = [
//     { b: 2 },
//     {},
//     {},
//     {},
//     {},
//     { w: 5 },
//     {},
//     { w: 3 },
//     {},
//     {},
//     {},
//     { b: 5 },
//     { w: 5 },
//     {},
//     {},
//     {},
//     { b: 3 },
//     {},
//     { b: 5 },
//     {},
//     {},
//     {},
//     {},
//     { w: 2 },
//   ];

//   populateCells(type) {
//     if (type == "init") {
//       this.initState.forEach((cell, index) => {
//         if (cell.b) {
//           this.cells[index].blackPieces = cell.b;
//         }
//         if (cell.w) {
//           this.cells[index].whitePieces = cell.w;
//         }
//       });
//     }
//   }

//   asDisplayRows() {
//     const rows = [];
//     for (let i = 0; i < 4; i++) {
//         rows.push(this.cells.slice(i * 6, i * 6 + 6));
//     }
//   }
// }
