package main

import (
	"bufio"
	"fmt"
	"tp2/hospital"
)

func procesarEntrada(clinica *hospital.Hospital, scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Print(clinica.EjecutarComando(scanner.Text()))
	}
}
