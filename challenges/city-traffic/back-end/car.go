package main

import(
	"time"
)

func crearCoche(streetCells []entidad, ch *chan int) {
	coches = make([]coche, 0)
	for i := 0; i < nCoches; i++ {
		n := r.Intn(len(streetCells))
		cell1 := streetCells[n]

		streetCells[len(streetCells)-1], streetCells[n] = streetCells[n], streetCells[len(streetCells)-1]
		streetCells = streetCells[:len(streetCells)-1]

		n2 := r.Intn(len(streetCells))
		cell2 := streetCells[n2]

		velocidad := r.Intn(250-50) + 50

		path := obtenerRuta(cell1, cell2)
		rutas = append(rutas, path)

		c := coche{i, cell1.x, cell1.y, velocidad, path, 0, "", ""}

		coches = append(coches, c)
		modificarCoche(c)
	}

	for i := 0; i < len(coches); i++ {
		index := i
		go func() {
			for len(coches[index].path) > 0 {
				time.Sleep(time.Duration(coches[index].velocidad) * time.Millisecond)
				trasladarCoche(&coches[index])
			}
			*ch <- coches[index].Id
			coches[index].velocidad = 0
			removerCoche(&coches[index])
		}()
	}
}

func modificarCoche(c coche) {
	i := c.x
	j := c.y
	if !tablero[i][j].hasCar {
		tablero[i][j].hasCar = true
	}
}

func removerCoche(c *coche) {
	i := c.x
	j := c.y
	if tablero[i][j].hasCar {
		tablero[i][j].hasCar = false
	}
}

func trasladarCoche(c *coche) {
	cX := c.x
	cY := c.y
	nextCell := c.path[0]
	nX := nextCell.x
	nY := nextCell.y
	if !tablero[nX][nY].hasCar && tablero[cX][cY].greenLight {
		tablero[cX][cY].hasCar = false
		c.x = nX
		c.y = nY
		tablero[nX][nY].hasCar = true
		c.path = c.path[1:]
		if c.velocidad > 20 {
			c.velocidad -= 10
		}
		c.idle = 0
	} else {
		if c.idle <= 2 {
			c.idle++
			if c.velocidad < 150 {
				c.velocidad += 10
			}
		} else {
			c.velocidad = 150
		}
	}
}

