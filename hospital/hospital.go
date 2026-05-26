package hospital

import (
  "tdas/tdas/cola"
  "tdas/tdas/cola_prioridad"
  "tdas/tdas/diccionario"
)

type Paciente struct {
  nombre       string
  ordenLlegada int
}

type Doctor struct {
  nombre       string
  especialidad *Especialidad
  atendidos    int
}

type Especialidad struct {
  nombre    string
  urgentes  cola.Cola[*Paciente]
  regulares cola_prioridad.ColaPrioridad[*Paciente]
}

type Clinica struct {
  pacientes      diccionario.Diccionario[string, *Paciente]
  especialidades diccionario.Diccionario[string, *Especialidad]
  doctores       diccionario.DiccionarioOrdenado[string, *Doctor]

  proximoOrden int
}
