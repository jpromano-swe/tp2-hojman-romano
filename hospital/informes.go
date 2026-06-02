package hospital

import (
	"fmt"
	"tp2/mensajes"
	"tp2/tdas/diccionario"
)

func crearInforme(arbol diccionario.DiccionarioOrdenado[string, *Medico], inicio, fin *string) {
	cantidadMedicos := 0

	arbol.IterarRango(inicio, fin, func(nombreMedico string, medico *Medico) bool {
		cantidadMedicos++
		return true
	})

	fmt.Printf(mensajes.DOCTORES_SISTEMA, cantidadMedicos)

	posicion := 1

	arbol.IterarRango(inicio, fin, func(nombreMedico string, medico *Medico) bool {
		fmt.Printf(mensajes.INFORME_DOCTOR, posicion, nombreMedico, medico.especialidad.nombre, medico.pacientesAtendidos)
		posicion++
		return true
	})
}
