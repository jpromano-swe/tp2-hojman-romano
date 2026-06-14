package hospital

import (
	"fmt"
	"strings"
	"tp2/mensajes"
)

const (
	COMANDO_PEDIR_TURNO       = "PEDIR_TURNO"
	COMANDO_ATENDER_SIGUIENTE = "ATENDER_SIGUIENTE"
	COMANDO_INFORME           = "INFORME"
)

func (hospital *Hospital) EjecutarComando(ingreso string) string {
	partes := strings.SplitN(ingreso, ":", 2)

	if len(partes) != 2 {
		return fmt.Sprintf(mensajes.ENOENT_FORMATO, ingreso)
	}

	switch partes[0] {
	case COMANDO_PEDIR_TURNO:
		return hospital.ChequearIngresoTurnos(partes[1])
	case COMANDO_ATENDER_SIGUIENTE:
		return hospital.ChequearIngresoMedicos(partes[1])
	case COMANDO_INFORME:
		return hospital.ChequearIngresoInforme(partes[1])
	default:
		return fmt.Sprintf(mensajes.ENOENT_CMD, partes[0])
	}
}
