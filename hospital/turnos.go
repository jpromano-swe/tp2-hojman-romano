package hospital

import (
	"fmt"
	"tp2/mensajes"
	"tp2/tdas/cola"
	"tp2/tdas/cola_prioridad"
)

func crearColasTurnos(funcion_cmp func(*Paciente, *Paciente) int) (cola.Cola[*Paciente], cola_prioridad.ColaPrioridad[*Paciente]) {
	urgentes := cola.CrearColaEnlazada[*Paciente]()
	regulares := cola_prioridad.CrearHeap(funcion_cmp)
	return urgentes, regulares
}

func sigTurno(medico *Medico) {
	especialidad := medico.especialidad
	if !especialidad.turnosUrgentes.EstaVacia() {
		paciente_atendido := especialidad.turnosUrgentes.Desencolar()
		especialidad.cantidadEspera--
		fmt.Printf(mensajes.PACIENTE_ATENDIDO, paciente_atendido.nombre)
		fmt.Printf(mensajes.CANT_PACIENTES_ENCOLADOS, especialidad.cantidadEspera, especialidad.nombre)
		medico.pacientesAtendidos++
	} else if !especialidad.turnosRegulares.EstaVacia() {
		paciente_atendido := especialidad.turnosRegulares.Desencolar()
		especialidad.cantidadEspera--
		fmt.Printf(mensajes.PACIENTE_ATENDIDO, paciente_atendido.nombre)
		fmt.Printf(mensajes.CANT_PACIENTES_ENCOLADOS, especialidad.cantidadEspera, especialidad.nombre)
		medico.pacientesAtendidos++
	} else {
		fmt.Printf(mensajes.SIN_PACIENTES)
	}
}

func asignarTurnos(paciente *Paciente, especialidad *Especialidad, urgencia string) {
	definirUrgencia(paciente, urgencia)
	especialidad.cantidadEspera++

	if paciente.urgencia == "URGENTE" {
		especialidad.turnosUrgentes.Encolar(paciente)
		fmt.Printf(mensajes.PACIENTE_ENCOLADO, paciente.nombre)
	} else {
		especialidad.turnosRegulares.Encolar(paciente)
		fmt.Printf(mensajes.PACIENTE_ENCOLADO, paciente.nombre)
	}
	fmt.Printf(mensajes.CANT_PACIENTES_ENCOLADOS, especialidad.cantidadEspera, especialidad.nombre)
}

func definirUrgencia(paciente *Paciente, urgencia string) {
	paciente.urgencia = urgencia
}
