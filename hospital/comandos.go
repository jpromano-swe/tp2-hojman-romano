package hospital

import (
  "fmt"
  "strings"
  "tp2/mensajes"
)

func AsignarComando(ingreso string, hospital *Hospital) {
  partes := strings.SplitN(ingreso, ":", 2)

  if len(partes) != 2 {
    fmt.Printf(mensajes.ENOENT_FORMATO, ingreso)
    return
  }

  switch partes[0] {
  case "PEDIR_TURNO":
    chequearIngresoTurnos(partes[1], hospital)
  case "ATENDER_SIGUIENTE":
    chequearIngresoMedicos(partes[1], hospital)
  case "INFORME":
    chequearIngresoInforme(partes[1], hospital)
  default:
    fmt.Printf(mensajes.ENOENT_CMD, partes[0])
  }
}
