package main

import (
	"log"
	"time"
)

func validarAncho() {
	if ancho < 9 || (ancho-2)%7 != 0 {
		log.Fatalf("ERROR: Ancho no aceptado")
	}
}

func validarCoches(ancho int) {
	if nCoches > ancho || nCoches < 0 {
		log.Fatalf("ERROR: Número de coches debe ser menor a Ancho")
	}
}

func validarSemaforos(nIntersections int) {
	if nSemaforos < 0 || nSemaforos > nIntersections {
		log.Fatalf("ERROR: Número de semaforos invalido")
	}
}

func mostrarMapa(ch *chan int) {
	for {
		imprimirTablero()
		if len(*ch) >= nCoches {
			break
		}
		time.Sleep(41 * time.Millisecond)
	}
}