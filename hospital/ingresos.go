package hospital

import (
	"fmt"
	"strings"
	"tp2/mensajes"
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
	asignarTurnos(hospital.pacientes.Obtener(nombre), hospital.especialidades.Obtener(nombre), urgencia)
}

func chequearIngresoMedicos(nombre string, hospital *Hospital) {
	if hospital.medicos.Pertenece(nombre) {
		sigTurno(hospital.medicos.Obtener(nombre))
	} else {
		fmt.Printf(mensajes.ENOENT_DOCTOR, nombre)
		panic("Panic")
	}
}

func chequearIngresoInforme(datos string, hospital *Hospital) {
	partes := strings.Split(datos, ",")

	if len(partes) != 2 {
		fmt.Printf(mensajes.ENOENT_PARAMS, "INFORME")
		return
	}

	inicioTexto := partes[0]
	finTexto := partes[1]

	var inicio *string
	var fin *string

	if inicioTexto != "" {
		inicio = &inicioTexto
	}

	if finTexto != "" {
		fin = &finTexto
	}

	crearInforme(hospital.medicos, inicio, fin)
}
