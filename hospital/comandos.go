package hospital
import (
  "strings"
  "fmt"
  "tp2/mensajes"
)

func asignarComando(ingreso string, hospital *Hospital){
  partes := strings.Split(ingreso,";")
  switch partes[0]{
    case "PEDIR_TURNO":
      chequearIngresoTurnos(partes[1], hospital)
    case "ATENDER_SIGUIENTE":
      chequearIngresoMedicos(partes[1], hospital)
    default:
      fmt.Printf(mensajes.ENOENT_CMD, partes[0] )
  }

}
