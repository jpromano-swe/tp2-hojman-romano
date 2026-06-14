package hospital

import (
	"bufio"
	"os"
)

func abrirArchivoCSV(archivo string) *os.File {
	datos, err := os.Open(archivo)
	if err != nil {
		panic("Error al leer el archivo")
	}
	return datos
}

func leerLineasArchivo(archivo string) []string {
	datos := abrirArchivoCSV(archivo)
	defer datos.Close()

	var lineas []string
	scanner := bufio.NewScanner(datos)

	for scanner.Scan() {
		lineas = append(lineas, scanner.Text())
	}
	if scanner.Err() != nil {
		panic("Error al leer el archivo")
	}
	return lineas
}
