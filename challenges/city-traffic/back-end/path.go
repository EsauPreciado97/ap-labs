package main

import (
	"github.com/golang-collections/collections/stack" // go get github.com/golang-collections/collections/stack
	"strconv"
	"fmt"
)

func obtenerRuta(source entidad, destination entidad) []entidad {
	q := BFS{nil, source}
	visited := make([]entidad, 0)
	queue := make([]BFS, 0)
	queue = append(queue, q)
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.c == destination {
			return construirRuta(&curr)
		}
		visited = append(visited, curr.c)
		vecinos := obtenerVecinos(visited, curr.c)
		for i := 0; i < len(vecinos); i++ {
			q2 := BFS{&curr, vecinos[i]}
			if !Visitados(visited, q2.c) {
				queue = append(queue, q2)
			}
		}
	}
	return nil
}

func obtenerVecinos(visited []entidad, source entidad) []entidad {
	var vecinos = make([]entidad, 0)
	x := source.x
	y := source.y

	d := source.dir
	if d == LEFT || d == LDOWN || d == LUP {
		if y > 0 {
			c := tablero[x][y-1]
			if !Visitados(visited, c) && c.typeOfCell != EDIFICIO {
				vecinos = append(vecinos, c)
			}
		}
	}

	if d == DOWN || d == LDOWN || d == RDOWN {
		if x < ancho-1 {
			c := tablero[x+1][y]
			if !Visitados(visited, c) && c.typeOfCell != EDIFICIO {
				vecinos = append(vecinos, c)
			}
		}
	}

	if d == RIGHT || d == RDOWN || d == RUP {
		if y < ancho-1 {
			c := tablero[x][y+1]
			if !Visitados(visited, c) && c.typeOfCell != EDIFICIO {
				vecinos = append(vecinos, c)
			}
		}
	}

	if d == UP || d == LUP || d == RUP {
		if x > 0 {
			c := tablero[x-1][y]
			if !Visitados(visited, c) && c.typeOfCell != EDIFICIO {
				vecinos = append(vecinos, c)
			}
		}
	}

	return vecinos
}

func Visitados(visited []entidad, c entidad) bool {
	length := len(visited)
	for i := 0; i < length; i++ {
		if visited[i] == c {
			return true
		}
	}
	return false
}

func construirRuta(q *BFS) []entidad {

	path := make([]entidad, 0)
	s := stack.New()
	curr := q
	for curr != nil {
		s.Push(curr.c)
		curr = curr.previous
	}
	s.Pop()
	for s.Len() > 0 {
		path = append(path, s.Pop().(entidad))
	}
	return path
}

func imprimirRutas() {

	for i := 0; i < len(rutas); i++ {
		// Accede al coche columna i

		for j := 0; j < len(rutas[i]); j++ {
			// Accede a elementos del coche fila j
			fmt.Printf("Ruta de coche %d , (%d,%d)", coches[i].Id, rutas[i][j].x, rutas[i][j].y)
			fmt.Println()


			pathOnX = pathOnX + "," + strconv.Itoa(rutas[i][j].x)
			pathOnY = pathOnY + "," + strconv.Itoa(rutas[i][j].y)
			
			if j == len(rutas[i])-1 {
				coches[i].RouteX = pathOnX 
				coches[i].RouteY = pathOnY
				pathOnX = ""
				pathOnY = ""
			}
		}
	}
}

func imprimirEntidad(c entidad) {
	fmt.Println("x:", c.x, ", y:", c.y)
}

func imprimirEntidadRebanada(slice []entidad) {
	length := len(slice)
	for i := 0; i < length; i++ {
		imprimirEntidad(slice[i])
	}
}