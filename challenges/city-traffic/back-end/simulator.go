package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
)

func main() {

	anchoFlag := flag.Int("ancho", 9, "Dato Invalido")
	ncarsFlag := flag.Int("coches", 9, "Dato Invalido")
	nSemFlag := flag.Int("semaforos", 4, "Dato Invalido")


	flag.Parse()

	ancho = *anchoFlag
	nCoches = *ncarsFlag
	nSemaforos = *nSemFlag

	validarAncho()

	r = rand.New(rand.NewSource(time.Now().UnixNano())) // seed
	var streetCells []entidad
	tablero, streetCells = createtablero()

	validarCoches(ancho)

	esquinas := obtenerEsquinas()

	validarSemaforos(len(esquinas))

	crearSemaforo(esquinas)

	ch := make(chan int, nCoches)
	crearCoche(streetCells, &ch)
	mostrarMapa(&ch)

	fmt.Print("Visualizar rutas: (s/n) ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

	for ;input.Text() != "s" && input.Text() != "n"; {
		fmt.Print("Visualizar rutas? (s/n): ")
		input = bufio.NewScanner(os.Stdin)
		input.Scan()
	}

	if input.Text() == "s" {
		imprimirRutas()
	}
	
	//initServer()
}







