package main

import (
	"time"
)

func crearSemaforo(esquinas [][]entidad) {
	semaforos := make([]semaforo, 0)

	for i := 0; i < nSemaforos; i++ {
		n := r.Intn(len(esquinas))

		esquina := esquinas[n]

		esquinas[len(esquinas)-1], esquinas[n] = esquinas[n], esquinas[len(esquinas)-1]
		esquinas = esquinas[:len(esquinas)-1]

		velocidad := r.Intn(1200-800) + 800
		s := semaforo{esquina, 0, velocidad}
		iniciarEstados(&s)
		semaforos = append(semaforos, s)
	}

	for i := 0; i < len(semaforos); i++ {
		index := i
		go func() {
			for {
				cambiarEstado(&semaforos[index])
				time.Sleep(time.Duration(semaforos[index].velocidad) * time.Millisecond)
			}
		}()
	}

}

func iniciarEstados(s *semaforo) {

	for i := 0; i < len(s.cells); i++ {
		x := s.cells[i].x
		y := s.cells[i].y
		tablero[x][y].greenLight = false
	}
}

func cambiarEstado(s *semaforo) {
	length := len(s.cells)
	cX := s.cells[s.index].x
	cY := s.cells[s.index].y
	tablero[cX][cY].greenLight = false
	s.index = (s.index + 1) % length
	nX := s.cells[s.index].x
	nY := s.cells[s.index].y
	tablero[nX][nY].greenLight = true
}