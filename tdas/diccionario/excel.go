package diccionario

import (
    "encoding/csv"
    "io"
    "os"
	"strconv" 
)

func abrirArchivoCSV(archivo string) *os.File{
	datos, err := os.Open(archivo)
    if err != nil {
        panic("Error al leer el archivo")
    }
	return datos
}

func hashearDatosCSV(csvReader *csv.Reader){
	pacientes := CrearHash[string,int]()
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
        pacientes.Guardar(rec[0], edadInt)
    }
}

//a la q llamamos es a esta
func obtenerDatosCSV(archivo string) {
    datos := abrirArchivoCSV(archivo)
	defer datos.Close()
    csvReader := csv.NewReader(datos)
	hashearDatosCSV(csvReader)
}
