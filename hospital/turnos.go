package hospital

import COLA_ENL"tdas/tdas/cola"

func CrearColas() COLA_ENL.Cola[T any],  {
	urgencia:= COLA_ENL.CrearColaEnlazada[*Paciente]
	
	return urgencia
}

func AsignacionTurnos()