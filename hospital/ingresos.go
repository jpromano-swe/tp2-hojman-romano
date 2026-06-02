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
	if !hospital.pacientes.Pertenece(nombre) {
		fmt.Printf(mensajes.ENOENT_PACIENTE, nombre)
		return

	}
	if !hospital.especialidades.Pertenece(especialidad) {
		fmt.Printf(mensajes.ENOENT_ESPECIALIDAD, especialidad)
		return

	}
	if urgencia != "REGULAR" && urgencia != "URGENTE" {
		fmt.Printf(mensajes.ENOENT_URGENCIA, especialidad)
		return

	}
	asignarTurnos(hospital.pacientes.Obtener(nombre), hospital.especialidades.Obtener(especialidad), urgencia)
}

func chequearIngresoMedicos(nombre string, hospital *Hospital) {
	if hospital.medicos.Pertenece(nombre) {
		sigTurno(hospital.medicos.Obtener(nombre))
	} else {
		fmt.Printf(mensajes.ENOENT_DOCTOR, nombre)
		return
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
