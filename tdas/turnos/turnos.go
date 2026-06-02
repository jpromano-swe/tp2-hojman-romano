package turnos

import (
	hospital "tdas/hospital"
	COLA_ENL "tdas/tdas/cola"
	COLA_PRI "tdas/tdas/cola_prioridad"

)

type Turnos[T any] interface {
}

func CrearColas(funcion_cmp func(*hospital.Paciente, *hospital.Paciente) int, especialidad hospital.Especialidad){
	especialidad.TurnosUrgentes = COLA_ENL.CrearColaEnlazada[*hospital.Paciente]()
	especialidad.TurnosRegulares = COLA_PRI.CrearHeap(funcion_cmp)
}

func AsignacionTurnos()
