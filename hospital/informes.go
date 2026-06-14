package hospital

import (
	"fmt"
	"tp2/mensajes"
	"tp2/tdas/diccionario"
)

func crearInforme(arbol diccionario.DiccionarioOrdenado[string, *Medico], inicio, fin *string) string {
	cantidadMedicos := 0

	arbol.IterarRango(inicio, fin, func(nombreMedico string, medico *Medico) bool {
		cantidadMedicos++
		return true
	})

	salida := fmt.Sprintf(mensajes.DOCTORES_SISTEMA, cantidadMedicos)

	posicion := 1

	arbol.IterarRango(inicio, fin, func(nombreMedico string, medico *Medico) bool {
		salida += fmt.Sprintf(mensajes.INFORME_DOCTOR, posicion, nombreMedico, medico.especialidad.nombre, medico.pacientesAtendidos)
		posicion++
		return true
	})
	return salida
}
