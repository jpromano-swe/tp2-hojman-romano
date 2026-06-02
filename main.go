package main

import (
	"bufio"
	"fmt"
	"os"
	"tp2/hospital"
	"tp2/mensajes"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print(mensajes.ENOENT_CANT_PARAMS)
		return
	}

	clinica := hospital.CrearHospital(os.Args[1], os.Args[2])
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		hospital.AsignarComando(scanner.Text(), clinica)
	}
}
