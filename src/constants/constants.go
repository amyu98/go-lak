package constants

import (
	"github.com/amyu98/go-lak/src/models"
)

func GetInitState() [24]models.Cell {

	return [24]models.Cell{
		{Index: 0, WhitePieces: 0, BlackPieces: 2},
		{Index: 1, WhitePieces: 0, BlackPieces: 0},
		{Index: 2, WhitePieces: 0, BlackPieces: 0},
		{Index: 3, WhitePieces: 0, BlackPieces: 0},
		{Index: 4, WhitePieces: 0, BlackPieces: 0},
		{Index: 5, WhitePieces: 5, BlackPieces: 0},
		{Index: 6, WhitePieces: 0, BlackPieces: 0},
		{Index: 7, WhitePieces: 3, BlackPieces: 0},
		{Index: 8, WhitePieces: 0, BlackPieces: 0},
		{Index: 9, WhitePieces: 0, BlackPieces: 0},
		{Index: 10, WhitePieces: 0, BlackPieces: 0},
		{Index: 11, WhitePieces: 0, BlackPieces: 5},
		{Index: 12, WhitePieces: 5, BlackPieces: 0},
		{Index: 13, WhitePieces: 0, BlackPieces: 0},
		{Index: 14, WhitePieces: 0, BlackPieces: 0},
		{Index: 15, WhitePieces: 0, BlackPieces: 0},
		{Index: 16, WhitePieces: 0, BlackPieces: 3},
		{Index: 17, WhitePieces: 0, BlackPieces: 0},
		{Index: 18, WhitePieces: 0, BlackPieces: 5},
		{Index: 19, WhitePieces: 0, BlackPieces: 0},
		{Index: 20, WhitePieces: 0, BlackPieces: 0},
		{Index: 21, WhitePieces: 0, BlackPieces: 0},
		{Index: 22, WhitePieces: 0, BlackPieces: 0},
		{Index: 23, WhitePieces: 2, BlackPieces: 0},
	}

}
