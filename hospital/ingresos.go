package hospital

import (
	"fmt"
	"strings"
	"tp2/mensajes"
)

func (hospital *Hospital) ChequearIngresoTurnos(datos string) string {
	partes := strings.Split(datos, ",")
	nombre := partes[0]
	especialidad := partes[1]
	urgencia := partes[2]
	if !hospital.pacientes.Pertenece(nombre) {
		return fmt.Sprintf(mensajes.ENOENT_PACIENTE, nombre)

	}
	if !hospital.especialidades.Pertenece(especialidad) {
		return fmt.Sprintf(mensajes.ENOENT_ESPECIALIDAD, especialidad)

	}
	if urgencia != TURNO_REGULAR && urgencia != TURNO_URGENTE {
		return fmt.Sprintf(mensajes.ENOENT_URGENCIA, especialidad)

	}
	return asignarTurnos(hospital.pacientes.Obtener(nombre), hospital.especialidades.Obtener(especialidad), urgencia)
}

func (hospital *Hospital) ChequearIngresoMedicos(nombre string) string {
	if !hospital.medicos.Pertenece(nombre) {
		return fmt.Sprintf(mensajes.ENOENT_DOCTOR, nombre)
	}
	return sigTurno(hospital.medicos.Obtener(nombre))
}

func (hospital *Hospital) ChequearIngresoInforme(datos string) string {
	partes := strings.Split(datos, ",")

	if len(partes) != 2 {
		return fmt.Sprintf(mensajes.ENOENT_PARAMS, COMANDO_INFORME)
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

	return crearInforme(hospital.medicos, inicio, fin)
}
