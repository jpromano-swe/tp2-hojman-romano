package hospital
import (
	"fmt"
	"tp2/mensajes"
	"strings"
)
func chequearIngresoTurnos(datos string, hospital *Hospital) {
	partes := strings.Split(datos, ",")
	nombre := partes[0]
	especialidad := partes[1]
	urgencia := partes[2]
	if hospital.pacientes.Pertenece(nombre) {
		fmt.Printf(mensajes.ENOENT_PACIENTE, nombre)
			panic("Panic")

	}
	if hospital.especialidades.Pertenece(especialidad) {
		fmt.Printf(mensajes.ENOENT_ESPECIALIDAD, especialidad)
			panic("Panic")

	}
	if urgencia != "REGULAR" && urgencia != "URGENTE" {
		fmt.Printf(mensajes.ENOENT_URGENCIA, especialidad)
			panic("Panic")

	}
	asignarTurnos(hospital.pacientes.Obtener(nombre),hospital.especialidades.Obtener(nombre),urgencia)
}

func chequearIngresoMedicos(nombre string, hospital *Hospital){
	if hospital.medicos.Pertenece(nombre){
		sigTurno(hospital.medicos.Obtener(nombre))
	}else{
		fmt.Printf(mensajes.ENOENT_DOCTOR, nombre)
		panic("Panic")
	}
}