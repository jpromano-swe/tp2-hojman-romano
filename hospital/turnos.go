package hospital

import (
	"fmt"
	"tp2/mensajes"
	"tp2/tdas/cola"
	"tp2/tdas/cola_prioridad"
)

const (
	TURNO_URGENTE = "URGENTE"
	TURNO_REGULAR = "REGULAR"
)

func crearColasTurnos(funcion_cmp func(*Paciente, *Paciente) int) (cola.Cola[*Paciente], cola_prioridad.ColaPrioridad[*Paciente]) {
	urgentes := cola.CrearColaEnlazada[*Paciente]()
	regulares := cola_prioridad.CrearHeap(funcion_cmp)
	return urgentes, regulares
}

func sigTurno(medico *Medico) string {
	especialidad := medico.especialidad
	if !especialidad.turnosUrgentes.EstaVacia() {
		paciente_atendido := especialidad.turnosUrgentes.Desencolar()
		especialidad.cantidadEspera--
		medico.pacientesAtendidos++
		return fmt.Sprintf(mensajes.PACIENTE_ATENDIDO, paciente_atendido.nombre) + fmt.Sprintf(mensajes.CANT_PACIENTES_ENCOLADOS, especialidad.cantidadEspera, especialidad.nombre)
	} else if !especialidad.turnosRegulares.EstaVacia() {
		paciente_atendido := especialidad.turnosRegulares.Desencolar()
		especialidad.cantidadEspera--
		medico.pacientesAtendidos++
		return fmt.Sprintf(mensajes.PACIENTE_ATENDIDO, paciente_atendido.nombre) + fmt.Sprintf(mensajes.CANT_PACIENTES_ENCOLADOS, especialidad.cantidadEspera, especialidad.nombre)
	}
	return fmt.Sprintf(mensajes.SIN_PACIENTES)
}

func asignarTurnos(paciente *Paciente, especialidad *Especialidad, urgencia string) string {
	definirUrgencia(paciente, urgencia)
	especialidad.cantidadEspera++

	if paciente.urgencia == TURNO_URGENTE {
		especialidad.turnosUrgentes.Encolar(paciente)
	} else {
		especialidad.turnosRegulares.Encolar(paciente)
	}
	return fmt.Sprintf(mensajes.PACIENTE_ENCOLADO, paciente.nombre) + fmt.Sprintf(mensajes.CANT_PACIENTES_ENCOLADOS, especialidad.cantidadEspera, especialidad.nombre)
}

func definirUrgencia(paciente *Paciente, urgencia string) {
	paciente.urgencia = urgencia
}
