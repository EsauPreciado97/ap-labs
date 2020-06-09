package main

import(
	"fmt"
	"strconv"
)

func createtablero() ([][]entidad, []entidad) {

	var tablero = make([][]entidad, 0)
	var streetCells = make([]entidad, 0)

	for i := 0; i < ancho; i++ {
		var line = make([]entidad, 0)
		iMod := i % 7
		for j := 0; j < ancho; j++ {
			jMod := j % 7
			if iMod < 2 || jMod < 2 {
				dir := NO_DIR
				if iMod == 0 && jMod == 0 {
					dir = LDOWN
				} else if iMod == 1 && jMod == 0 {
					dir = RDOWN
				} else if iMod == 0 && jMod == 1 {
					dir = LUP
				} else if iMod == 1 && jMod == 1 {
					dir = RUP
				} else if iMod == 0 {
					dir = LEFT
				} else if iMod == 1 {
					dir = RIGHT
				} else if jMod == 0 {
					dir = DOWN
				} else if jMod == 1 {
					dir = UP
				}
				c := entidad{i, j, CALLE, dir, false, true}
				if dir == LEFT || dir == RIGHT || dir == UP || dir == DOWN {
					streetCells = append(streetCells, c)
				}
				line = append(line, c) // That entidad is part of a CALLE
			} else {
				c := entidad{i, j, EDIFICIO, NO_DIR, false, true}
				line = append(line, c) // That entidad is part of a EDIFICIO
			}
		}
		tablero = append(tablero, line)
	}
	return tablero, streetCells

}

func obtenerEsquinas() [][]entidad {
	esquinas := make([][]entidad, 0)
	for i := 0; i < ancho; i += 7 {
		for j := 0; j < ancho; j += 7 {
			esquina := make([]entidad, 0)
			if i > 0 {
				esquina = append(esquina, tablero[i-1][j])
			}
			if j > 0 {
				esquina = append(esquina, tablero[i+1][j-1])
			}
			if i < ancho-2 {
				esquina = append(esquina, tablero[i+2][j+1])
			}
			if j < ancho-2 {
				esquina = append(esquina, tablero[i][j+2])
			}
			esquinas = append(esquinas, esquina)
		}
	}
	return esquinas
}

func imprimirTablero() {

	fmt.Println("\033[H\033[2J")

	for i := 0; i < ancho; i++ {
		line := ""
		for j := 0; j < ancho; j++ {
			if tablero[i][j].hasCar {
				line += " ■ "
			} else if !tablero[i][j].greenLight {
				d := tablero[i][j].dir
				if d == LEFT || d == RIGHT {
					line += "__"
				}
				if d == DOWN || d == UP {
					line += " | "
				}
			} else {
				switch(tablero[i][j].typeOfCell) {
				case CALLE:
					switch(tablero[i][j].dir) {
					case LEFT:
						line += "<<<"
					case RIGHT:
						line += ">>>"
					case DOWN:
						line += " | "
					case UP:
						line += " | "
					case LDOWN:
						line += "  \\"
					case RDOWN:
						line += "  /"
					case LUP:
						line += "/  "
					case RUP:
						line += "\\  "
					default:
						line += "   "
					}
					break;
				case EDIFICIO:
					line += "***"
				}
			}
		}
		if i < nCoches {
			cIndex := strconv.Itoa(i)
			if i < 10 {
				cIndex = "0" + cIndex
			}

			if coches[i].velocidad == 0 {
				line += "	              coche #" + cIndex + ": 		Completado"
			} else {
				if coches[i].velocidad > 240 {
					line += "	              coche #" + cIndex + "'s velocidad: 0 km/h"
				} else {
					line += "	              coche #" + cIndex + "'s velocidad: " + strconv.Itoa(1250 / coches[i].velocidad) + "km/h"
				}
			}
		}
		fmt.Println(line);
	}

}