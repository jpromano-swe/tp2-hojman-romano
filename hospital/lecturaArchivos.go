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

/* REFACTOR: Habria que borrar eso antes de entregar si estas de acuerdo con el refactor

func hashearDatosCSV(csvReader *csv.Reader) diccionario.Diccionario[string, *Paciente] {
	pacientes := diccionario.CrearHash[string, *Paciente]()
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("Error al leer el archivo")
		}
		edadInt, err := strconv.Atoi(rec[1])
		if err != nil {
			panic("El valor de la columna no es un número válido: " + rec[1])
		}
		paciente := crearPaciente(rec[0], edadInt)
		pacientes.Guardar(rec[0], paciente)
	}
	return pacientes
}

// a la q llamamos es a esta
func obtenerDatosCSV(archivo string) diccionario.Diccionario[string, *Paciente] {
	datos := abrirArchivoCSV(archivo)
	defer datos.Close()
	csvReader := csv.NewReader(datos)
	return hashearDatosCSV(csvReader)
}

*/
