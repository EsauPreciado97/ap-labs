package main

import (
	"math/rand"
)

var ancho int
var nCoches int
var nSemaforos int
var r *rand.Rand
var tablero [][]entidad
var rutas [][]entidad
var coches []coche
var pathOnX string
var pathOnY string
var autos string
var semaforos string


const CALLE = 0
const EDIFICIO = 1

const NO_DIR = -1
const LEFT = 0
const RIGHT = 1
const DOWN = 2
const UP = 3
const LDOWN = 4
const RDOWN = 5
const LUP = 6
const RUP = 7

type semaforo struct {
	cells []entidad
	index int
	velocidad int
}

type entidad struct {
	x          int
	y          int
	typeOfCell int
	dir        int
	hasCar     bool
	greenLight bool
}

type coche struct {
	Id    	int		`json:"id"`
	x     	int		
	y     	int
	velocidad 	int		
	path  	[]entidad	
	idle  	int
	RouteX	string	`json:"routeX"`
	RouteY 	string	`json:"routeY"`
}

type BFS struct {
	previous *BFS
	c        entidad
}

